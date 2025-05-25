package handler

import (
	"net/http"
	"strconv"

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

func (u *UserHandler) GetUsers(c *gin.Context) {
	res := u.userService.GetUsers()
	utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (u *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idUint64, _ := strconv.ParseUint(id, 10, 64)

	res := u.userService.GetUserByID(uint(idUint64))

	utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	res := u.userService.CreateUser(user)
	utils.SendHandlerResponse(c, res.StatusCode, res)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Updated user", "id": id})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted user", "id": id})
}
