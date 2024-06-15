package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/auth-service/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/login", handlers.LoginService)
	r.POST("/signup", handlers.SignupService)
	r.GET("/logout", handlers.Logout)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
