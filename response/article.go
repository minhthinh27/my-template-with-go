package response

import "my-template-with-go/internal/entity"

type ArticleListRes struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (r *ArticleListRes) SetAttributes(item *entity.Article) {
	r.ID = item.ID
	r.Title = item.Title
	r.Author = item.Author
}

type ArticleDetailRes struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (r *ArticleDetailRes) SetAttributes(item *entity.Article) {
	r.Title = item.Title
	r.Author = item.Author
}
