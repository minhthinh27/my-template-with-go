package main

import (
	"context"
	"log"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/helper/nl_cron"
	"my-template-with-go/internal/server"
	"my-template-with-go/logger"
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

	zap, err := logger.InitLogger(config)
	if err != nil {
		zap.GetZapLogger().Fatal(err)
	}

	provider, err := container.NewContainer(config, zap)
	if err != nil {
		zap.GetZapLogger().Fatal(err)
	}

	app, cleanup, err := server.NewCRONServer(provider, zap, config)
	if err != nil {
		cleanup()
		zap.GetZapLogger().Fatal(err)
	}

	run(app, config, zap)
}

func run(app nl_cron.ICronApp, cf bootstrap.Config, zap logger.ILogger) {
	app.Start()

	sugar := zap.GetZapLogger()
	sugar.Infof("[CronJOB] server starting on background...")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	sugar.Infof("[Shutdown Server]")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cf.Server.Http.Timeout)*time.Second)
	defer cancel()

	ctx.Done()
	sugar.Infof("[Exited]")
}
