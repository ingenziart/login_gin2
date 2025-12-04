package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     string      `json:"status"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}

func ResponseSucess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:     "success",
		Message:    message,
		Data:       data,
		StatusCode: http.StatusOK,
	})

}

func ResponseError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, Response{
		Status:     "error",
		Message:    err,
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	})
}
