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

func ResponceSucess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, data)

}

func ResponceError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, Response{
		Status:     "error",
		Message:    err,
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	})
}

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}
