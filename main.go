// @title Go Gin User Service API
// @version 1.0
// @description API documentation for the User Service built with Go, Gin, and GORM.
// @termsOfService http://swagger.io/terms/

// @contact.name aime
// @contact.email aimembabazi15@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ingenziart/myapp/api/routes"
	"github.com/ingenziart/myapp/config"
	"github.com/ingenziart/myapp/db"
	_ "github.com/ingenziart/myapp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Create clean Gin router (no default middlewares)
	r := gin.New()

	// Must be FIRST
	//r.SetTrustedProxies([]string{"127.0.0.1"})

	// Load environment variables
	config.LoadEnv()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Database connection
	db.ConnectingDb()

	// API group / versioning
	api := r.Group("/api/v1")
	routes.UserRoutes(api)

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Graceful HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server asynchronously
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
