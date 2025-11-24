package response

import (
	"github.com/gin-gonic/gin"
)

func ResponceSucess(c *gin.Context, status int, model interface{}) {
	c.JSON(status, gin.H{
		"status": true,
		"succes": model,
	})

}

func ResponceError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status": false,
		"err":    message,
	})

}
