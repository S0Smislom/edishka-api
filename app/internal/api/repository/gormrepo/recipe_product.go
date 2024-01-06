package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type RecipeProductRepository struct {
	db *gorm.DB
}

func NewRecipeProductRepository(db *gorm.DB) *RecipeProductRepository {
	return &RecipeProductRepository{db: db}
}

func (r *RecipeProductRepository) Create(data *model.CreateRecipeProduct) (int, error) {
	res := r.db.Table("recipe_product").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *RecipeProductRepository) GetById(id int) (*model.RecipeProduct, error) {
	RecipeProduct := &model.RecipeProduct{}
	res := r.db.First(&RecipeProduct, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "RecipeProduct not found"}
		}
		return nil, res.Error
	}
	return RecipeProduct, nil
}

func (r *RecipeProductRepository) GetList(limit, offset int, filters *model.RecipeProductFilter) ([]*model.RecipeProduct, error) {
	RecipeProduct := []*model.RecipeProduct{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&RecipeProduct)
	if res.Error != nil {
		return []*model.RecipeProduct{}, res.Error
	}
	return RecipeProduct, nil
}

func (r *RecipeProductRepository) Count(filters *model.RecipeProductFilter) (int, error) {
	var total int64
	res := r.db.Table("recipe_product").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *RecipeProductRepository) Update(id int, data *model.UpdateRecipeProduct) error {
	return r.db.Model(&model.RecipeProduct{}).Where("id=?", id).Updates(data).Error
}

func (r *RecipeProductRepository) Delete(id int) error {
	return r.db.Model(&model.RecipeProduct{}).Delete(&model.RecipeProduct{}, id).Error

}

func (r *RecipeProductRepository) getFilterQuery(filters *model.RecipeProductFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filters.RecipeId != nil {
			db = db.Where("recipe_id=?", filters.RecipeId)
		}
		if filters.ProductId != nil {
			db = db.Where("product_id=?", filters.ProductId)
		}
		return db
	}
}
