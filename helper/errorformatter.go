package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatErrorValidation(err error) []string {
	tempError := ""
	var errors []string
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, v := range validationErrors {
				switch v.Tag() {
				case "required":
					tempError = fmt.Sprintf("%s is required", v.Field())
				case "email":
					tempError = fmt.Sprintf("%s is not a valid email", v.Field())
				case "gte":
					tempError = fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param())
				case "lte":
					tempError = fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param())
				case "min":
					tempError = fmt.Sprintf("%s character must be at least %s", v.Field(), v.Param())
				case "max":
					tempError = fmt.Sprintf("%s character must be at most %s", v.Field(), v.Param())
				default:
					tempError = v.Error()
				}
				errors = append(errors, tempError)
			}
		} else {
			// Handle the case where err is not a ValidationErrors type
			errors = append(errors, err.Error())
		}
	}
	return errors
}
