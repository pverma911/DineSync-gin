package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Get all users"})
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Get user by ID", "id": id})
}

func CreateUser(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Updated user", "id": id})
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Deleted user", "id": id})
}
