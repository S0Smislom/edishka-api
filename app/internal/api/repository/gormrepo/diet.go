package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type DietRepository struct {
	db *gorm.DB
}

func NewDietRepository(db *gorm.DB) *DietRepository {
	return &DietRepository{db: db}
}

func (r *DietRepository) Create(data *model.CreateDiet) (int, error) {
	res := r.db.Table("diet").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *DietRepository) GetById(id int) (*model.Diet, error) {
	diet := &model.Diet{}
	res := r.db.Preload("Items").First(&diet, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Diet not found"}
		}
		return nil, res.Error
	}
	return diet, nil
}

func (r *DietRepository) Count(filters *model.DietFilter) (int, error) {
	var total int64
	res := r.db.Table("diet").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *DietRepository) GetList(limit, offset int, filters *model.DietFilter) ([]*model.Diet, error) {
	diets := []*model.Diet{}
	res := r.db.Preload("Items").Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&diets)
	if res.Error != nil {
		return []*model.Diet{}, res.Error
	}
	return diets, nil
}

func (r *DietRepository) Delete(id int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.DietItem{}).Where("diet_id=?", id).Delete(&model.DietItem{}).Error; err != nil {
			return err
		}
		return tx.Model(&model.Diet{}).Where("id=?", id).Delete(&model.Diet{}).Error
	})
}

func (r *DietRepository) Update(id int, data *model.UpdateDiet) error {
	return r.db.Model(&model.Diet{}).Where("id=?", id).Updates(data).Error
}

func (r *DietRepository) getFilterQuery(filters *model.DietFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filters.UserId != 0 {
			db = db.Where("user_id = ?", filters.UserId)
		}
		if filters.Title != nil {
			db = db.Where("LOWER(title) like LOWER(?)", "%%"+*filters.Title+"%%")
		}
		return db
	}
}
