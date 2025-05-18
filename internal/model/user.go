package model

import "time"

type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name" required:"true"`
	Age       int        `json:"age" required:"true"`
	Birthday  *time.Time `json:"birthday"`// if pointer is not used, then value can't be null, it will have default go time
	Is_Active bool       `json:"isActive" required:"true"`
	CreatedAt time.Time  `gorm:"->" json:"createdAt"`// Automatically managed by GORM for creation time
	UpdatedAt time.Time  `gorm:"->" json:"updatedAt"`// Automatically managed by GORM for update time
}
