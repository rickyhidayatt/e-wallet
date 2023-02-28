package model

import "time"

type Receiver struct {
	Id            string `db:"id"`
	UserId        string `db:"user_id" json:"user_id" binding:"required`
	Name          string `db:"name" json:"receiver_name" binding:"required`
	BankName      string `db:"bank_name" json:"bank_name" binding:"required`
	AccountNumber string `db:"account_number" json:"account_number" binding:"required`
}

type TransactionReceiver struct {
	Id              string    `db:"id"`
	UserId          string    `db:"user_id"`
	TransactionType string    `db:"transaction_type"`
	Amount          int       `db:"amount"`
	Category        string    `db:"category"`
	TransactionDate time.Time `db:"transaction_date"`
	ReceiverName    string    `db:"name"`
	ReceiverBank    string    `db:"bank_name"`
	ReceiverAccount string    `db:"account_number"`
	ReceiverId      string    `db:"receiver_id"`
}
