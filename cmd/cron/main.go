package main

import (
	"context"
	"log"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/helper/nl_cron"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/cron"
	"my-template-with-go/internal/data"
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

	app, cleanup, err := register(config, zap, provider)
	if cleanup != nil {
		defer cleanup()
	}

	if err != nil {
		zap.GetZapLogger().Panic(err)
	}

	run(app, config, zap)
}

func register(
	cf bootstrap.Config,
	zap logger.ILogger,
	provider container.IContainerProvider,
) (nl_cron.ICronApp, func(), error) {
	userRedisRepo := data.NewRedisRepo(provider.RedisProvider())
	mailUC := biz.NewMailBoxUC(zap, userRedisRepo)

	iMailBoxCron, cleanup, err := cron.NewMailBoxCron(cf, zap, mailUC)
	if err != nil {
		return nil, nil, err
	}

	iCronApp, cleanup1, err := server.NewCRONServer(iMailBoxCron)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	return iCronApp, func() {
		cleanup1()
		cleanup()
	}, nil
}

func run(app nl_cron.ICronApp, cf bootstrap.Config, zap logger.ILogger) {
	app.Start()

	sugar := zap.GetZapLogger()
	sugar.Infof("[CronJOB] server starting on background...")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	sugar.Infof("[Shutdown Server]")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cf.Server.GetHttp().GetTimeout())*time.Second)
	defer cancel()

	ctx.Done()
	sugar.Infof("[Exited]")
}
