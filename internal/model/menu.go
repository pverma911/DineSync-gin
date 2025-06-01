package model

import "time"

type MenuItem struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Available   bool      `json:"available"`
    CreatedAt   time.Time `gorm:"->" json:"createdAt"`
    UpdatedAt   time.Time `gorm:"->" json:"updatedAt"`
}