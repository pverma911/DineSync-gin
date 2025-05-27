package router

import (
    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/handler"
)

type AddressRoutes struct {
    addressHandler *handler.AddressHandler
}

func NewAddressRoutes(handler *handler.AddressHandler) *AddressRoutes {
    return &AddressRoutes{addressHandler: handler}
}

func (a *AddressRoutes) RegisterAddressRoutes(r *gin.RouterGroup) {
    addressRoutes := r.Group("/addresses")
    {
        addressRoutes.GET("/user/:userId", a.addressHandler.GetAddressesByUserID)
        addressRoutes.GET("/:id", a.addressHandler.GetAddressByID)
        addressRoutes.POST("/", a.addressHandler.CreateAddress)
    }
}