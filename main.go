package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/config"
	"github.com/ingenziart/myapp/db"
)

// function to load env
func init() {
	config.LoadEnv()
}

func main() {
	r := gin.Default()
	//database connection
	db.ConnectingDb()

	r.Run()

}
