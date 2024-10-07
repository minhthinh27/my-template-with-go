package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-template-with-go/bootstrap"
	"my-template-with-go/constant"
	"my-template-with-go/container"
	"my-template-with-go/helper/gderror"
	"my-template-with-go/internal/api"
	"my-template-with-go/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitRouter(container container.IContainerProvider, zap logger.ILogger, cf bootstrap.Config) *echo.Echo {
	router := echo.New()
	router.Use(middleware.Recover())
	router.HTTPErrorHandler = gderror.CustomHTTPErrorHandler
	cors(router)

	if cf.Server.Env.Mode != constant.EnvProduction {
		router.Use(middleware.Logger())
	}

	api.RegisterV1Router(router, container, cf)

	return router
}

func Start(engine *echo.Echo, zap logger.ILogger, cf bootstrap.Server) {
	var (
		sugar   = zap.GetZapLogger()
		timeOut = time.Duration(cf.Http.Timeout) * time.Second
		address = cf.Http.Address
	)

	// start and wait for stop signal
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", address),
		Handler:           engine,
		ReadHeaderTimeout: timeOut,
		ReadTimeout:       timeOut,
		WriteTimeout:      timeOut,
	}

	go func() {
		sugar.Infof("start server on: %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sugar.Fatal(err)
		}
	}()
	<-ctx.Done()
	stop()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatal(err)
	}
}

func cors(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Token"},
		AllowCredentials: true,
	}))
}
