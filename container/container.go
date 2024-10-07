package container

import (
	"go.uber.org/zap"
	"my-template-with-go/bootstrap"
	"my-template-with-go/logger"
)

type IContainerProvider interface {
	RedisProvider() IRedisProvider
	DatabaseProvider() IDataBaseProvider
}

type containerProvider struct {
	redisProvider    IRedisProvider
	databaseProvider IDataBaseProvider
}

func NewContainer(cf bootstrap.Config, zap logger.ILogger) IContainerProvider {
	var (
		sugar    = zap.GetZapLogger()
		provider = &containerProvider{}
	)

	provider.databaseProvider = buildDatabase(cf, sugar)
	provider.redisProvider = buildRedis(cf, sugar)

	return provider
}

func buildDatabase(cf bootstrap.Config, sugar *zap.SugaredLogger) IDataBaseProvider {
	database, cleanup, err := NewDatabase(cf.Database, sugar)
	if err != nil {
		cleanup()
		sugar.Fatalf("Error init database: %v", err)
	}

	return database
}

func buildRedis(cf bootstrap.Config, sugar *zap.SugaredLogger) IRedisProvider {
	redis, cleanup, err := NewRedis(cf.Cache, sugar)
	if err != nil {
		cleanup()
		sugar.Fatalf("Error init redis: %v", err)
	}

	return redis
}

func (p containerProvider) RedisProvider() IRedisProvider {
	return p.redisProvider
}

func (p containerProvider) DatabaseProvider() IDataBaseProvider {
	return p.databaseProvider
}
