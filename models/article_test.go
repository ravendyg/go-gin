package models

import (
	"testing"
)

func TestShowIndexPage(t *testing.T) {
	alist := ShowIndexPage()

	if len(alist) != len(ArticleList) {
		t.Fail()
	}

	for i, v := range alist {
		article := ArticleList[i]
		if v.Content != article.Content ||
			v.ID != article.ID ||
			v.Title != article.Title {
			t.Fail()
		}
	}
}
