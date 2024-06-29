package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/auth-service/pkg/auth"
	"github.com/lokeshllkumar/auth-service/pkg/storage"
)

func OAuthLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := "random"
		url := auth.GetOAuthURL(state)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func OAuthCallbackHandler(db *storage.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		state := c.Query("state")
		code := c.Query("code")

		if state != "random" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
			return
		}

		token, err := auth.GetUser(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fexchange token"})
			return
		}

		res, err := auth.GetUserInfo(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
			return
		}
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)

		var userInfo map[string]interface{}
		if err := json.Unmarshal(data, &userInfo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
			return
		}

		username := userInfo["email"].(string)
		db.AddUser(username, "")

		jwtToken, err := auth.GenToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": jwtToken})
	}
}
