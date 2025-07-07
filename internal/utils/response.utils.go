package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/pverma911/go-gin-tonic/internal/helpers"
)

type ServiceResponse struct {
	Success    bool   `json:"success"`
	Data       gin.H  `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

var errorCodes []int = []int{400, 401, 403, 404, 412}

func SendServiceResponse(statusCode int, data gin.H, message string) ServiceResponse {
	return ServiceResponse{
		Success:    !helpers.DoesInclude(statusCode, errorCodes),
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
}

func SendHandlerResponse(c *gin.Context,statusCode int,data any){
	c.JSON(statusCode,data) 
}