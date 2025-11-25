package validation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ingenziart/myapp/models"
	"github.com/ingenziart/myapp/utils/response"
)

func IsValidateStatus(s models.Status) bool {
	switch s {
	case models.StatusActive, models.StatusInactive, models.StatusDeleted:
		return true

	}
	return false

}

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

func ValidationErrorResponse(c *gin.Context, err error) {

	var validationErrors []string

	// Check if it's a validator error
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			msg := getValidationErrorMessage(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			validationErrors = append(validationErrors, msg)
		}
	} else {
		// Non-validation error
		validationErrors = append(validationErrors, err.Error())
	}

	// Send the error response
	response.ResponceError(c, http.StatusBadRequest, strings.Join(validationErrors, ", "))
}
func ValidateStruct(c *gin.Context, s interface{}) bool {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		ValidationErrorResponse(c, err)
		return false
	}
	return true
}
