package controller

import (
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func (uc *UserController) Login(c *gin.Context) {
	var login model.User
	err := c.ShouldBindJSON(&login)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	login, err = uc.userUseCase.Login(login.Email, login.Password)
	if err != nil {
		response := utils.ApiResponse("login failed for user", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successful login to your account", http.StatusOK, "success", login)
		c.JSON(http.StatusOK, response)
	}

}

func NewUserController(router *gin.Engine, userArg usecase.UserUseCase) *UserController {
	userController := UserController{
		userUseCase: userArg,
	}

	r := router.Group("api/user")
	r.POST("/login", userController.Login)
	return &userController
}
