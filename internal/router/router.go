package router

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", handler.GetUsers)
        userRoutes.GET("/:id", handler.GetUserByID)
        userRoutes.POST("/", handler.CreateUser)
        userRoutes.PUT("/:id", handler.UpdateUser)
        userRoutes.DELETE("/:id", handler.DeleteUser)
    }

    return r
}
