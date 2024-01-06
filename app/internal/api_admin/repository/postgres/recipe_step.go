package postgres

import (
	"database/sql"
	"fmt"
	"food/internal/api_admin/model"
	"food/pkg/exceptions"
	"strconv"
	"strings"
	"time"
)

type RecipeStepRepository struct {
	db *sql.DB
}

func NewRecipeStepRepository(db *sql.DB) *RecipeStepRepository {
	return &RecipeStepRepository{db: db}
}

func (r *RecipeStepRepository) Create(data *model.CreateRecipeStep) (int, error) {
	var id int
	row := r.db.QueryRow(
		"INSERT INTO recipe_step (title, description, ordering, recipe_id, created_by_id) values ($1, $2, $3, $4, $5) RETURNING id",
		data.Title,
		data.Description,
		data.Ordering,
		data.RecipeId,
		data.CreatedById,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecipeStepRepository) GetById(id int) (*model.RecipeStep, error) {
	p := &model.RecipeStep{}
	if err := r.db.QueryRow(
		"select id, title, description, ordering, recipe_id, created_at, updated_at, photo, created_by_id, updated_by_id from recipe_step where id=$1",
		id,
	).Scan(
		&p.Id,
		&p.Title,
		&p.Description,
		&p.Ordering,
		&p.RecipeId,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Photo,
		&p.CreatedById,
		&p.UpdatedById,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ObjectNotFoundError{Msg: "RecipeStep not found"}
		}
		return nil, err
	}
	return p, nil
}

func (r *RecipeStepRepository) GetList(limit, offset int, filters *model.RecipeStepFilter) ([]*model.RecipeStep, error) {
	tempQuery := "select id, title, description, ordering, recipe_id, created_at, updated_at, photo, created_by_id, updated_by_id from recipe_step"
	query, values, err := r.prepareFilters(tempQuery, filters)
	if err != nil {
		return nil, err
	}
	query += " limit " + strconv.Itoa(limit)
	query += " offset " + strconv.Itoa(offset)
	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recipeSteps := []*model.RecipeStep{}
	for rows.Next() {
		p := &model.RecipeStep{}
		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Description,
			&p.Ordering,
			&p.RecipeId,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Photo,
			&p.CreatedById,
			&p.UpdatedById,
		); err != nil {
			return nil, err
		}
		recipeSteps = append(recipeSteps, p)
	}
	return recipeSteps, nil
}

func (r *RecipeStepRepository) Count(filters *model.RecipeStepFilter) (int, error) {
	tempQuery := "select count(*) from recipe_step"
	query, values, err := r.prepareFilters(tempQuery, filters)
	if err != nil {
		return 0, err
	}
	fmt.Println(query, values)
	var total int
	if err := r.db.QueryRow(query, values...).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func (r *RecipeStepRepository) Update(id int, data *model.UpdateRecipeStep) error {
	var queryParams []string
	var queryValues []interface{}
	// TODO придумать более изящный способ обновления
	queryValues = append(queryValues, time.Now())
	queryParams = append(queryParams, "updated_at=$"+strconv.Itoa(len(queryValues)))
	queryValues = append(queryValues, *data.UpdatedById)
	queryParams = append(queryParams, "updated_by_id=$"+strconv.Itoa(len(queryValues)))
	if data.Title != nil {
		queryValues = append(queryValues, *data.Title)
		queryParams = append(queryParams, "title=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Description != nil {
		queryValues = append(queryValues, *data.Description)
		queryParams = append(queryParams, "description=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Ordering != nil {
		queryValues = append(queryValues, *data.Ordering)
		queryParams = append(queryParams, "ordering=$"+strconv.Itoa(len(queryValues)))
	}
	if len(queryParams) == 0 || len(queryValues) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "UPDATE recipe_step SET " + strings.Join(queryParams, ", ") + " where id=$" + strconv.Itoa(len(queryValues))
	fmt.Println(query)
	_, err := r.db.Exec(query, queryValues...)
	return err
}

func (r *RecipeStepRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from recipe_step where id=$1", id)
	return err
}

func (r *RecipeStepRepository) UpdatePhoto(id int, photo *string) error {
	_, err := r.db.Exec("update recipe_step set photo=$1 where id=$2", photo, id)
	return err
}

func (r *RecipeStepRepository) prepareFilters(query string, filters *model.RecipeStepFilter) (string, []interface{}, error) {
	var filterValues []interface{}
	var filterQuery []string
	// Recipe id
	filterValues = append(filterValues, filters.RecipeID)
	filterQuery = append(filterQuery, "recipe_id = $"+strconv.Itoa(len(filterValues)))
	// Title
	if filters.Title != nil {
		filterValues = append(filterValues, "%%"+*filters.Title+"%%")
		filterQuery = append(filterQuery, "LOWER(title) like LOWER($"+strconv.Itoa(len(filterValues))+")")
	}
	if len(filterValues) == 0 {
		return query, filterValues, nil
	}
	return query + " where " + strings.Join(filterQuery, " AND "), filterValues, nil
}
