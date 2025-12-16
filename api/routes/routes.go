package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/controller"
)

func UserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")

	//post create user
	users.POST("", controller.CreateUser)
	//GET all user
	users.GET("", controller.FindAllUser)
	//GET user by id
	users.GET("/:id", controller.GetUserByID)
	//update user by id
	users.PATCH("/:id", controller.UpdateUser)

	//soft delete
	users.DELETE("/:id", controller.SoftDeleteUser)

	//restore delete
	users.PATCH("/:id/restore", controller.RestoreUser)

}
