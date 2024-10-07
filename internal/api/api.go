package api

import (
	"github.com/labstack/echo/v4"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/controller"
	"my-template-with-go/internal/repo"
)

func RegisterV1Router(router *echo.Echo, provider container.IContainerProvider, cf bootstrap.Config) {
	articleRepo := repo.NewArticleRepo(provider.DatabaseProvider())
	articleUC := biz.NewArticleUseCase(articleRepo)
	articleCtl := controller.NewArticleCtl(articleUC)

	registerArticleRouter(router, articleCtl)
}

func registerArticleRouter(router *echo.Echo, ctl controller.IArticleCtl) {
	articleGroup := router.Group("/api/v1/article")
	{
		articleGroup.GET("/list", ctl.List)
		articleGroup.GET("/:id/detail", ctl.Detail)
		articleGroup.POST("", ctl.Create)
		articleGroup.PUT("/:id/edit", ctl.Edit)
		articleGroup.DELETE("/delete", ctl.Delete)
	}
}
