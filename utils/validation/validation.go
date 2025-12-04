package validation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ingenziart/myapp/utils/response"
)

func getValidationErrorMessage(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required ", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s character long", field, param)
	case "email":
		return fmt.Sprintf("%s must be a valid email adres  ", field)
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", field, param)
	}
	return fmt.Sprintf("%s failed validation (%s)", field, tag)

}

func ValidationErrorMessage(c *gin.Context, err error) {
	var validationErrors []string

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		//check if it error and sssign it to error message
		for _, fieldError := range validationErrs {
			message := getValidationErrorMessage(fieldError.Field(), fieldError.Tag(), fieldError.Param())
			validationErrors = append(validationErrors, message)

		}
	} else {
		//non valodator error (different error )
		validationErrors = append(validationErrors, err.Error())
	}
	response.ResponseError(c, http.StatusBadRequest, strings.Join(validationErrors, ", "))
}

func ValidateStruct(c *gin.Context, s interface{}) bool {
	validate := validator.New() //CREATE VALIDATION

	err := validate.Struct(s)
	if err != nil {
		ValidationErrorMessage(c, err)
		return false
	} //USE STRUCT
	return true
}
