package config

import (
	"os"

	"user-service/common/utils"

	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"signatureKey"`
	Database              Database `json:"database"`
	RateLimiterRequest    float64  `json:"rateLimiterRequest"`
	RateLimiterTimeSecond int      `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string   `json:"jwtSecretKey"`
	JwtExpirationTime     int      `json:"jwtExpirationTime"`
}

type Database struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Name                   string `json:"name"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	MaxOpenConnections     int    `json:"maxOpenConnections"`
	MaxLifetimeConnections int    `json:"maxLifetimeConnections"`
	MaxIdleTime            int    `json:"maxIdleTime"`
	MaxIdleConnections     int    `json:"maxIdleConnections"`
}

func Init() {
	// Try loading from a local file first
	err := utils.BindfromJSON(&Config, "config.json", "-")
	if err != nil {
		// If local file failed, try loading from Consul using environment variables
		err = utils.BindfromJSON(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_KV_PATH"))
		if err != nil {
			logrus.Fatalf("Failed to load config: %v", err)
		}
	}
}
