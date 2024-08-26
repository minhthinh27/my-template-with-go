package container

import (
	"go.uber.org/zap"
	"my-template-with-go/bootstrap"
	"my-template-with-go/logger"
)

type IContainerProvider interface {
	RedisProvider() IRedisProvider
	DatabaseProvider() IDatabaseProvider
}

type containerProvider struct {
	redisProvider    IRedisProvider
	databaseProvider IDatabaseProvider
}

func NewContainer(cf bootstrap.Config, zap logger.ILogger) (IContainerProvider, error) {
	var (
		sugar    = zap.GetZapLogger()
		provider = &containerProvider{}
	)

	provider.databaseProvider = buildDatabase(cf, sugar)
	provider.redisProvider = buildRedis(cf, sugar)

	return provider, nil
}

func buildDatabase(cf bootstrap.Config, sugar *zap.SugaredLogger) IDatabaseProvider {
	database, cleanup, err := NewDatabase(cf.Database, sugar)
	if err != nil {
		cleanup()
		sugar.Panic("init database err")
	}

	return database
}

func buildRedis(cf bootstrap.Config, sugar *zap.SugaredLogger) IRedisProvider {
	redis, cleanup, err := NewRedis(cf.Cache, sugar)
	if err != nil {
		cleanup()
		sugar.Panic("init redis err")
	}

	return redis
}

func (p containerProvider) RedisProvider() IRedisProvider {
	return p.redisProvider
}

func (p containerProvider) DatabaseProvider() IDatabaseProvider {
	return p.databaseProvider
}
