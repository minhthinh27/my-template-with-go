package logger

import (
	"go.uber.org/zap"
	"my-template-with-go/bootstrap"
)

type ILogger interface {
	GetZapLogger() *zap.SugaredLogger
}

type logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(sugar *zap.SugaredLogger) ILogger {
	return &logger{Zap: sugar}
}

func InitLogger(config bootstrap.Config) (ILogger, error) {
	zapLogger, err := build(config)
	defer zapLogger.Sync()

	if err != nil {
		return nil, err
	}

	log := NewLogger(zapLogger.Sugar())
	log.GetZapLogger().Info("Success zap logger configuration")

	return log, nil
}

func build(config bootstrap.Config) (*zap.Logger, error) {
	var (
		env = config.Server.Env
	)

	zapLogger, err := zap.NewDevelopment()
	if env.Mode == "PRODUCTION" {
		zapLogger, err = zap.NewProduction()
	}

	return zapLogger, err
}

func (l logger) GetZapLogger() *zap.SugaredLogger {
	return l.Zap
}
