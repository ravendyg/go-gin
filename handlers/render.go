package handlers

import (
	"net/http"

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
		c.HTML(http.StatusOK, templateName, data)
	}
}
