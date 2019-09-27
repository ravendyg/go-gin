package handlers

import (
	"net/http"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// ShowUserPage -
func ShowUserPage(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user != nil {
		user := _user.(*models.User)
		data := gin.H{
			"title":    "User",
			"username": user.Username,
		}
		render(c, data, "user.html")
	} else {
		c.Redirect(http.StatusFound, "/u/login")
	}
}
