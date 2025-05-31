package router

import (
    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/handler"
)

type CustomerRoutes struct {
    customerHandler *handler.CustomerHandler
}

func NewCustomerRoutes(handler *handler.CustomerHandler) *CustomerRoutes {
    return &CustomerRoutes{customerHandler: handler}
}

func (cr *CustomerRoutes) RegisterCustomerRoutes(r *gin.RouterGroup) {
    customerRoutes := r.Group("/customers")
    {
        customerRoutes.GET("/", cr.customerHandler.GetCustomers)
        customerRoutes.GET("/:id", cr.customerHandler.GetCustomerByID)
        customerRoutes.POST("/", cr.customerHandler.CreateCustomer)
    }
}