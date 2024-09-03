package data

import (
	"errors"
	"gorm.io/gorm"
	"my-template-with-go/container"
	"my-template-with-go/internal/entity"
)

type IArticleRepo interface {
	List() ([]*entity.Article, error)
	Detail(id uint) (*entity.Article, error)
	Create(item *entity.Article) error
	Update(id uint, items map[string]interface{}) error
	Delete(ids []uint) error
}

type articleRepo struct {
	db container.IDatabaseProvider
}

func (d *articleRepo) List() ([]*entity.Article, error) {
	var (
		tx     = d.db.GetDBSlave()
		result []*entity.Article
	)

	if err := tx.Model(&entity.Article{}).
		Find(&result).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return result, nil
}

func (d *articleRepo) Detail(id uint) (*entity.Article, error) {
	var (
		tx     = d.db.GetDBSlave()
		result = &entity.Article{}
	)

	if err := tx.Model(&entity.Article{}).
		Where("id = ?", id).
		First(result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	return result, nil
}

func (d *articleRepo) Create(item *entity.Article) error {
	var (
		tx = d.db.GetDBMain()
	)

	if err := tx.Create(item).Error; err != nil {
		return err
	}

	return nil
}

func (d *articleRepo) Update(id uint, items map[string]interface{}) error {
	var (
		tx = d.db.GetDBMain()
	)

	if err := tx.Model(&entity.Article{}).
		Where("id = ?", id).
		Updates(items).Error; err != nil {
		return err
	}

	return nil
}

func (d *articleRepo) Delete(ids []uint) error {
	var (
		tx = d.db.GetDBMain()
	)

	return tx.Delete(&entity.Article{}, ids).Error
}

func NewArticleRepo(db container.IDatabaseProvider) IArticleRepo {
	return &articleRepo{
		db: db,
	}
}
