package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/service"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type MenuHandler struct {
    menuService *service.MenuService
}

func NewMenuHandler(ms *service.MenuService) *MenuHandler {
    return &MenuHandler{menuService: ms}
}

func (h *MenuHandler) GetMenuItems(c *gin.Context) {
    res := h.menuService.GetMenuItems()
    utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *MenuHandler) CreateMenuItem(c *gin.Context) {
    var item model.MenuItem
    if err := c.ShouldBindJSON(&item); err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res := h.menuService.CreateMenuItem(item)
    utils.SendHandlerResponse(c, res.StatusCode, res)
}