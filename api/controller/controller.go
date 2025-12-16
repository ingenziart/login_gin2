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

// GetUserByID godoc
// @Summary Get a single user by ID
// @Description Fetch a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response "User retrieved successfully"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [get]
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

// UpdateUserStatus godoc
// @Summary Update a user's status
// @Description Set user's status by ID (body has the new status)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body dto.UpdateStatusDTO true "New status (active|inactive|deleted)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /users/{id}/status [patch]
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

// FindAllUser godoc
// @Summary Get all users with pagination
// @Description Retrieve a paginated list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Success 200 {object} response.Response "Users retrieved successfully"
// @Failure 400 {object} response.Response "Invalid pagination parameters"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [get]
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
