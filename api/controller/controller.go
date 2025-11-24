package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/api/service"
	"github.com/ingenziart/myapp/utils/response"
)

func CreateUser(c *gin.Context) {
	var inputs dto.CreateUserDto
	if err := c.ShouldBind(&inputs); err != nil {
		response.ResponceError(c, http.StatusBadRequest, err.Error())
		return

	}

	user, err := service.CreateUser(inputs)

	if err != nil {
		response.ResponceError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ResponceSucess(c, http.StatusCreated, user)

}

func GetAllUser(c *gin.Context) {

	user, err := service.GetAllUser()
	if err != nil {
		response.ResponceError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.ResponceSucess(c, http.StatusOK, user)

}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		response.ResponceError(c, http.StatusInternalServerError, err.Error())
	}
	response.ResponceSucess(c, http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var inputs dto.UpdateUserDto

	if err := c.ShouldBind(&inputs); err != nil {
		response.ResponceError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	user, err := service.UpdateUser(id, inputs)

	if err != nil {
		response.ResponceError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.ResponceSucess(c, http.StatusOK, user)

}

func UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")

	var inputs dto.UpdateStatusDTO

}
