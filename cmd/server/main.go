package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/internal/server"
	"my-template-with-go/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configPath = "./configs"

func main() {
	config, err := bootstrap.InitConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	sugar, err := logger.InitLogger(config)
	if err != nil {
		sugar.GetZapLogger().Fatal(err)
	}

	provider, err := container.NewContainer(config, sugar)
	if err != nil {
		sugar.GetZapLogger().Fatal(err)
	}

	router, err := server.Router(provider, config)
	if err != nil {
		sugar.GetZapLogger().Fatal(err)
	}

	run(router, sugar, config.Server)
}

func run(engine *echo.Echo, zap logger.ILogger, cf bootstrap.Server) {
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
