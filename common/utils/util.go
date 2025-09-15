package utils

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

// BindfromJSON loads a JSON config file (filename without extension) from path into dest.
// dest should be a pointer to the target struct or map.
func BindfromJSON(dest any, filename, path string) error {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType("json")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// pass dest (not &dest) because dest is expected to be a pointer already
	if err := v.Unmarshal(dest); err != nil {
		log.Printf("unable to decode into struct: %v", err)
		return err
	}

	return nil
}

// SetEnvFromConsulLKV unmarshals a viper key-value map into environment variables.
func SetEnvFromConsulLKV(v *viper.Viper) error {
	env := make(map[string]any)
	if err := v.Unmarshal(&env); err != nil {
		return err
	}

	for k, val := range env {
		valOf := reflect.ValueOf(val)
		var strVal string

		switch valOf.Kind() {
		case reflect.String:
			strVal = valOf.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strVal = strconv.FormatInt(valOf.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			strVal = strconv.FormatUint(valOf.Uint(), 10)
		case reflect.Float32:
			strVal = strconv.FormatFloat(valOf.Float(), 'f', -1, 32)
		case reflect.Float64:
			strVal = strconv.FormatFloat(valOf.Float(), 'f', -1, 64)
		case reflect.Bool:
			strVal = strconv.FormatBool(valOf.Bool())
		default:
			log.Printf("unsupported type for env var %s: %s", k, valOf.Kind())
			continue
		}

		if err := os.Setenv(k, strVal); err != nil {
			return err
		}
	}

	return nil
}
