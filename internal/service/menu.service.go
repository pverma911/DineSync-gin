package service

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/repository"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type MenuService struct {
    MenuRepo *repository.MenuRepository
}

func NewMenuService(mr *repository.MenuRepository) *MenuService {
    return &MenuService{MenuRepo: mr}
}

func (s *MenuService) GetMenuItems() utils.ServiceResponse {
    items, err := s.MenuRepo.GetMenuItems()
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"menu": items}, "Menu fetched")
}

func (s *MenuService) CreateMenuItem(item model.MenuItem) utils.ServiceResponse {
    id, err := s.MenuRepo.CreateMenuItem(&item)
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"id": id}, "Menu item created")
}