package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/handler"
)

type OrderRoutes struct {
	orderHandler *handler.OrderHandler
}

func NewOrderRoutes(handler *handler.OrderHandler) *OrderRoutes {
	return &OrderRoutes{orderHandler: handler}
}

func (o *OrderRoutes) RegisterOrderRoutes(r *gin.RouterGroup) {
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", o.orderHandler.GetOrders)
		orderRoutes.GET("/:id", o.orderHandler.GetOrderByID)
		orderRoutes.POST("/", o.orderHandler.CreateOrder)
	}
}
