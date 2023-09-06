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

type StepProductRepository struct {
	db *sql.DB
}

func NewStepProductRepository(db *sql.DB) *StepProductRepository {
	return &StepProductRepository{db: db}
}

func (r *StepProductRepository) Create(data *model.CreateStepProduct) (int, error) {
	var id int
	row := r.db.QueryRow(
		"INSERT INTO step_product (recipe_step_id, product_id, amount) values ($1, $2, $3) RETURNING id",
		data.RecipeStepId,
		data.ProductId,
		data.Amount,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StepProductRepository) GetById(id int) (*model.StepProduct, error) {
	p := &model.StepProduct{}
	if err := r.db.QueryRow(
		"select id, recipe_step_id, product_id, amount, created_at, updated_at from step_product where id=$1",
		id,
	).Scan(
		&p.Id,
		&p.RecipeStepId,
		&p.ProductId,
		&p.Amount,
		&p.CreatedAt,
		&p.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("StepProduct not found")
		}
		return nil, err
	}
	return p, nil
}

func (r *StepProductRepository) GetList(limit, offset int, filters *model.StepProductFilter) ([]*model.StepProduct, error) {
	tempQuery := "select id, recipe_step_id, product_id, amount, created_at, updated_at from step_product"
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
	StepProducts := []*model.StepProduct{}
	for rows.Next() {
		p := &model.StepProduct{}
		if err := rows.Scan(
			&p.Id,
			&p.RecipeStepId,
			&p.ProductId,
			&p.Amount,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		StepProducts = append(StepProducts, p)
	}
	return StepProducts, nil
}

func (r *StepProductRepository) Count(filters *model.StepProductFilter) (int, error) {
	tempQuery := "select count(*) from step_product"
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

func (r *StepProductRepository) Update(id int, data *model.UpdateStepProduct) error {
	var queryParams []string
	var queryValues []interface{}
	queryValues = append(queryValues, time.Now())
	queryParams = append(queryParams, "updated_at=$"+strconv.Itoa(len(queryValues)))
	if data.Amount != nil {
		queryValues = append(queryValues, *data.Amount)
		queryParams = append(queryParams, "amount=$"+strconv.Itoa(len(queryValues)))
	}
	if len(queryParams) == 0 || len(queryValues) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "UPDATE step_product SET " + strings.Join(queryParams, ", ") + " where id=$" + strconv.Itoa(len(queryValues))
	fmt.Println(query)
	_, err := r.db.Exec(query, queryValues...)
	return err
}

func (r *StepProductRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from step_product where id=$1", id)
	return err
}

func (r *StepProductRepository) prepareFilters(query string, filters *model.StepProductFilter) (string, []interface{}, error) {
	var filterValues []interface{}
	var filterQuery []string
	// RecipeStepId
	if filters.RecipeStepId != nil {
		filterValues = append(filterValues, *filters.RecipeStepId)
		filterQuery = append(filterQuery, "recipe_step_id=$"+strconv.Itoa(len(filterValues)))
	}
	// ProductId
	if filters.ProductId != nil {
		filterValues = append(filterValues, *filters.ProductId)
		filterQuery = append(filterQuery, "product_id=$"+strconv.Itoa(len(filterValues)))
	}
	if len(filterValues) == 0 {
		return query, filterValues, nil
	}
	return query + " where " + strings.Join(filterQuery, " AND "), filterValues, nil
}
