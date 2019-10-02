package middlewares

import (
	"net/http"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// GetUser -
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if len(token) > 0 {
			user := models.FindByToken(token)
			if user != nil {
				c.Set("user", user)
				return
			}
		}
	}
}

// RequireAuth -
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_user, _ := c.Get("user")
		if _user == nil {
			c.Redirect(http.StatusFound, "/u/login")
			c.Abort()
		}
	}
}
