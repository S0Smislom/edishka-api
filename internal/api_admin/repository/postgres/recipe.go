package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"food/internal/api_admin/model"
	"strconv"
	"strings"
	"time"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) Create(data *model.CreateRecipe) (int, error) {
	var id int
	row := r.db.QueryRow(
		"INSERT INTO recipe (title, slug, description, cooking_time, preparing_time, kitchen, difficulty_level) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		data.Title,
		data.Slug,
		data.Description,
		data.CookingTime,
		data.PreparingTime,
		data.Kitchen,
		data.DifficultyLevel,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecipeRepository) GetById(id int) (*model.Recipe, error) {
	query := r.querySelect() + " where r.id=$1 group by r.id"
	queryRow := r.db.QueryRow(query, id)
	p, err := r.scan(queryRow)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Recipe not found")
		}
		return nil, err
	}
	p.Products = r.getRecipeProduct(id)
	return p, nil
}

func (r *RecipeRepository) GetList(limit, offset int, filters *model.RecipeFilter) ([]*model.Recipe, error) {
	tempQuery := r.querySelect()
	query, values, err := r.prepareFilters(tempQuery, filters)
	if err != nil {
		return nil, err
	}
	query += " group by r.id"
	query += " limit " + strconv.Itoa(limit)
	query += " offset " + strconv.Itoa(offset)
	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recipes := []*model.Recipe{}
	for rows.Next() {
		p, err := r.scan(rows)
		if err != nil {
			continue
		}
		p.Products = r.getRecipeProduct(p.Id)
		recipes = append(recipes, p)
	}
	return recipes, nil
}

