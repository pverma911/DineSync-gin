package repository

import (
	"errors"

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

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(user *model.User) (uint, error) {
	res := ur.db.Create(user)

	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}
