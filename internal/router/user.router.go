package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/handler"
)

type UserRoutes struct {
    userHandler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
    return &UserRoutes{userHandler: handler}
}

func(u UserRoutes) RegisterUserRoutes(r *gin.RouterGroup) {

	userRoutes := r.Group("/users")
	{
		// userRoutes.GET("/", u.userHandler.GetUsers)
		userRoutes.GET("/:id",  u.userHandler.GetUserByID)
		userRoutes.POST("/",  u.userHandler.CreateUser)
		userRoutes.PUT("/:id",  u.userHandler.UpdateUser)
		userRoutes.DELETE("/:id",  u.userHandler.DeleteUser)
	}
}
