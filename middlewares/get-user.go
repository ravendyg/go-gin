package middlewares

import (
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// GetUser -
func GetUser(c *gin.Context) {
	token, _ := c.Cookie("token")
	if len(token) > 0 {
		user := models.FindByToken(token)
		if user != nil {
			c.Set("user", user)
		}
	}
}
