package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/api/service"
	"github.com/ingenziart/myapp/utils/validation"
)

func CreateUser(c *gin.Context) {
	var input dto.CreateUserDto
	if err := c.ShouldBindJSON(&input); err != nil {
		validation.ValidationErrorResponse(c, err)
		return
	}

	if !validation.ValidateStruct(c, &input) {
		return
	}

	user, err := service.CreateUser(input)
	if err != nil {
		response.ResponseError(c, 400, err.Error())
		return
	}

	response.ResponseSuccess(c, user, "User created")

}
