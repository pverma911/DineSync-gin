package repository

import (
    "errors"

    "github.com/pverma911/go-gin-tonic/internal/model"
    "gorm.io/gorm"
)

type AddressRepository struct {
    db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
    return &AddressRepository{db: db}
}

func (r *AddressRepository) GetAddressesByUserID(userID uint) ([]model.Address, error) {
    var addresses []model.Address
    if err := r.db.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
        return nil, err
    }
    return addresses, nil
}

func (r *AddressRepository) GetAddressByID(id uint) (*model.Address, error) {
    var address model.Address
    if err := r.db.First(&address, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("address not found")
        }
        return nil, err
    }
    return &address, nil
}

func (r *AddressRepository) CreateAddress(address *model.Address) (uint, error) {
    if err := r.db.Create(address).Error; err != nil {
        return 0, err
    }
    return address.ID, nil
}