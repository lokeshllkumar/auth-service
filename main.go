package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/auth-service/pkg/handlers"
	"github.com/lokeshllkumar/auth-service/pkg/middleware"
	"github.com/lokeshllkumar/auth-service/pkg/storage"
)

func main() {
	db, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler(db))
	r.GET("/oauth/login", handlers.OAuthLoginHandler())
	r.GET("/callback", handlers.OAuthCallbackHandler(db))

	authorized := r.Group("/")
	authorized.Use(middleware.JWTMiddleware())
	authorized.GET("/protected", func(c *gin.Context) {
		username := c.GetString("Username")
		c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username + "!"})
	})
}
