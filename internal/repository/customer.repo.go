package repository

import (
    "errors"

    "github.com/pverma911/go-gin-tonic/internal/model"
    "gorm.io/gorm"
)

type CustomerRepository struct {
    db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
    return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetCustomers() ([]model.Customer, error) {
    var customers []model.Customer
    if err := r.db.Find(&customers).Error; err != nil {
        return nil, err
    }
    return customers, nil
}

func (r *CustomerRepository) GetCustomerByID(id uint) (*model.Customer, error) {
    var customer model.Customer
    if err := r.db.First(&customer, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("customer not found")
        }
        return nil, err
    }
    return &customer, nil
}

func (r *CustomerRepository) CreateCustomer(customer *model.Customer) (uint, error) {
    if err := r.db.Create(customer).Error; err != nil {
        return 0, err
    }
    return customer.ID, nil
}