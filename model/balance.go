package model

type Balances struct {
	// Id      string `db:"id"`
	UserId  string `json:"user_id" db:"user_id"`
	Balance int    `json:"balance" db:"balance"`
}
