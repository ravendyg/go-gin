package handlers

import (
	"net/http"
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

// GetNewArticleForm -
func GetNewArticleForm(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user != nil {
		data := gin.H{}
		render(c, data, "create-article.html")
	} else {
		c.Redirect(http.StatusFound, "/u/login")
	}
}

// CreateArticle -
func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	models.NewArticle(title, content)
	c.Redirect(http.StatusFound, "/")
}
