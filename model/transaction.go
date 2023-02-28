package model

import "time"

type Transaction struct {
	Id              string    `db:"id"`
	UserId          string    `db:"user_id"`
	TransactionDate time.Time `db:"transaction_date"`
	TransactionType string    `db:"transaction_type"`
	Amount          int       `db:"amount"`
	ReciverId       string    `db:"receiver_id"`
	Category        string    `db:"category"`
}
