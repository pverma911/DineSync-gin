package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/service"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type AddressHandler struct {
    addressService *service.AddressService
}

func NewAddressHandler(as *service.AddressService) *AddressHandler {
    return &AddressHandler{addressService: as}
}

func (h *AddressHandler) GetAddressesByUserID(c *gin.Context) {
    userIDStr := c.Param("userId")
    userIDUint64, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    res := h.addressService.GetAddressesByUserID(uint(userIDUint64))
    utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *AddressHandler) GetAddressByID(c *gin.Context) {
    idStr := c.Param("id")
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
        return
    }
    res := h.addressService.GetAddressByID(uint(idUint64))
    utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
    var address model.Address
    if err := c.ShouldBindJSON(&address); err != nil {
        utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res := h.addressService.CreateAddress(address)
    utils.SendHandlerResponse(c, res.StatusCode, res)
}