package model

import "time"

type Receiver struct {
	Id            string `db:"id"`
	UserId        string `db:"user_id" json:"user_id"`
	Name          string `db:"name" json:"receiver_name"`
	BankName      string `db:"bank_name" json:"bank_name"`
	AccountNumber string `db:"account_number" json:"account_number"`
}
type TransactionReceiver struct {
	Id              string    `db:"id" json:"id"`
	UserId          string    `db:"user_id" json:"user_id"`
	TransactionType string    `db:"transaction_type" json:"transaction_type"`
	Amount          int       `db:"amount" json:"amount"`
	Category        string    `db:"category" json:"category"`
	TransactionDate time.Time `db:"transaction_date" json:"transaction_date"`
	ReceiverName    string    `db:"name" json:"receiver_name"`
	ReceiverBank    string    `db:"bank_name" json:"receiver_bank"`
	ReceiverAccount string    `db:"account_number" json:"receiver_account"`
	ReceiverId      string    `db:"receiver_id" json:"receiver_id"`
}
