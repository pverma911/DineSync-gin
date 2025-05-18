package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/model"
	"github.com/pverma911/go-gin-tonic/internal/service"
	"github.com/pverma911/go-gin-tonic/internal/utils"
)

type UserHandler struct {
	userService *service.UserService
}

// Constructor to initialize dependencies
func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

// func (u *UserHandler) GetUsers(c *gin.Context) {
// 	response := u.userService.GetUsers()
// 	c.JSON(http.StatusOK, gin.H{"data": response})
// }

func (u *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get user by ID", "id": id})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	res:= u.userService.CreateUser(user)
	utils.SendHandlerResponse(c,res.StatusCode,res)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Updated user", "id": id})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted user", "id": id})
}
