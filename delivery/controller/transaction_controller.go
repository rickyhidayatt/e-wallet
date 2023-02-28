package controller

import (
	"e-wallet/model"
	"e-wallet/usecase"
	"e-wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

// / bikin Method Handler nya disini
func (tc *TransactionController) TopUp(c *gin.Context) {

	var topup model.Transaction
	err := c.ShouldBindJSON(&topup)
	if err != nil {
		utils.HandleBadRequest(c, err.Error())
	}

	amount, err := tc.transactionUseCase.TopUp(topup.UserId, topup.Amount)
	if err != nil {
		utils.HandleInternalServerError(c, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"balance": amount,
	})

}

func (tc *TransactionController) Transfer(c *gin.Context) {
	var Transfer model.Transfer
	err := c.ShouldBindJSON(&Transfer)
	if err != nil {
		utils.HandleBadRequest(c, err.Error())
	}
	trf, err := tc.transactionUseCase.SendMoney(Transfer.UserId, Transfer.Amount, Transfer.BankName, Transfer.Category, Transfer.AccountNumber, Transfer.ReceiverName)

	if err != nil {
		utils.HandleInternalServerError(c, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success Transaction": trf,
		})
	}
}

func NewTransactionController(router *gin.Engine, transactionUc usecase.TransactionUseCase) *TransactionController {
	trxController := TransactionController{
		transactionUseCase: transactionUc,
	}

	r := router.Group("api/transaction")
	r.POST("/topup", trxController.TopUp)
	r.POST("/transfer", trxController.Transfer)

	return &trxController
}
