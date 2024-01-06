package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"
	"log"

	"gorm.io/gorm"
)

type RecipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) Create(data *model.CreateRecipe) (int, error) {
	res := r.db.Table("recipe").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *RecipeRepository) GetById(id int) (*model.Recipe, error) {
	recipe := &model.Recipe{}
	res := r.db.Scopes(r.getSelectQuery, r.getQuery).Group("recipe.id").First(&recipe, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Recipe not found"}
		}
		return nil, res.Error
	}
	r.fetchProducts([]*model.Recipe{recipe})
	return recipe, nil
}

func (r *RecipeRepository) GetList(limit, offset int, filters *model.RecipeFilter) ([]*model.Recipe, error) {
	recipes := []*model.Recipe{}
	res := r.db.
		Scopes(r.getSelectQuery, r.getQuery, r.getFilterQuery(filters)).
		Group("recipe.id").
		Limit(limit).
		Offset(offset).
		Find(&recipes)
	if res.Error != nil {
		return recipes, res.Error
	}
	r.fetchProducts(recipes)
	return recipes, nil
}
func (r *RecipeRepository) Count(filters *model.RecipeFilter) (int, error) {
	var total int64

	res := r.db.
		Table("recipe").
		Scopes(r.getQuery, r.getFilterQuery(filters)).
		Group("recipe.id").
		Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}
func (r *RecipeRepository) Update(id int, data *model.UpdateRecipe) error {
	return r.db.Model(&model.Recipe{}).Where("id = ?", id).Updates(data).Error
}
func (r *RecipeRepository) Delete(id int) error {
	return r.db.Model(&model.Recipe{}).Delete(&model.Recipe{}, id).Error
}
func (r *RecipeRepository) GetOne(filters *model.RecipeFilter) (*model.Recipe, error) {
	recipe := &model.Recipe{}
	res := r.db.Scopes(r.getSelectQuery, r.getQuery, r.getFilterQuery(filters)).Group("recipe.id").First(&recipe)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Recipe not found"}
		}
		return nil, res.Error
	}

	r.fetchProducts([]*model.Recipe{recipe})

	return recipe, nil
}

func (r *RecipeRepository) getFilterQuery(filters *model.RecipeFilter) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		// Title
		if filters.Title != nil {
			query = query.Where("LOWER(recipe.title) like LOWER(?)", "%%"+*filters.Title+"%%")
		}
		// Slug
		if filters.Slug != nil {
			query = query.Where("LOWER(recipe.slug) like LOWER(?)", "%%"+*filters.Slug+"%%")
		}
		// CookingTime
		if filters.CookingTimeGTE != nil {
			query = query.Where("recipe.cooking_time >= ?", filters.CookingTimeGTE)
		}
		if filters.CookingTimeLTE != nil {
			query = query.Where("recipe.cooking_time <= ?", filters.CookingTimeLTE)
		}
		// Kitchen
		if filters.Kitchen != nil {
			query = query.Where("LOWER(recipe.kitchen) like LOWER(?)", "%%"+*filters.Kitchen+"%%")
		}
		// DifficultyLevel
		if filters.DifficultyLevel != nil {
			query = query.Where("recipe.difficulty_level like ?", filters.DifficultyLevel)
		}
		// Id
		if filters.Id != nil {
			query = query.Where("recipe.id=?", filters.Id)
		}
		// Published
		if filters.Published != nil {
			query = query.Where("recipe.published=?", filters.Published)
		}
		// Created by id
		if filters.CreatedById != nil {
			query = query.Where("recipe.created_by_id=?", filters.CreatedById)
		}
		return query
	}
}

func (r *RecipeRepository) getQuery(query *gorm.DB) *gorm.DB {
	return query.
		Joins("left join recipe_product on recipe.id=recipe_product.recipe_id").
		Joins("left join product on recipe_product.product_id = product.id")
}

func (r *RecipeRepository) getSelectQuery(query *gorm.DB) *gorm.DB {
	return query.Preload("Gallery").
		Select(`recipe.*,
			sum(coalesce(recipe_product.amount, 0) * coalesce(product.calories, 0)/100) as calories,
			sum(coalesce(recipe_product.amount, 0) * coalesce(product.fats, 0)/100) as fats,
			sum(coalesce(recipe_product.amount, 0) * coalesce(product.squirrels , 0)/100) as squirrels,
			sum(coalesce(recipe_product.amount, 0) * coalesce(product.carbohydrates , 0)/100) as carbohydrates`)
}

func (r *RecipeRepository) fetchProducts(recipes []*model.Recipe) {
	recipeIds := make([]int, len(recipes))
	for i, recipe := range recipes {
		recipeIds[i] = recipe.Id
	}
	products := []*model.RecipeProductMinimal{}
	res := r.db.Table("recipe").
		Select("product.*, sum(recipe_product.amount) as amount, recipe.id as recipe_id").
		Joins("inner join recipe_product on recipe.id=recipe_product.recipe_id").
		Joins("inner join product on recipe_product.product_id=product.id").
		Where("recipe.id IN (?)", recipeIds).
		Group("recipe.id, product.id").
		Find(&products)
	if res.Error != nil {
		log.Println("err", res.Error)
		return
	}
	recipeMap := map[int]*model.Recipe{}
	for _, recipe := range recipes {
		recipeMap[recipe.Id] = recipe
	}
	for _, product := range products {
		dbRecipe := recipeMap[product.RecipeId]
		dbRecipe.Products = append(dbRecipe.Products, product)
	}
}
