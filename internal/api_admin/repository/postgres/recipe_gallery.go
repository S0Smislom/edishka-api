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

type RecipeGalleryRepository struct {
	db *sql.DB
}

func NewRecipeGalleryRepository(db *sql.DB) *RecipeGalleryRepository {
	return &RecipeGalleryRepository{db: db}
}

func (r *RecipeGalleryRepository) Create(data *model.CreateRecipeGallery) (int, error) {
	var id int
	row := r.db.QueryRow(
		"insert into recipe_gallery (photo, recipe_id, published, ordering) values ($1, $2, $3, $4) returning id",
		data.Photo,
		data.RecipeId,
		data.Published,
		data.Ordering,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecipeGalleryRepository) GetById(id int) (*model.RecipeGallery, error) {
	p := &model.RecipeGallery{}
	if err := r.db.QueryRow(
		"select id, ordering, recipe_id, created_at, updated_at, photo, published from recipe_gallery where id=$1",
		id,
	).Scan(
		&p.Id,
		&p.Ordering,
		&p.RecipeId,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Photo,
		&p.Published,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("RecipeGallery not found")
		}
		return nil, err
	}
	return p, nil
}

func (r *RecipeGalleryRepository) GetList(limit, offset int, filters *model.RecipeGalleryFilter) ([]*model.RecipeGallery, error) {
	tempQuery := "select id, ordering, recipe_id, created_at, updated_at, photo, published from recipe_gallery"
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
	recipeGallerys := []*model.RecipeGallery{}
	for rows.Next() {
		p := &model.RecipeGallery{}
		if err := rows.Scan(
			&p.Id,
			&p.Ordering,
			&p.RecipeId,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Photo,
			&p.Published,
		); err != nil {
			return nil, err
		}
		recipeGallerys = append(recipeGallerys, p)
	}
	return recipeGallerys, nil
}

func (r *RecipeGalleryRepository) Count(filters *model.RecipeGalleryFilter) (int, error) {
	tempQuery := "select count(*) from recipe_gallery"
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

func (r *RecipeGalleryRepository) Update(id int, data *model.UpdateRecipeGallery) error {
	var queryParams []string
	var queryValues []interface{}
	// TODO придумать более изящный способ обновления
	queryValues = append(queryValues, time.Now())
	queryParams = append(queryParams, "updated_at=$"+strconv.Itoa(len(queryValues)))
	if data.Published != nil {
		queryValues = append(queryValues, *data.Published)
		queryParams = append(queryParams, "published=$"+strconv.Itoa(len(queryValues)))
	}
	if data.Ordering != nil {
		queryValues = append(queryValues, *data.Ordering)
		queryParams = append(queryParams, "ordering=$"+strconv.Itoa(len(queryValues)))
	}
	if len(queryParams) == 0 || len(queryValues) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "UPDATE recipe_gallery SET " + strings.Join(queryParams, ", ") + " where id=$" + strconv.Itoa(len(queryValues))
	fmt.Println(query)
	_, err := r.db.Exec(query, queryValues...)
	return err
}

func (r *RecipeGalleryRepository) Delete(id int) error {
	_, err := r.db.Exec("delete from recipe_gallery where id=$1", id)
	return err
}

func (r *RecipeGalleryRepository) UpdatePhoto(id int, photo *string) error {
	_, err := r.db.Exec("update recipe_gallery set photo=$1 where id=$2", photo, id)
	return err
}

func (r *RecipeGalleryRepository) prepareFilters(query string, filters *model.RecipeGalleryFilter) (string, []interface{}, error) {
	var filterValues []interface{}
	var filterQuery []string
	// Recipe id
	filterValues = append(filterValues, filters.RecipeId)
	filterQuery = append(filterQuery, "recipe_id = $"+strconv.Itoa(len(filterValues)))
	if len(filterValues) == 0 {
		return query, filterValues, nil
	}
	return query + " where " + strings.Join(filterQuery, " AND "), filterValues, nil
}
