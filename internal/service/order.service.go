package service

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/repository"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type OrderService struct {
    OrderRepo *repository.OrderRepository
}

func NewOrderService(or *repository.OrderRepository) *OrderService {
    return &OrderService{OrderRepo: or}
}

func (s *OrderService) GetOrders() utils.ServiceResponse {
    orders, err := s.OrderRepo.GetOrders()
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"orders": orders}, "Orders fetched")
}

func (s *OrderService) GetOrderByID(id uint) utils.ServiceResponse {
    order, err := s.OrderRepo.GetOrderByID(id)
    if err != nil {
        return utils.SendServiceResponse(http.StatusNotFound, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"order": order}, "Order fetched")
}

func (s *OrderService) CreateOrder(order model.Order) utils.ServiceResponse {
    id, err := s.OrderRepo.CreateOrder(&order)
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"id": id}, "Order created")
}