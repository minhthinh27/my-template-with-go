package biz

import (
	"my-template-with-go/internal/data"
	"my-template-with-go/internal/model"
	"my-template-with-go/request"
	"my-template-with-go/response"
)

type IArticleUC interface {
	List() (interface{}, error)
	Detail(id uint) (interface{}, error)
	Create(jBody *request.ArticleCreateReq) error
	Update(id uint, jBody *request.ArticleUpdateReq) error
	Delete(jBody *request.ArticleDeleteReq) error
}

type articleUC struct {
	repo data.IArticleRepo
}

func (b *articleUC) List() (interface{}, error) {
	var (
		res []*response.ArticleListRes
	)

	// do something with login business
	articles, err := b.repo.List()
	if err != nil {
		return nil, err
	}

	if len(articles) > 0 {
		res = make([]*response.ArticleListRes, 0, len(articles))
		for _, a := range articles {
			temp := &response.ArticleListRes{}
			temp.SetAttributes(a)
			res = append(res, temp)
		}
	}

	return res, nil
}

func (b *articleUC) Detail(id uint) (interface{}, error) {
	var (
		res *response.ArticleDetailRes
	)

	// do something with login business
	article, err := b.repo.Detail(id)
	if err != nil {
		return nil, err
	}

	res = &response.ArticleDetailRes{}
	res.SetAttributes(article)

	return res, nil
}

func (b *articleUC) Create(jBody *request.ArticleCreateReq) error {
	// do something with login business
	articleEntity := model.ToArticleEntity(jBody.Author, jBody.Title)

	if err := b.repo.Create(articleEntity); err != nil {
		return err
	}
	return nil
}

func (b *articleUC) Update(id uint, jBody *request.ArticleUpdateReq) error {
	var (
		updateItems = map[string]interface{}{}
	)

	// do something with login business
	if jBody.Author != nil {
		updateItems["author"] = jBody.Author
	}

	if jBody.Title != nil {
		updateItems["title"] = jBody.Title
	}

	if err := b.repo.Update(id, updateItems); err != nil {
		return err
	}

	return nil
}

func (b *articleUC) Delete(jBody *request.ArticleDeleteReq) error {
	// do something with login business
	return b.repo.Delete(jBody.IDs)
}

func NewLinkUseCase(repo data.IArticleRepo) IArticleUC {
	return &articleUC{repo: repo}
}
