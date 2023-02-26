package model

import "time"

type Transaction struct {
	Id              string
	UserId          string
	TransactionDate time.Time
	TransactionType string
	Amount          int
	SenderId        string
	ReciverId       string
	CategoryId      string
}
