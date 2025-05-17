package model

import "time"

type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	Birthday  *time.Time // if pointer is not used, then value can't be null, it will have default go time
	Is_Active bool       `json:"isActive"`
	CreatedAt time.Time  `gorm:"->" json:"createdAt"`// Automatically managed by GORM for creation time
	UpdatedAt time.Time  `gorm:"->" json:"updatedAt"`// Automatically managed by GORM for update time
}
