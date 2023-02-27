package model

type Balances struct {
	UserId  string `json:"user_id" db:"user_id"`
	Balance int    `json:"balance" db:"balance"`
}
