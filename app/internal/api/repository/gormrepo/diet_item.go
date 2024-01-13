package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"
	"food/pkg/utils"

	"gorm.io/gorm"
)

type DietItemRepository struct {
	db *gorm.DB
}

func NewDietItemRepository(db *gorm.DB) *DietItemRepository {
	return &DietItemRepository{db: db}
}

func (r *DietItemRepository) Create(data *model.CreateDietItem) (int, error) {
	res := r.db.Table("diet_item").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}

func (r *DietItemRepository) GetById(id int) (*model.DietItem, error) {
	dbModel := &model.DietItem{}
	res := r.db.First(&dbModel, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "Diet item not found"}
		}
		return nil, res.Error
	}
	return dbModel, nil
}

func (r *DietItemRepository) Count(filters *model.DietItemFilter) (int, error) {
	var total int64
	res := r.db.Table("diet_item").Scopes(r.getFilterQuery(filters)).Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(total), nil
}

func (r *DietItemRepository) GetList(limit, offset int, filters *model.DietItemFilter) ([]*model.DietItem, error) {
	dbModels := []*model.DietItem{}
	res := r.db.Scopes(r.getFilterQuery(filters)).Limit(limit).Offset(offset).Find(&dbModels)
	if res.Error != nil {
		return []*model.DietItem{}, res.Error
	}
	return dbModels, nil
}

func (r *DietItemRepository) Update(id int, data *model.UpdateDietItem) error {
	return r.db.Model(&model.DietItem{}).Where("id=?", id).Updates(data).Error
}
func (r *DietItemRepository) Delete(id int) error {
	return r.db.Model(&model.DietItem{}).Delete(&model.DietItem{}, id).Error
}

func (r *DietItemRepository) getFilterQuery(filters *model.DietItemFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filters.DietId != 0 {
			db = db.Where("diet_id=?", filters.DietId)
		}
		if filters.IdList != nil {
			db = db.Where("id IN ?", utils.ConvertIntListToStringList(*filters.IdList))
		}
		if filters.CreatedById != nil {
			db = db.Where("created_by_id = ?", filters.CreatedById)
		}
		return db
	}
}
