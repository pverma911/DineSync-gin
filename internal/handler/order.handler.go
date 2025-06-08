package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/service"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type OrderHandler struct {
    orderService *service.OrderService
}

func NewOrderHandler(os *service.OrderService) *OrderHandler {
    return &OrderHandler{orderService: os}
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
    res := h.orderService.GetOrders()
    utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
    idStr := c.Param("id")
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }
    res := h.orderService.GetOrderByID(uint(idUint64))
    utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var order model.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res := h.orderService.CreateOrder(order)
    utils.SendHandlerResponse(c, res.StatusCode, res)
}