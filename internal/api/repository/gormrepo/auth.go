package gormrepo

import (
	"food/internal/api/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(data *model.Login) (int, error) {
	res := r.db.Table("\"user\"").Create(data)
	if res.Error != nil {
		return 0, res.Error
	}
	return data.Id, nil
}
