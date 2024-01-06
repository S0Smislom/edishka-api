package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
	"log"
)

type ShoppingItemService struct {
	repo repository.ShoppingItem
}

func NewShoppingItemService(repo repository.ShoppingItem) *ShoppingItemService {
	return &ShoppingItemService{repo: repo}
}

func (s *ShoppingItemService) GetById(currentUserId, id int) (*model.ShoppingItem, error) {
	dbModel, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if dbModel.UserId != currentUserId {
		return nil, &exceptions.UserPermissionError{Msg: "Forbiden"}
	}
	return dbModel, err
}

func (s *ShoppingItemService) GetList(currentUserId, limit, offset int, filters *model.ShoppingItemFilter) (*model.ShoppingList, error) {
	filters.UserId = currentUserId
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.ShoppingList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.ShoppingItem{}
		return result, nil
	}
	data, err := s.repo.GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}

func (s *ShoppingItemService) Create(data *model.CreateShoppingItem) (*model.ShoppingItem, error) {
	log.Print(data)
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return s.GetById(data.UserId, id)
}

func (s *ShoppingItemService) Update(currentUserId, id int, data *model.UpdateShoppingItem) (*model.ShoppingItem, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	_, err := s.GetById(currentUserId, id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *ShoppingItemService) Delete(currentUserId, id int) (*model.ShoppingItem, error) {
	dbModel, err := s.GetById(currentUserId, id)
	if err != nil {
		return nil, err
	}
	s.repo.Delete(id)
	return dbModel, nil
}
