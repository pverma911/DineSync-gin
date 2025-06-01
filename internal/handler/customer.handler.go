package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/model"
	"github.com/pverma911/go-gin-tonic/internal/service"
	"github.com/pverma911/go-gin-tonic/internal/utils"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(cs *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: cs}
}

func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	res := h.customerService.GetCustomers()
	utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	res := h.customerService.GetCustomerByID(uint(idUint64))
	utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		utils.SendHandlerResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := h.customerService.CreateCustomer(customer)
	utils.SendHandlerResponse(c, res.StatusCode, res)
}
