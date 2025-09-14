package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var ErrorValidator = map[string]string{
	"min": "The %s field must be at least %s",
	"max": "The %s field must be at most %s",
}

func ErrValidationResponse(err error) ([]ValidationResponse, error) {
	var fieldErrors validator.ValidationErrors
	if !errors.As(err, &fieldErrors) {
		// Not a validation error; let the caller handle it.
		return nil, err
	}

	var out []ValidationResponse
	for _, fe := range fieldErrors {
		switch fe.Tag() {
		case "required":
			out = append(out, ValidationResponse{
				Field:   fe.Field(),
				Message: fmt.Sprintf("The %s field is required", fe.Field()),
			})
		case "email":
			out = append(out, ValidationResponse{
				Field:   fe.Field(),
				Message: fmt.Sprintf("The %s field must be a valid email address", fe.Field()),
			})
		default:
			if tmpl, ok := ErrorValidator[fe.Tag()]; ok {
				switch strings.Count(tmpl, "%s") {
				case 1:
					out = append(out, ValidationResponse{
						Field:   fe.Field(),
						Message: fmt.Sprintf(tmpl, fe.Field()),
					})
				case 2:
					out = append(out, ValidationResponse{
						Field:   fe.Field(),
						Message: fmt.Sprintf(tmpl, fe.Field(), fe.Param()),
					})
				default:
					out = append(out, ValidationResponse{
						Field:   fe.Field(),
						Message: "Invalid validation error template",
					})
				}
			} else {
				out = append(out, ValidationResponse{
					Field:   fe.Field(),
					Message: "Unknown validation error",
				})
			}
		}
	}
	return out, nil
}

func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	logrus.Errorf("%s: %v", message, err)
	return fmt.Errorf("%s: %w", message, err)
}