func (r *RecipeRepository) Count(filters *model.RecipeFilter) (int, error) {
	tempQuery := "select count(*) from recipe"
	query, values, err := r.prepareFilters(tempQuery, filters)
	if err != nil {
		return 0, err
	}
	fmt.Println(query)
	var total int
	if err := r.db.QueryRow(query, values...).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func (r *RecipeRepository) Update(id int, data *model.UpdateRecipe) error {
	var queryParams []string
	var queryValues []interface{}
	// TODO придумать более изящный способ обновления
	queryValues = append(queryValues, time.Now())
	queryParams = append(queryParams, "updated_at=$"+strconv.Itoa(len(queryValues)))
	if data.Title != nil {
		queryValues = append(queryValues, *data.Title)
		queryParams = append(queryParams, "title=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Slug != nil {
		queryValues = append(queryValues, *data.Slug)
		queryParams = append(queryParams, "slug=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Description != nil {
		queryValues = append(queryValues, *data.Description)
		queryParams = append(queryParams, "description=$"+strconv.Itoa(len(queryValues)))
	}
	if data.CookingTime != nil {
		queryValues = append(queryValues, *data.CookingTime)
		queryParams = append(queryParams, "cooking_time=$"+strconv.Itoa(len(queryValues)))
	}
	if data.PreparingTime != nil {
		queryValues = append(queryValues, *data.PreparingTime)
		queryParams = append(queryParams, "preparing_time=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Kitchen != nil {
		queryValues = append(queryValues, *data.Kitchen)
		queryParams = append(queryParams, "kitchen=$"+strconv.Itoa(len(queryValues)))
	}
	if data.DifficultyLevel != nil {
		queryValues = append(queryValues, *data.DifficultyLevel)
		queryParams = append(queryParams, "difficulty_level=$"+strconv.Itoa(len(queryValues)))
	}
	if len(queryParams) == 0 || len(queryValues) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "UPDATE recipe SET " + strings.Join(queryParams, ", ") + " where id=$" + strconv.Itoa(len(queryValues))
	fmt.Println(query)
	_, err := r.db.Exec(query, queryValues...)
	return err
}

func (r *RecipeRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from recipe where id=$1", id)
	return err
}

func (r *RecipeRepository) prepareFilters(query string, filters *model.RecipeFilter) (string, []interface{}, error) {
	var filterValues []interface{}
	var filterQuery []string
	// Title
	if filters.Title != nil {
		filterValues = append(filterValues, "%%"+*filters.Title+"%%")
		filterQuery = append(filterQuery, "LOWER(title) like LOWER($"+strconv.Itoa(len(filterValues))+")")
	}
	// Slug
	if filters.Slug != nil {
		filterValues = append(filterValues, "%%"+*filters.Slug+"%%")
		filterQuery = append(filterQuery, "LOWER(slug) like LOWER($"+strconv.Itoa(len(filterValues))+")")
	}
	// CookingTime
	if filters.CookingTimeGTE != nil {
		filterValues = append(filterValues, *filters.CookingTimeGTE)
		filterQuery = append(filterQuery, "cooking_time >= $"+strconv.Itoa(len(filterValues)))
	}
	if filters.CookingTimeLTE != nil {
		filterValues = append(filterValues, *filters.CookingTimeLTE)
		filterQuery = append(filterQuery, "cooking_time <= $"+strconv.Itoa(len(filterValues)))
	}
	// Kitchen
	if filters.Kitchen != nil {
		filterValues = append(filterValues, "%%"+*filters.Kitchen+"%%")
		filterQuery = append(filterQuery, "LOWER(kitchen) like LOWER($"+strconv.Itoa(len(filterValues))+")")
	}
	// DifficultyLevel
	if filters.DifficultyLevel != nil {
		filterValues = append(filterValues, *filters.DifficultyLevel)
		filterQuery = append(filterQuery, "difficulty_level like $"+strconv.Itoa(len(filterValues)))
	}
	if len(filterValues) == 0 {
		return query, filterValues, nil
	}
	return query + " where " + strings.Join(filterQuery, " AND "), filterValues, nil
}

func (r *RecipeRepository) scan(row interface {
	Scan(dest ...any) error
}) (*model.Recipe, error) {
	p := &model.Recipe{}
	if err := row.Scan(
		&p.Id,
		&p.Title,
		&p.Slug,
		&p.Description,
		&p.CookingTime,
		&p.PreparingTime,
		&p.Kitchen,
		&p.DifficultyLevel,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Photo,
		&p.Published,
		&p.Calories,
		&p.Squirrels,
		&p.Fats,
		&p.Carbohydrates,
	); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *RecipeRepository) querySelect() string {
	query := `
		select 
			r.id,
			r.title,
			r.slug,
			r.description,
			r.cooking_time,
			r.preparing_time,
			r.kitchen,
			r.difficulty_level,
			r.created_at,
			r.updated_at,
			r.photo,
			r.published,
			sum(coalesce(sp.amount, 0) * coalesce(p.calories, 0)/100),
			sum(coalesce(sp.amount, 0) * coalesce(p.fats, 0)/100),
			sum(coalesce(sp.amount, 0) * coalesce(p.squirrels , 0)/100),
			sum(coalesce(sp.amount, 0) * coalesce(p.carbohydrates , 0)/100)
		from recipe as r
		left join recipe_step rs
			left join step_product sp 
				left join product p 
				on sp.product_id = p.id
			on rs.id = sp.recipe_step_id
		on r.id=rs.recipe_id
	`
	return query
}

func (r *RecipeRepository) getRecipeProduct(recipeId int) []*model.RecipeProduct {
	query := `
	select 
		p.id,
		p.title,
		p.slug,
		p.description,
		p.calories,
		p.squirrels,
		p.fats,
		p.carbohydrates,
		p.created_at,
		p.updated_at,
		p.photo,
		sum(sp.amount) 
	from recipe r 
	left join recipe_step rs
		left join step_product sp 
			left join product p 
			on sp.product_id = p.id
		on rs.id = sp.recipe_step_id
	on r.id=rs.recipe_id
	where r.id=$1
	group by r.id, p.id
	`
	products := []*model.RecipeProduct{}
	rows, err := r.db.Query(query, recipeId)
	if err != nil {
		return products
	}
	defer rows.Close()
	for rows.Next() {
		p := &model.RecipeProduct{}
		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Slug,
			&p.Description,
			&p.Calories,
			&p.Squirrels,
			&p.Fats,
			&p.Carbohydrates,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Photo,
			&p.Amount,
		); err != nil {
			continue
		}
		products = append(products, p)
	}
	return products
}
