package gormrepo

import (
	"errors"
	"food/internal/api/model"
	"food/pkg/exceptions"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	user := &model.User{}
	res := r.db.First(&user, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &exceptions.ObjectNotFoundError{Msg: "User not found"}
		}
		return nil, res.Error
	}
	return user, nil
}

func (r *UserRepository) GetByPhone(phone string) (int, error) {
	user := &model.User{}
	res := r.db.Select("id").Where("phone=?", phone).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return 0, &exceptions.ObjectNotFoundError{Msg: "User not found"}
		}
		return 0, res.Error
	}
	return user.ID, nil
}

func (r *UserRepository) UpdateCode(id int, code string) error {
	return r.db.Model(&model.User{}).Where("id=?", id).Update("code", code).Error
}

func (r *UserRepository) UpdatePhoto(id int, photo *string) error {
	return r.db.Model(&model.User{}).Where("id=?", id).Update("photo", photo).Error
}

func (r *UserRepository) Update(id int, data *model.UpdateUser) error {
	return r.db.Model(&model.User{}).Where("id=?", id).Updates(data).Error
}
