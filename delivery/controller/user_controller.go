package controller

import (
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func (uc *UserController) ViewAllUser(ctx *gin.Context) {
	page,_ := strconv.Atoi(ctx.Query("page"))
	 if page == 0 {
		page = 1
	}
	size,_ := strconv.Atoi(ctx.Query("size"))
	if size == 0 {
		size = 5
	}
	users, err := uc.userUseCase.ViewAllUser(page,size)
	if err != nil {
		utils.HandleInternalServerError(ctx,err.Error())
	}else{
		utils.HandleSuccess(ctx, users, "Success get all users")
	}
}

func (uc *UserController) CreateNewUser(ctx *gin.Context) {
	var newUser *model.User
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := uc.userUseCase.CreateNewUser(newUser)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		}else{
			utils.HandleSuccessCreated(ctx,"Success create new user", newUser)
		}
	}
}

func NewUserController(router *gin.Engine, userUc usecase.UserUseCase) *UserController {
	newUserController := UserController {
		userUseCase: userUc,
	}
	//routerGroup := router.Group("")
	router.GET("user", newUserController.ViewAllUser)
	router.POST("user", newUserController.CreateNewUser)
	return &newUserController
}