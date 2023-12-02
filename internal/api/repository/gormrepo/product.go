package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(data *model.CreateProduct) (int, error) {
	res := r.db.Table("product").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *ProductRepository) GetById(id int) (*model.Product, error) {
	product := &model.Product{}
	res := r.db.First(&product, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Product not found"}
		}
		return nil, res.Error
	}
	return product, nil
}

func (r *ProductRepository) Count(filters *model.ProductFilter) (int, error) {
	var total int64
	res := r.db.Table("product").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *ProductRepository) UpdatePhoto(id int, photo *string) error {
	return r.db.Model(&model.Product{}).Where("id=?", id).Update("photo", photo).Error
}

func (r *ProductRepository) Delete(id int) error {
	return r.db.Model(&model.Product{}).Delete(&model.Product{}, id).Error
}

func (r *ProductRepository) Update(id int, data *model.UpdateProduct) error {
	return r.db.Model(&model.Product{}).Where("id=?", id).Updates(data).Error
}

func (r *ProductRepository) GetList(limit, offset int, filters *model.ProductFilter) ([]*model.Product, error) {
	products := []*model.Product{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&products)
	if res.Error != nil {
		return []*model.Product{}, res.Error
	}
	return products, nil
}

func (r *ProductRepository) getFilterQuery(filters *model.ProductFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Title
		if filters.Title != nil {
			db = db.Where("LOWER(title) like LOWER(?)", "%%"+*filters.Title+"%%")
		}
		// Slug
		if filters.Slug != nil {
			db = db.Where("LOWER(slug) like LOWER(?)", "%%"+*filters.Slug+"%%")
		}
		// Calories
		if filters.CaloriesGTE != nil {
			db = db.Where("calories >= ?", filters.CaloriesGTE)
		}
		if filters.CaloriesLTE != nil {
			db = db.Where("calories <= ?", filters.CaloriesLTE)
		}
		// Squirrels
		if filters.SquirrelsGTE != nil {
			db = db.Where("squirrels >= ?", filters.SquirrelsGTE)
		}
		if filters.SquirrelsLTE != nil {
			db = db.Where("squirrels <= ?", filters.SquirrelsLTE)
		}
		// Fats
		if filters.FatsGTE != nil {
			db = db.Where("fats >= ?", filters.FatsGTE)
		}
		if filters.FatsLTE != nil {
			db = db.Where("fats <= ?", filters.FatsLTE)
		}
		// Carbohydrates
		if filters.CarbohydratesGTE != nil {
			db = db.Where("carbohydrates >= ?", filters.CarbohydratesGTE)
		}
		if filters.CarbohydratesLTE != nil {
			db = db.Where("carbohydrates <= ?", filters.CarbohydratesLTE)
		}
		return db
	}
}
