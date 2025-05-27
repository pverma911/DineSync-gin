package config

import (
	"log"

	"github.com/pverma911/go-gin-tonic/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=go_gin_tonic port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error while initializing DB %s", err.Error())
	}

	return db
}

func CreateOrUpdateTables(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Address{})
}
