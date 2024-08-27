package server

import (
	"github.com/gin-gonic/gin"
	"my-template-with-go/internal/service"
)

func setupArticleRouter(
	router *gin.Engine,
	ctl service.IArticleCtl,
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
