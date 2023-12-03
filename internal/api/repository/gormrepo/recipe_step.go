package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type RecipeStepRepository struct {
	db *gorm.DB
}

func NewRecipeStepRepository(db *gorm.DB) *RecipeStepRepository {
	return &RecipeStepRepository{db: db}
}

func (r *RecipeStepRepository) Create(data *model.CreateRecipeStep) (int, error) {
	res := r.db.Table("recipe_step").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *RecipeStepRepository) GetById(id int) (*model.RecipeStep, error) {
	recipeStep := &model.RecipeStep{}
	res := r.db.First(&recipeStep, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "RecipeStep not found"}
		}
		return nil, res.Error
	}
	return recipeStep, nil
}

func (r *RecipeStepRepository) GetList(limit, offset int, filters *model.RecipeStepFilter) ([]*model.RecipeStep, error) {
	recipeSteps := []*model.RecipeStep{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&recipeSteps)
	if res.Error != nil {
		return []*model.RecipeStep{}, res.Error
	}
	return recipeSteps, nil
}

func (r *RecipeStepRepository) Count(filters *model.RecipeStepFilter) (int, error) {
	var total int64
	res := r.db.Table("recipe_step").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *RecipeStepRepository) Update(id int, data *model.UpdateRecipeStep) error {
	return r.db.Model(&model.RecipeStep{}).Where("id=?", id).Updates(data).Error
}

func (r *RecipeStepRepository) Delete(id int) error {
	return r.db.Model(&model.RecipeStep{}).Delete(&model.RecipeStep{}, id).Error
}

func (r *RecipeStepRepository) UpdatePhoto(id int, photo *string) error {
	return r.db.Model(&model.RecipeStep{}).Where("id=?", id).Update("photo", photo).Error
}

func (r *RecipeStepRepository) getFilterQuery(filters *model.RecipeStepFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filters.Title != nil {
			db = db.Where("LOWER(title) like LOWER(?)", "%%"+*filters.Title+"%%")
		}
		return db.Where("recipe_id=?", filters.RecipeID)
	}
}
