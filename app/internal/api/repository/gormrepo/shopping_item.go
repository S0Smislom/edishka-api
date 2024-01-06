package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"
	"time"

	"gorm.io/gorm"
)

type ShoppingItemRepository struct {
	db *gorm.DB
}

func NewShoppingItemRepository(db *gorm.DB) *ShoppingItemRepository {
	return &ShoppingItemRepository{db: db}
}

func (r *ShoppingItemRepository) Create(data *model.CreateShoppingItem) (int, error) {
	now := time.Now().UTC()
	data.CreatedAt = now
	data.UpdatedAt = now
	res := r.db.Table("shopping_item").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *ShoppingItemRepository) GetById(id int) (*model.ShoppingItem, error) {
	dbModel := &model.ShoppingItem{}
	res := r.db.First(&dbModel, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Item not found"}
		}
		return nil, res.Error
	}
	return dbModel, nil
}

func (r *ShoppingItemRepository) Count(filters *model.ShoppingItemFilter) (int, error) {
	var total int64
	res := r.db.Table("shopping_item").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *ShoppingItemRepository) GetList(limit, offset int, filters *model.ShoppingItemFilter) ([]*model.ShoppingItem, error) {
	dbModels := []*model.ShoppingItem{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&dbModels)
	if res.Error != nil {
		return []*model.ShoppingItem{}, res.Error
	}
	return dbModels, nil
}

func (r *ShoppingItemRepository) Update(id int, data *model.UpdateShoppingItem) error {
	data.UpdatedAt = time.Now().UTC()
	return r.db.Model(&model.ShoppingItem{}).Where("id=?", id).Updates(data).Error
}

func (r *ShoppingItemRepository) Delete(id int) error {
	return r.db.Model(&model.ShoppingItem{}).Delete(&model.ShoppingItem{}, id).Error
}

func (r *ShoppingItemRepository) getFilterQuery(filters *model.ShoppingItemFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
