package controller

import (
	"e-wallet/config"
	"e-wallet/delivery/middleware"
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		return
	}

	// Pembuatan JWT Token
	expTime := time.Now().Add(time.Minute * 120)
	claims := &config.JWTClaim{
		Email: login.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "REFACTOR PROJEK",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	fmt.Println("INI TOKEN", token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	response := utils.ApiResponse("successful login to your account", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)
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

func (uc *UserController) CheckEmail(c *gin.Context) {
	var input model.CheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := uc.userUseCase.IsEmailAvailable(input)
	if err != nil {
		data := gin.H{
			"is_available": true,
		}
		response := utils.ApiResponse("email is is available", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
	} else {
		data := gin.H{
			"is_available": isEmailAvailable,
		}

		response := utils.ApiResponse("email has been registered", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	}

}

func (uc *UserController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("no file you uploaded", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//save file
	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// nanti dapet dari jwt
	fakeId := "4e3c55684bc748b2a2b495191e4243a1"
	// nanti dari jwt
	_, err = uc.userUseCase.SaveAvatar(fakeId, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := utils.ApiResponse("success uploaded avatar", http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)

}

func (uc *UserController) Logout(c *gin.Context) {

	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(c.Writer, cookie)

	data := gin.H{
		"message": "Bye bye",
	}

	response := utils.ApiResponse("success Logout", http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}

func NewUserController(router *gin.Engine, userArg usecase.UserUseCase) *UserController {
	userController := UserController{
		userUseCase: userArg,
	}

	router.POST("/check-email", userController.CheckEmail)
	router.POST("/signup", userController.RegisterUser)
	router.POST("/login", userController.Login)

	r := router.Group("api/user")
	r.Use(middleware.AuthMiddleware())
	r.PUT("/update", userController.UpdateUser)
	r.POST("/avatars", userController.UploadAvatar)
	r.GET("/logout", userController.Logout)

	return &userController
}
