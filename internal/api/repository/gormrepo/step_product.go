package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type StepProductRepository struct {
	db *gorm.DB
}

func NewStepProductRepository(db *gorm.DB) *StepProductRepository {
	return &StepProductRepository{db: db}
}

func (r *StepProductRepository) Create(data *model.CreateStepProduct) (int, error) {
	res := r.db.Table("step_product").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *StepProductRepository) GetById(id int) (*model.StepProduct, error) {
	StepProduct := &model.StepProduct{}
	res := r.db.First(&StepProduct, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "StepProduct not found"}
		}
		return nil, res.Error
	}
	return StepProduct, nil
}

func (r *StepProductRepository) GetList(limit, offset int, filters *model.StepProductFilter) ([]*model.StepProduct, error) {
	stepProduct := []*model.StepProduct{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&stepProduct)
	if res.Error != nil {
		return []*model.StepProduct{}, res.Error
	}
	return stepProduct, nil
}

func (r *StepProductRepository) Count(filters *model.StepProductFilter) (int, error) {
	var total int64
	res := r.db.Table("step_product").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *StepProductRepository) Update(id int, data *model.UpdateStepProduct) error {
	return r.db.Model(&model.StepProduct{}).Where("id=?", id).Updates(data).Error
}

func (r *StepProductRepository) Delete(id int) error {
	return r.db.Model(&model.StepProduct{}).Delete(&model.StepProduct{}, id).Error

}

func (r *StepProductRepository) getFilterQuery(filters *model.StepProductFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filters.RecipeStepId != nil {
			db = db.Where("recipe_step_id=?", filters.RecipeStepId)
		}
		if filters.ProductId != nil {
			db = db.Where("product_id=?", filters.ProductId)
		}
		return db
	}
}
