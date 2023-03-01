package controller

import (
	"e-wallet/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

/// bikin Method Handler nya disini

func NewTransactionController(router *gin.Engine, transactionUc usecase.TransactionUseCase) *TransactionController {
	trxController := TransactionController{
		transactionUseCase: transactionUc,
	}

	router.POST("/topup")

	return &trxController
}
