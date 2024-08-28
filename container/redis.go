package container

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"my-template-with-go/bootstrap"
)

type IRedisProvider interface {
	GetClient() *redis.Client
}

type redisProvider struct {
	redisClient *redis.Client
}

func NewRedis(config bootstrap.Cache, sugar *zap.SugaredLogger) (IRedisProvider, func(), error) {
	var (
		data    = &redisProvider{}
		cfRedis = config.Redis
	)

	cleanup := func() {
		if data != nil && data.GetClient() != nil {
			data.GetClient().Close()
		}
		sugar.Info("closing the data resources")
	}

	if cfRedis.Host != "" {
		redisClient, err := connectRedis(cfRedis)
		if err != nil {
			return data, cleanup, err
		} else {
			data.redisClient = redisClient
		}
	}

	return data, cleanup, nil
}

func connectRedis(cfRedis bootstrap.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfRedis.Host + ":" + cfRedis.Port,
		Password: cfRedis.Password,
		DB:       cfRedis.Db,
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
