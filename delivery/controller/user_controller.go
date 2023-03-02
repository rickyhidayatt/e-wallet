package controller

import (
	"e-wallet/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

// func (uc *UserController)GetAllUser

func NewUserController(router *gin.Engine, userArg usecase.UserUseCase) *UserController {
	userController := UserController{
		userUseCase: userArg,
	}

	r := router.Group("api/user")
	r.GET("")
	return &userController
}
