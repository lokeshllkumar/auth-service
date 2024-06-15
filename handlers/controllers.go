package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/auth-service/models"
)

func LoginService(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func SignupService(c *gin.Context) {
}

func Logout(c *gin.Context) {
}
