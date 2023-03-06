package controller

import (
	"e-wallet/delivery/middleware"
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

func (tc *TransactionController) TopUp(c *gin.Context) {
	var topup model.Transaction
	err := c.ShouldBindJSON(&topup)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	amount, err := tc.transactionUseCase.TopUp(topup.UserId, topup.Amount)

	if err != nil {
		response := utils.ApiResponse("top-up failed", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successful top-up", http.StatusOK, "success", amount)
		c.JSON(http.StatusOK, response)
	}
}

func (tc *TransactionController) Transfer(c *gin.Context) {
	var Transfer model.TransactionSend
	err := c.ShouldBindJSON(&Transfer)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	trf, err := tc.transactionUseCase.SendMoney(Transfer)

	if err != nil {
		response := utils.ApiResponse("transaction failed", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := utils.ApiResponse("successful transaction", http.StatusOK, "success", trf)
		c.JSON(http.StatusOK, response)
	}
}

func (tc *TransactionController) RequestMoney(c *gin.Context) {
	var transactions model.TransactionRequest
	err := c.ShouldBindJSON(&transactions)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req, err := tc.transactionUseCase.RequestMoney(transactions)

	if err != nil {
		response := utils.ApiResponse("payment Request failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := utils.ApiResponse("payment Request success", http.StatusOK, "success", req)
		c.JSON(http.StatusOK, response)
	}
}

func (tc *TransactionController) PrintTransactionHistory(c *gin.Context) {
	id := c.Param("id")
	transactions, err := tc.transactionUseCase.PrintHistoryTransactionsById(id)

	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := utils.ApiResponse("transaction history", http.StatusOK, "success", transactions)
		c.JSON(http.StatusOK, response)
	}
}

func (tc *TransactionController) GiftBirthDay(c *gin.Context) {
	id := c.Param("id")

	err := tc.transactionUseCase.GetBonus(id)
	if err != nil {
		response := utils.ApiResponse("server error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	message := "You receive a cash prize of IDR 25,000 from an online wallet"
	response := utils.ApiResponse("Happy BirthDay", http.StatusOK, "success", message)
	c.JSON(http.StatusOK, response)

}

func NewTransactionController(router *gin.Engine, transactionArg usecase.TransactionUseCase) *TransactionController {
	trxController := TransactionController{
		transactionUseCase: transactionArg,
	}

	r := router.Group("api/transaction")
	r.Use(middleware.AuthMiddleware())
	r.GET("/:id/gift", trxController.GiftBirthDay)
	r.POST("/topup", trxController.TopUp)
	r.POST("/transfer", trxController.Transfer)
	r.POST("/request", trxController.RequestMoney)
	r.GET("/:id/history", trxController.PrintTransactionHistory)

	return &trxController
}
