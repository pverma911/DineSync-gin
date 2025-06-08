package repository

import (
    "errors"

    "github.com/pverma911/go-gin-tonic/internal/model"
    "gorm.io/gorm"
)

type OrderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrders() ([]model.Order, error) {
    var orders []model.Order
    if err := r.db.Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (r *OrderRepository) GetOrderByID(id uint) (*model.Order, error) {
    var order model.Order
    if err := r.db.First(&order, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("order not found")
        }
        return nil, err
    }
    return &order, nil
}

func (r *OrderRepository) CreateOrder(order *model.Order) (uint, error) {
    if err := r.db.Create(order).Error; err != nil {
        return 0, err
    }
    return order.ID, nil
}