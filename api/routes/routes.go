package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/controller"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/user")

	//post create user
	api.POST("/", controller.CreateUser)
	//GET all user
	api.GET("/", controller.FindAllUser)
	//GET user by id
	api.GET("/:id", controller.GetUserByID)
	//update user by id
	api.PATCH("/:id", controller.UpdateUser)

}
