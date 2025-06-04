package router

import (
    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/handler"
)

type MenuRoutes struct {
    menuHandler *handler.MenuHandler
}

func NewMenuRoutes(handler *handler.MenuHandler) *MenuRoutes {
    return &MenuRoutes{menuHandler: handler}
}

func (mr *MenuRoutes) RegisterMenuRoutes(r *gin.RouterGroup) {
    menuRoutes := r.Group("/menu")
    {
        menuRoutes.GET("/", mr.menuHandler.GetMenuItems)
        menuRoutes.POST("/", mr.menuHandler.CreateMenuItem)
    }
}