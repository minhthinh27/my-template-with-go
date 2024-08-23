package biz

import "my-template-with-go/internal/model"

type IUserRedisRepo interface {
	GetAllUser() ([]*model.UserLogin, error)
	GetUserByID(id string) error
}
