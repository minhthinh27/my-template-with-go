package service

import (
	"github.com/gin-gonic/gin"
	"my-template-with-go/internal/biz"
	"my-template-with-go/request"
	"my-template-with-go/request/uri"
	"net/http"
)

type IArticleCtl interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Create(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type articleCtl struct {
	uc biz.IArticleUC
}

func (s *articleCtl) List(ctx *gin.Context) {
	articles, err := s.uc.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, articles)
	return
}

func (s *articleCtl) Detail(ctx *gin.Context) {
	idUri := &uri.IDUri{}
	if err := ctx.Bind(idUri); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	article, err := s.uc.Detail(idUri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, article)
	return
}

func (s *articleCtl) Create(ctx *gin.Context) {
	req := &request.ArticleCreateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.uc.Create(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, nil)
	return
}

func (s *articleCtl) Edit(ctx *gin.Context) {
	idUri := &uri.IDUri{}
	if err := ctx.Bind(idUri); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	req := &request.ArticleUpdateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.uc.Edit(idUri.ID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
	return
}

func (s *articleCtl) Delete(ctx *gin.Context) {
	req := &request.ArticleDeleteReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.uc.Delete(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
	return
}

func NewArticleService(uc biz.IArticleUC) IArticleCtl {
	return &articleCtl{
		uc: uc,
	}
}
