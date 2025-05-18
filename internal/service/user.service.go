package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/model"
	"github.com/pverma911/go-gin-tonic/internal/repository"
	"github.com/pverma911/go-gin-tonic/internal/utils"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

type IUserService interface {
	GetUsers() []model.User
	CreateUser(user model.User) uint
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{UserRepo: ur}
}

func (s *UserService) CreateUser(user model.User) utils.ServiceResponse {
	res, err := s.UserRepo.CreateUser(&user)

	if err != nil {
		return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
	}

	return utils.SendServiceResponse(http.StatusOK, gin.H{"id": res}, "User created")
}
