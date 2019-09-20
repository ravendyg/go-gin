package handlers

import (
	"strconv"
	"web-server/models"

	"github.com/gin-gonic/gin"
)

// GetArticle -
func GetArticle(c *gin.Context) {
	articleID := c.Param("article_id")
	id, err := strconv.ParseInt(articleID, 10, 64)
	if err != nil {
		panic(err)
	}

	article, err := models.GetArticleByID(id)

	data := gin.H{
		"title":   article.Title,
		"payload": article,
	}
	render(c, data, "article.html")
}
