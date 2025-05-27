package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/handler"
	"github.com/pverma911/go-gin-tonic/internal/repository"
	"github.com/pverma911/go-gin-tonic/internal/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	router := r.Group("/api/v1")

	// Create service, handler, and routes
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	NewUserRoutes(userHandler).RegisterUserRoutes(router)

	addressRepo := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)
	addressRoutes := NewAddressRoutes(addressHandler)
	addressRoutes.RegisterAddressRoutes(router)

	return r
}
