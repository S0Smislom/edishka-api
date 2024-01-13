package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
	"time"
)

type DietService struct {
	repo repository.Repository
}

func NewDietService(repo repository.Repository) *DietService {
	return &DietService{repo: repo}
}

func (s *DietService) GetByIdPrivate(currentUserId, id int) (*model.Diet, error) {
	dbModel, err := s.repo.Diet().GetById(id)
	if err != nil {
		return nil, err
	}
	if dbModel.UserId != currentUserId {
		return nil, &exceptions.UserPermissionError{Msg: "Forbidden"}
	}
	return dbModel, err
}

func (s *DietService) GetListPrivate(currentUserId, limit, offset int, filters *model.DietFilter) (*model.DietList, error) {
	filters.UserId = currentUserId
	total, err := s.repo.Diet().Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.DietList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.Diet{}
		return result, nil
	}
	data, err := s.repo.Diet().GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}

func (s *DietService) Create(currentUserId int, data *model.CreateDiet) (*model.Diet, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}

	data.UserId = currentUserId
	now := time.Now().UTC()
	data.CreatedAt = now
	data.UpdatedAt = now
	id, err := s.repo.Diet().Create(data)
	if err != nil {
		return nil, err
	}
	if data.ItemIds != nil {
		s.updateItems(currentUserId, id, data.ItemIds)
	}
	return s.GetByIdPrivate(currentUserId, id)
}

func (s *DietService) Update(currentUserId, id int, data *model.UpdateDiet) (*model.Diet, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	if _, err := s.GetByIdPrivate(currentUserId, id); err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	data.UpdatedAt = now
	if err := s.repo.Diet().Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.Diet().GetById(id)
}

func (s *DietService) Delete(currentUserId, id int) (*model.Diet, error) {
	dbModel, err := s.GetByIdPrivate(currentUserId, id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Diet().Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *DietService) updateItems(currentUserId, dietId int, itemIds *[]int) {
	dbItems, err := s.repo.DietItem().GetList(100, 0, &model.DietItemFilter{
		CreatedById: &currentUserId,
		IdList:      itemIds,
	})
	if err != nil || len(dbItems) == 0 {
		return
	}
	// TODO Run in goroutine
	for _, item := range dbItems {
		if item.DietId != nil {
			continue
		}
		s.repo.DietItem().Update(
			item.Id,
			&model.UpdateDietItem{
				DietId: &dietId,
			},
		)
	}
}
