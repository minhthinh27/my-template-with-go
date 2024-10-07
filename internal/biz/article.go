package biz

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"my-template-with-go/internal/model"
	"my-template-with-go/internal/repo"
	"my-template-with-go/request"
	"my-template-with-go/response"
)

type IArticleUC interface {
	Sync(ctx context.Context) error

	List(ctx echo.Context) (interface{}, error)
	Detail(id uint) (interface{}, error)
	Create(jBody *request.ArticleCreateReq) error
	Edit(id uint, jBody *request.ArticleUpdateReq) error
	Delete(jBody *request.ArticleDeleteReq) error
}

type articleUC struct {
	articleRepo repo.IArticleRepo
}

func (b *articleUC) Sync(ctx context.Context) error {
	fmt.Println("syncing")
	return nil
}

func (b *articleUC) List(ctx echo.Context) (interface{}, error) {
	var (
		res []*response.ArticleListRes
	)

	// do something with login business
	articles, err := b.articleRepo.List()
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
	article, err := b.articleRepo.Detail(id)
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

	if err := b.articleRepo.Create(articleEntity); err != nil {
		return err
	}
	return nil
}

func (b *articleUC) Edit(id uint, jBody *request.ArticleUpdateReq) error {
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

	if err := b.articleRepo.Update(id, updateItems); err != nil {
		return err
	}

	return nil
}

func (b *articleUC) Delete(jBody *request.ArticleDeleteReq) error {
	// do something with login business
	return b.articleRepo.Delete(jBody.IDs)
}

func NewArticleUseCase(articleRepo repo.IArticleRepo) IArticleUC {
	return &articleUC{
		articleRepo: articleRepo,
	}
}
