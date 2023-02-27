package model

type Receiver struct {
	Id            string `db:"id"`
	UserId        string `db:"user_id"`
	Name          string `db:"name"`
	BankName      string `db:"bank_name"`
	AccountNumber string `db:"account_number"`
}
