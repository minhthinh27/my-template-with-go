package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/controller"
	"my-template-with-go/internal/repo"
	"net/http"
)

func Router(container container.IContainerProvider, cf bootstrap.Config) (*echo.Echo, error) {
	router := echo.New()
	router.Use(middleware.Recover())
	cors(router)

	if cf.Server.Env.Mode != "PRODUCTION" {
		router.Use(middleware.Logger())
	}

	userRepo := repo.NewUserRepo(container.DatabaseProvider())

	articleRepo := repo.NewArticleRepo(container.DatabaseProvider())
	articleUC := biz.NewArticleUseCase(articleRepo, userRepo)
	articleCtl := controller.NewArticleService(articleUC)

	setupArticleRouter(router, articleCtl)

	return router, nil
}

func cors(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Token"},
		AllowCredentials: true,
	}))
}
