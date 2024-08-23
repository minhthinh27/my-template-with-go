package data

import (
	"context"
	"fmt"
	"my-template-with-go/container"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/model"
)

const (
	UserLogin    = "user_loging_email_inbox"
	EmailAccount = "email_syncing_email_inbox"
)

type userRedisRepo struct {
	redis container.IRedisProvider
}

func (d *userRedisRepo) GetUserByID(id string) error {
	var (
		rd = d.redis.GetClient()
	)

	value, err := rd.Get(context.Background(), fmt.Sprintf("%s_%s", EmailAccount, id)).Result()
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}

func (d *userRedisRepo) GetAllUser() ([]*model.UserLogin, error) {
	var (
		rd = d.redis.GetClient()
		//res []*model.UserLogin
	)

	result, err := rd.HGetAll(context.Background(), UserLogin).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return nil, nil
}

func NewRedisRepo(redis container.IRedisProvider) biz.IUserRedisRepo {
	return &userRedisRepo{
		redis: redis,
	}
}
