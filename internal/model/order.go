package model

import "time"

type Order struct {
    ID         uint      `json:"id"`
    CustomerID uint      `json:"customerId"`
    Items      string    `json:"items"` // Could be JSON or a separate table for order items
    Total      float64   `json:"total"`
    Status     string    `json:"status"`
    CreatedAt  time.Time `gorm:"->" json:"createdAt"`
    UpdatedAt  time.Time `gorm:"->" json:"updatedAt"`
}