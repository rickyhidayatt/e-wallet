package model

import "time"

type Transaction struct {
	Id              string    `db:"id"`
	UserId          string    `db:"user_id" json:"user_id" binding:"required"`
	TransactionDate time.Time `db:"transaction_date"`
	TransactionType string    `db:"transaction_type"`
	Amount          int       `db:"amount"`
	ReciverId       string    `db:"receiver_id"`
	Category        string    `db:"category"`
}

type Transfer struct {
	UserId        string `json:"user_id" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	Category      string `json:"category"`
	AccountNumber string `json:"account_number" binding:"required"`
	ReceiverName  string `json:"receiver_name"`
}

type TransactionRequest struct {
	UserId        string `json:"user_id" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	Category      string `json:"category" binding:"required"`
	ReceiverID    string `json:"receiver_id" binding:"required"`
}

//perubahan disini
type TransactionSend struct {
	UserId        string `json:"user_id" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	Category      string `json:"category" binding:"required"`
	ReceiverName  string `json:"receiver_name" binding:"required"`
}
