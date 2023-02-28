package model

import "time"

type Receiver struct {
	Id            string `db:"id"`
	UserId        string `db:"user_id"`
	Name          string `db:"name"`
	BankName      string `db:"bank_name"`
	AccountNumber string `db:"account_number"`
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
