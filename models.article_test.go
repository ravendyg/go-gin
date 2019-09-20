package main

import (
	"testing"
	"web-server/models"
)

func TestShowIndexPage(t *testing.T) {
	alist := models.ShowIndexPage()

	if len(alist) != len(models.ArticleList) {
		t.Fail()
	}

	for i, v := range alist {
		article := models.ArticleList[i]
		if v.Content != article.Content ||
			v.ID != article.ID ||
			v.Title != article.Title {
			t.Fail()
		}
	}
}
