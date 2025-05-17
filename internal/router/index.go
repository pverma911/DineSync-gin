package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/handler"
	"github.com/pverma911/go-gin-tonic/internal/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	router:=r.Group("/api/v1")

 // Create service, handler, and routes
    userService := &service.UserService{}
    userHandler := handler.NewUserHandler(userService)
    NewUserRoutes(userHandler).RegisterUserRoutes(router)

	return r
}
