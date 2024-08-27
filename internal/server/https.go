package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/data"
	"my-template-with-go/internal/service"
)

func Router(container container.IContainerProvider, cf bootstrap.Config) (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	Cors(router)

	articleRepo := data.NewArticleRepo(container.DatabaseProvider())
	articleUC := biz.NewArticleUseCase(articleRepo)
	articleCtl := service.NewArticleService(articleUC)

	setupArticleRouter(router, articleCtl)

	return router, nil
}

func Cors(e *gin.Engine) {
	e.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length",
			"Content-Type", "Token"},
		AllowCredentials: true,
	}))
}
