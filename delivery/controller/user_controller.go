package controller

import (
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func (uc *UserController) Login(c *gin.Context) {
	var login model.UserLogin

	err := c.ShouldBindJSON(&login)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	users, err := uc.userUseCase.Login(login)
	if err != nil {
		response := utils.ApiResponse("login failed for user", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successful login to your account", http.StatusOK, "success", users)
		c.JSON(http.StatusOK, response)
	}

}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user model.UserRegister
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	register, err := uc.userUseCase.RegisterUser(&user)
	if err != nil {
		response := utils.ApiResponse("failed to register account please try again later", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successfully register your account", http.StatusOK, "success", register)
		c.JSON(http.StatusOK, response)
	}
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user model.UserUpdate
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(&user)
	update, err := uc.userUseCase.UpdateUser(&user)
	if err != nil {
		response := utils.ApiResponse("failed to update account please try again later", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successfully update account", http.StatusOK, "success", update)
		c.JSON(http.StatusOK, response)
	}
}

func NewUserController(router *gin.Engine, userArg usecase.UserUseCase) *UserController {
	userController := UserController{
		userUseCase: userArg,
	}

	r := router.Group("api/user")
	r.POST("/login", userController.Login)
	r.POST("/signup", userController.RegisterUser)
	r.POST("/update", userController.UpdateUser)
	return &userController
}
