package container

import (
	"my-template-with-go/bootstrap"
	"my-template-with-go/logger"
)

type IContainerProvider interface {
	RedisProvider() IRedisProvider
}

type containerProvider struct {
	redisProvider IRedisProvider
}

func NewContainer(
	cf bootstrap.Config,
	zap logger.ILogger,
) (IContainerProvider, error) {
	var (
		sugar    = zap.GetZapLogger()
		provider = &containerProvider{}
	)

	redis, cleanup, err := NewRedis(cf.Cache, zap)
	if err != nil {
		cleanup()
		sugar.Panic("init data err")
	}
	provider.redisProvider = redis

	return provider, nil
}

func (p containerProvider) RedisProvider() IRedisProvider {
	return p.redisProvider
}
