package handlers

import (
	"net/http"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

type content struct {
	Article interface{}
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, content{Article: data["payload"]})
	default:
		token, _ := c.Cookie("token")
		logged := false
		if len(token) > 0 {
			user := models.FindByToken(token)
			if user != nil {
				logged = true
			}
		}
		data["logged"] = logged
		c.HTML(http.StatusOK, templateName, data)
	}
}
