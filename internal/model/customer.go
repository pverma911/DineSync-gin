package model

import "time"

type Customer struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Phone     string    `json:"phone"`
    Email     string    `json:"email"`
    CreatedAt time.Time `gorm:"->" json:"createdAt"`
    UpdatedAt time.Time `gorm:"->" json:"updatedAt"`
}