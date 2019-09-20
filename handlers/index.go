package handlers

import (
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage -
func ShowIndexPage(c *gin.Context) {
	articles := models.ShowIndexPage()
	data := gin.H{
		"title":   "Home Page",
		"payload": articles,
	}
	render(c, data, "index.html")
}
