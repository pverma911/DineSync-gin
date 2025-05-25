package service

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/pverma911/go-gin-tonic/internal/model"
    "github.com/pverma911/go-gin-tonic/internal/repository"
    "github.com/pverma911/go-gin-tonic/internal/utils"
)

type AddressService struct {
    AddressRepo *repository.AddressRepository
}

func NewAddressService(ar *repository.AddressRepository) *AddressService {
    return &AddressService{AddressRepo: ar}
}

func (s *AddressService) GetAddressesByUserID(userID uint) utils.ServiceResponse {
    addresses, err := s.AddressRepo.GetAddressesByUserID(userID)
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"addresses": addresses}, "Addresses fetched")
}

func (s *AddressService) GetAddressByID(id uint) utils.ServiceResponse {
    address, err := s.AddressRepo.GetAddressByID(id)
    if err != nil {
        return utils.SendServiceResponse(http.StatusNotFound, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"address": address}, "Address fetched")
}

func (s *AddressService) CreateAddress(address model.Address) utils.ServiceResponse {
    id, err := s.AddressRepo.CreateAddress(&address)
    if err != nil {
        return utils.SendServiceResponse(http.StatusInternalServerError, gin.H{}, err.Error())
    }
    return utils.SendServiceResponse(http.StatusOK, gin.H{"id": id}, "Address created")
}