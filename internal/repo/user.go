package repo

import (
	"my-template-with-go/container"
	"my-template-with-go/internal/entity"
)

type IUserRepo interface {
	CheckUserExist(id uint) (bool, error)
}

type userRepo struct {
	db container.IDatabaseProvider
}

func (u *userRepo) CheckUserExist(id uint) (bool, error) {
	var (
		tx    = u.db.GetDBSlave()
		count int64
	)
	if err := tx.Model(&entity.User{}).
		Where("id = ?", id).
		Count(&count).
		Error; err != nil {
		return count > 0, err
	}

	return count > 0, nil
}

func NewUserRepo(db container.IDatabaseProvider) IUserRepo {
	return &userRepo{
		db: db,
	}
}
