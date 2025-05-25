package model

import "time"

type Address struct {
    ID        uint      `json:"id"`
    UserID    uint      `json:"userId"`
    Street    string    `json:"street"`
    City      string    `json:"city"`
    State     string    `json:"state"`
    ZipCode   string    `json:"zipCode"`
    Country   string    `json:"country"`
    CreatedAt time.Time `gorm:"->" json:"createdAt"`
    UpdatedAt time.Time `gorm:"->" json:"updatedAt"`
}