package model

import "my-template-with-go/internal/entity"

func ToArticleEntity(author, title string) *entity.Article {
	return &entity.Article{
		Author: author,
		Title:  title,
	}
}
