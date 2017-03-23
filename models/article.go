package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []article{
	{ID: 1, Title: "Art 1", Content: "Content 1"},
	{ID: 2, Title: "Art 2", Content: "Content 2"},
}

func GetAllArticles() []article {
	return ArticleList
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
