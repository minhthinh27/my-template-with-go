package container

import (
	"context"
	"github.com/redis/go-redis/v9"
	"my-template-with-go/bootstrap"
	"my-template-with-go/logger"
)

type IRedisProvider interface {
	GetClient() *redis.Client
}

type redisProvider struct {
	redisClient *redis.Client
}

func NewRedis(
	config bootstrap.Cache,
	zap logger.ILogger,
) (IRedisProvider, func(), error) {
	var (
		data    = &redisProvider{}
		sugar   = zap.GetZapLogger()
		cfRedis = config.Redis
	)

	cleanup := func() {
		if data != nil && data.GetClient() != nil {
			data.GetClient().Close()
		}

		sugar.Info("closing the data resources")
	}

	if cfRedis.GetHost() != "" {
		redisClient, err := data.connectRedis(cfRedis)
		if err != nil {
			return data, cleanup, err
		} else {
			data.redisClient = redisClient
		}
	}

	return data, cleanup, nil
}

func (p *redisProvider) connectRedis(cfRedis bootstrap.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfRedis.GetHost() + ":" + cfRedis.GetPort(),
		Password: cfRedis.GetPassword(),
		DB:       cfRedis.GetDB(),
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

func (p *redisProvider) GetClient() *redis.Client {
	return p.redisClient
}
