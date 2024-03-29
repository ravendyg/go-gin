package models

import (
	"errors"
	"sync/atomic"
)

var id int64 = 2

// Article -
type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ArticleList -
var ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// ShowIndexPage -
func ShowIndexPage() []Article {
	return ArticleList
}

// GetArticleByID -
func GetArticleByID(articleID int64) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == articleID {
			return &a, nil
		}
	}

	return nil, errors.New("sdfsd")
}

// NewArticle -
func NewArticle(title string, content string) {
	article := Article{
		ID:      getNextID(),
		Title:   title,
		Content: content,
	}
	ArticleList = append(ArticleList, article)
}

func getNextID() int64 {
	atomic.AddInt64(&id, 1)
	return id
}
