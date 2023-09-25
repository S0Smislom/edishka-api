package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"food/internal/api/model"
	"strconv"
	"strings"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(data *model.CreateProduct) (int, error) {
	var id int
	row := r.db.QueryRow(
		"INSERT INTO product (title, slug, description, calories, squirrels, fats, carbohydrates, created_by_id) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		data.Title,
		data.Slug,
		data.Description,
		data.Calories,
		data.Squirrels,
		data.Fats,
		data.Carbohydrates,
		data.CreatedById,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) GetById(id int) (*model.Product, error) {
	p := &model.Product{}
	if err := r.db.QueryRow(
		"select id, title, slug, description, calories, squirrels, fats, carbohydrates, photo, suggested_by_user, created_by_id from product where id=$1",
		id,
	).Scan(
		&p.Id,
		&p.Title,
		&p.Slug,
		&p.Description,
		&p.Calories,
		&p.Squirrels,
		&p.Fats,
		&p.Carbohydrates,
		&p.Photo,
		&p.SuggestedByUser,
		&p.CreatedById,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Product not found")
		}
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) Count(filters *model.ProductFilter) (int, error) {
	tempQuery := "select count(*) from product"
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

func (r *ProductRepository) UpdatePhoto(id int, photo *string) error {
	_, err := r.db.Exec("update product set photo=$1 where id=$2", photo, id)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from product where id=$1", id)
	return err
}

func (r *ProductRepository) Update(id int, data *model.UpdateProduct) error {
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
	if data.Slug != nil {
		queryValues = append(queryValues, *data.Slug)
		queryParams = append(queryParams, "slug=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Description != nil {
		queryValues = append(queryValues, *data.Description)
		queryParams = append(queryParams, "description=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Calories != nil {
		queryValues = append(queryValues, *data.Calories)
		queryParams = append(queryParams, "calories=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Squirrels != nil {
		queryValues = append(queryValues, *data.Squirrels)
		queryParams = append(queryParams, "squirrels=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Fats != nil {
		queryValues = append(queryValues, *data.Fats)
		queryParams = append(queryParams, "fats=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Carbohydrates != nil {
		queryValues = append(queryValues, *data.Carbohydrates)
		queryParams = append(queryParams, "carbohydrates=$"+strconv.Itoa(len(queryValues)))
	}
	if len(queryParams) == 0 || len(queryValues) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "UPDATE product SET " + strings.Join(queryParams, ", ") + " where id=$" + strconv.Itoa(len(queryValues))
	fmt.Println(query)
	_, err := r.db.Exec(query, queryValues...)
	return err
}

func (r *ProductRepository) GetList(limit, offset int, filters *model.ProductFilter) ([]*model.Product, error) {
	tempQuery := "select id, title, slug, description, calories, squirrels, fats, carbohydrates, photo, suggested_by_user, created_by_id from product"
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
	products := []*model.Product{}
	for rows.Next() {
		p := &model.Product{}
		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Slug,
			&p.Description,
			&p.Calories,
			&p.Squirrels,
			&p.Fats,
			&p.Carbohydrates,
			&p.Photo,
			&p.SuggestedByUser,
			&p.CreatedById,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) prepareFilters(query string, filters *model.ProductFilter) (string, []interface{}, error) {
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
	// Calories
	if filters.CaloriesGTE != nil {
		filterValues = append(filterValues, *filters.CaloriesGTE)
		filterQuery = append(filterQuery, "calories >= $"+strconv.Itoa(len(filterValues)))
	}
	if filters.CaloriesLTE != nil {
		filterValues = append(filterValues, *filters.CaloriesLTE)
		filterQuery = append(filterQuery, "calories <= $"+strconv.Itoa(len(filterValues)))
	}
	// Squirrels
	if filters.SquirrelsGTE != nil {
		filterValues = append(filterValues, *filters.SquirrelsGTE)
		filterQuery = append(filterQuery, "squirrels >= $"+strconv.Itoa(len(filterValues)))
	}
	if filters.SquirrelsLTE != nil {
		filterValues = append(filterValues, *filters.SquirrelsLTE)
		filterQuery = append(filterQuery, "squirrels <= $"+strconv.Itoa(len(filterValues)))
	}
	// Fats
	if filters.FatsGTE != nil {
		filterValues = append(filterValues, *filters.FatsGTE)
		filterQuery = append(filterQuery, "fats >= $"+strconv.Itoa(len(filterValues)))
	}
	if filters.FatsLTE != nil {
		filterValues = append(filterValues, *filters.FatsLTE)
		filterQuery = append(filterQuery, "fats <= $"+strconv.Itoa(len(filterValues)))
	}
	// Carbohydrates
	if filters.CarbohydratesGTE != nil {
		filterValues = append(filterValues, *filters.CarbohydratesGTE)
		filterQuery = append(filterQuery, "carbohydrates >= $"+strconv.Itoa(len(filterValues)))
	}
	if filters.CarbohydratesLTE != nil {
		filterValues = append(filterValues, *filters.CarbohydratesLTE)
		filterQuery = append(filterQuery, "carbohydrates <= $"+strconv.Itoa(len(filterValues)))
	}
	if len(filterValues) == 0 {
		return query, filterValues, nil
	}
	return query + " where " + strings.Join(filterQuery, " AND "), filterValues, nil
}
