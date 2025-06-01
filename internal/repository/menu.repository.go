package repository

import (
    "github.com/pverma911/go-gin-tonic/internal/model"
    "gorm.io/gorm"
)

type MenuRepository struct {
    db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
    return &MenuRepository{db: db}
}

func (r *MenuRepository) GetMenuItems() ([]model.MenuItem, error) {
    var items []model.MenuItem
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *MenuRepository) CreateMenuItem(item *model.MenuItem) (uint, error) {
    if err := r.db.Create(item).Error; err != nil {
        return 0, err
    }
    return item.ID, nil
}