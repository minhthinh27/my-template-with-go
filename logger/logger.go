package logger

import (
	"go.uber.org/zap"
	"log"
	"my-template-with-go/bootstrap"
)

const (
	EnvProduction = "production"
)

type ILogger interface {
	GetZapLogger() *zap.SugaredLogger
}

type logger struct {
	Zap *zap.SugaredLogger
}

func InitLogger(config bootstrap.Config) ILogger {
	zapLogger, err := build(config)
	defer zapLogger.Sync()

	if err != nil {
		log.Fatalf("Error init zap logger: %v", err)
	}

	return &logger{Zap: zapLogger.Sugar()}
}

func build(config bootstrap.Config) (*zap.Logger, error) {
	var (
		env = config.Server.Env
	)

	zapLogger, err := zap.NewDevelopment()
	if env.Mode == EnvProduction {
		zapLogger, err = zap.NewProduction()
	}

	return zapLogger, err
}

func (l logger) GetZapLogger() *zap.SugaredLogger {
	return l.Zap
}
