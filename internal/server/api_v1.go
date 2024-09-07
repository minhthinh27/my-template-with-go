package server

import (
	"github.com/labstack/echo/v4"
	"my-template-with-go/internal/controller"
)

func setupArticleRouter(
	router *echo.Echo,
	ctl controller.IArticleCtl,
) {
	articleGroup := router.Group("/api/v1/article")
	{
		articleGroup.GET("/list", ctl.List)
		articleGroup.GET("/:id/detail", ctl.Detail)
		articleGroup.POST("", ctl.Create)
		articleGroup.PUT("/:id/edit", ctl.Edit)
		articleGroup.DELETE("/delete", ctl.Delete)
	}
}
