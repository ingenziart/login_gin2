package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/api/service"
	"github.com/ingenziart/myapp/utils/response"
	"github.com/ingenziart/myapp/utils/validation"
)

func CreateUser(c *gin.Context) {
	var inputs dto.CreateUserDto

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&inputs); err != nil {
		validation.ValidationErrorMessage(c, err)
		return
	}

	// 2. Validate DTO fields
	if !validation.ValidateStruct(c, inputs) {
		return
	}

	// 3. Call service
	user, err := service.CreateUser(inputs)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 4. Success response
	response.ResponseSucess(c, user, "User created successfully")
}

func GetAllUser()
