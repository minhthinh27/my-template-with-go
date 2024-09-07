package controller

import (
	"github.com/labstack/echo/v4"
	"my-template-with-go/internal/biz"
	"my-template-with-go/request"
	"my-template-with-go/request/uri"
	"net/http"
	"strconv"
)

type IArticleCtl interface {
	List(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Create(ctx echo.Context) error
	Edit(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type articleCtl struct {
	uc biz.IArticleUC
}

func (s *articleCtl) List(ctx echo.Context) error {
	articles, err := s.uc.List(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, articles)
}

func (s *articleCtl) Detail(ctx echo.Context) error {
	idUri := &uri.IDUri{}
	if err := ctx.Bind(idUri); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	article, err := s.uc.Detail(idUri.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, article)
}

func (s *articleCtl) Create(ctx echo.Context) error {
	req := &request.ArticleCreateReq{}
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := s.uc.Create(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (s *articleCtl) Edit(ctx echo.Context) error {
	req := &request.ArticleUpdateReq{}
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	id := ctx.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := s.uc.Edit(uint(u64), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusAccepted, nil)
}

func (s *articleCtl) Delete(ctx echo.Context) error {
	req := &request.ArticleDeleteReq{}
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := s.uc.Delete(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusAccepted, nil)
}

func NewArticleService(uc biz.IArticleUC) IArticleCtl {
	return &articleCtl{
		uc: uc,
	}
}
