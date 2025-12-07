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
	//read
	var inputs dto.CreateUserDto
	if err := c.ShouldBindJSON(&inputs); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	//validation
	if !validation.ValidateStruct(c, inputs) {
		return

	}
	//calling service

	user, err := service.CreateUser(inputs)

	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.ResponseSucess(c, user, "user created successfully")

}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := service.GetUserByID(id)

	if err != nil {
		if err == service.ErrUserNotFound {
			response.ResponseError(c, http.StatusNotFound, err.Error())
			return
		}
		response.ResponseError(c, http.StatusInternalServerError, err.Error())
		return

	}
	response.ResponseSucess(c, user, "success")
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

}
