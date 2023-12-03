package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type RecipeGalleryRepository struct {
	db *gorm.DB
}

func NewRecipeGalleryRepository(db *gorm.DB) *RecipeGalleryRepository {
	return &RecipeGalleryRepository{db: db}
}

func (r *RecipeGalleryRepository) Create(data *model.CreateRecipeGallery) (int, error) {
	res := r.db.Table("recipe_gallery").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *RecipeGalleryRepository) GetById(id int) (*model.RecipeGallery, error) {
	recipeGallery := &model.RecipeGallery{}
	res := r.db.First(&recipeGallery, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "RecipeGallery not found"}
		}
		return nil, res.Error
	}
	return recipeGallery, nil
}

func (r *RecipeGalleryRepository) GetList(limit, offset int, filters *model.RecipeGalleryFilter) ([]*model.RecipeGallery, error) {
	recipeGallery := []*model.RecipeGallery{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&recipeGallery)
	if res.Error != nil {
		return []*model.RecipeGallery{}, res.Error
	}
	return recipeGallery, nil
}

func (r *RecipeGalleryRepository) Count(filters *model.RecipeGalleryFilter) (int, error) {
	var total int64
	res := r.db.Table("recipe_gallery").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *RecipeGalleryRepository) Update(id int, data *model.UpdateRecipeGallery) error {
	return r.db.Model(&model.RecipeGallery{}).Where("id=?", id).Updates(data).Error
}

func (r *RecipeGalleryRepository) Delete(id int) error {
	return r.db.Model(&model.RecipeGallery{}).Delete(&model.RecipeGallery{}, id).Error

}

func (r *RecipeGalleryRepository) UpdatePhoto(id int, photo *string) error {
	return r.db.Model(&model.RecipeGallery{}).Where("id=?", id).Update("photo", photo).Error

}

func (r *RecipeGalleryRepository) getFilterQuery(filters *model.RecipeGalleryFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("recipe_id=?", filters.RecipeId)
	}
}
