package repository

import (
	"github.com/pverma911/go-gin-tonic/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user *model.User) (uint, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *model.User) (uint, error) {
	res := ur.db.Create(user)

	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}
