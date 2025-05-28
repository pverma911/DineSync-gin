package service

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/repository"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type CustomerService struct {
    CustomerRepo *repository.CustomerRepository
}

func NewCustomerService(cr *repository.CustomerRepository) *CustomerService {
    return &CustomerService{CustomerRepo: cr}
}

func (s *CustomerService) GetCustomers() utils.ServiceResponse {
    customers, err := s.CustomerRepo.GetCustomers()
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"customers": customers}, "Customers fetched")
}

func (s *CustomerService) GetCustomerByID(id uint) utils.ServiceResponse {
    customer, err := s.CustomerRepo.GetCustomerByID(id)
    if err != nil {
        return utils.SendServiceResponse(http.StatusNotFound, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"customer": customer}, "Customer fetched")
}

func (s *CustomerService) CreateCustomer(customer model.Customer) utils.ServiceResponse {
    id, err := s.CustomerRepo.CreateCustomer(&customer)
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"id": id}, "Customer created")
}