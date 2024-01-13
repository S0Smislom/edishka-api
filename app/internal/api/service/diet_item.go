package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
)

type DietItemService struct {
	repo repository.Repository
}

func NewDietItemService(repo repository.Repository) *DietItemService {
	return &DietItemService{repo: repo}
}

func (s *DietItemService) GetByIdPrivate(currentUserId, id int) (*model.DietItem, error) {
	dbModel, err := s.repo.DietItem().GetById(id)
	if err != nil {
		return nil, err
	}

	if dbModel.CreatedById != currentUserId {
		return nil, &exceptions.UserPermissionError{Msg: "Forbidden"}
	}
	return dbModel, err
}

func (s *DietItemService) GetListPrivate(currentUserId, limit, offset int, filters *model.DietItemFilter) (*model.DietItemList, error) {
	if err := filters.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}

	filters.CreatedById = &currentUserId
	total, err := s.repo.DietItem().Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.DietItemList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.DietItem{}
		return result, nil
	}
	data, err := s.repo.DietItem().GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}

func (s *DietItemService) Create(currentUserId int, data *model.CreateDietItem) (*model.DietItem, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	if data.DietId != nil {
		dbModel, err := s.repo.Diet().GetById(*data.DietId)
		if err != nil {
			return nil, err
		}
		if dbModel.UserId != currentUserId {
			return nil, &exceptions.UserPermissionError{Msg: "Forbidden"}
		}
	}
	data.CreatedById = currentUserId
	id, err := s.repo.DietItem().Create(data)
	if err != nil {
		return nil, err
	}
	return s.GetByIdPrivate(data.CreatedById, id)
}

func (s *DietItemService) Update(currentUserId, id int, data *model.UpdateDietItem) (*model.DietItem, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	_, err := s.GetByIdPrivate(currentUserId, id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.DietItem().Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.DietItem().GetById(id)
}

func (s *DietItemService) Delete(currentUserId, id int) (*model.DietItem, error) {
	dbModel, err := s.GetByIdPrivate(currentUserId, id)
	if err != nil {
		return nil, err
	}
	s.repo.DietItem().Delete(id)
	return dbModel, nil
}
