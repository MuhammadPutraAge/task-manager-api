package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateRequest(s interface{}) map[string]string {
	errors := make(map[string]string)

	if err := validate.Struct(s); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				message := ""

				switch fieldError.Tag() {
				case "required":
					message = fmt.Sprintf("%s must be filled", fieldError.Field())
				case "min":
					message = fmt.Sprintf("%s must be at least %s characters", fieldError.Field(), fieldError.Param())
				case "boolean":
					message = fmt.Sprintf("%s must be a boolean", fieldError.Field())
				default:
					message = fmt.Sprintf("%s has a validation error", fieldError.Field())
				}

				errors[fieldError.Field()] = message
			}
		}
	}

	return errors
}
