package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Message      string
	ResponseCode int
	Result       interface{}
}

func HandleSuccess(c *gin.Context, data interface{}, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 200,
		Result:       data,
	}
	c.JSON(http.StatusOK, response)
}

func HandleSuccessCreated(c *gin.Context, message string, data interface{}) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 201,
		Result:       data,
	}
	c.JSON(http.StatusCreated, response)
}

func HandleNotFound(c *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 404,
		Result:       nil,
	}
	c.JSON(http.StatusNotFound, response)
}

func HandleInternalServerError(c *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 500,
		Result:       nil,
	}
	c.JSON(http.StatusInternalServerError, response)
}

func HandleBadRequest(c *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 400,
		Result:       nil,
	}
	c.JSON(http.StatusBadRequest, response)
}
