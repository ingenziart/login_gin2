package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/api/service"
	"github.com/ingenziart/myapp/utils/response"
	"github.com/ingenziart/myapp/utils/validation"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with full name, email, password, phone, status, and role.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User info"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 409 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [post]
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
	var inputs dto.UpdateUserDto

	if err := c.ShouldBindJSON(&inputs); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	//validate
	if !validation.ValidateStruct(c, inputs) {
		return
	}

	user, err := service.UpdateUser(id, inputs)

	if err != nil {
		if err == service.ErrUserNotFound {
			response.ResponseError(c, http.StatusNotFound, err.Error())
			return
		}

		if err == service.ErrEmailInUse {
			response.ResponseError(c, http.StatusConflict, err.Error())
			return
		}
		if err == service.ErrHashPassword {
			response.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

	}
	response.ResponseSucess(c, user, "successfully updated")

}
func FindAllUser(c *gin.Context) {
	pageNumber := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageNumber)

	if err != nil || page < 1 {
		response.ResponseError(c, http.StatusBadRequest, "invalid page number")
		return
	}
	limit, err := strconv.Atoi(pageSize)

	if err != nil || limit < 10 {
		response.ResponseError(c, http.StatusBadRequest, "invalid page size")
		return

	}
	user, err := service.FindAllUser(page, limit)

	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.ResponseSucess(c, user, "users retrieved successfully")
}
