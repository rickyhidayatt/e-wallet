package model

import "time"

type Transaction struct {
	Id              string    `db:"id"`
	UserId          string    `db:"user_id" json:"user_id" binding:"required"`
	TransactionDate time.Time `db:"transaction_date"`
	TransactionType string    `db:"transaction_type"`
	Amount          int       `db:"amount" json:"amount" binding:"required"`
	ReciverId       string    `db:"receiver_id"`
	Category        string    `db:"category" json:"category"`
}

type Transfer struct {
	UserId        string `json:"user_id" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	Category      string `json:"category"`
	AccountNumber string `json:"account_number" binding:"required"`
	ReceiverName  string `json:"receiver_name" binding:"required"`
}
