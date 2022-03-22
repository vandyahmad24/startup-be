package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatErrorValidation(err error) []string {
	var errors []string
	tempError := ""
	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "required":
			tempError = fmt.Sprintf("%s is required", v.Field())
		case "email":
			tempError = fmt.Sprintf("%s is not valid email", v.Field())
		case "gte":
			tempError = fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param())
		case "lte":
			tempError = fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param())
		case "min":
			tempError = fmt.Sprintf("%s character must be min %s", v.Field(), v.Param())
		case "max":
			tempError = fmt.Sprintf("%s character must be max %s", v.Field(), v.Param())
		default:
			tempError = v.Error()

		}
		errors = append(errors, tempError)
	}
	return errors
}
