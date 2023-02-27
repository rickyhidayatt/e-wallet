package utils

const (

	// TRANSACTIONS
	INSERT_BALANCE      = "INSERT INTO balances (user_id, balance) VALUES (:user_id, :balance)"
	ADD_BALANCE         = "UPDATE balances SET balance = balance + :balance WHERE user_id = :user_id"
	SEND_BALANCE        = "UPDATE balances SET balance = balance - :balance WHERE user_id = :user_id"
	CHECK_BALANCE_BY_ID = "SELECT * FROM balances WHERE user_id = $1"

	INSERT_TRANSACTION = `
	INSERT INTO transactions (id, user_id, transaction_type, amount, receiver_id, category, transaction_date) 
	VALUES (:id, :user_id, :transaction_type, :amount, :receiver_id, :category, :transaction_date)
`
	INSERT_RECEIVER = "INSERT INTO receivers (id, user_id, name, bank_name, account_number) VALUES (:id, :user_id, :name, :bank_name, :account_number)"

	//USERS
	USER_BY_ID       = "SELECT * FROM users WHERE id=$1"
	SELECT_ALL_USER  = "SELECT * FROM users"
	INSERT_NEW_USER  = "INSERT INTO users (name, email, phonenumber, password, address, birthDate) VALUES ($1, $2, $3, $4, $5, $6)"
	UPDATE_USER_BYID = "UPDATE users SET name=$1, email=$2, phonenumber=$3, password=$4, address=$5, birthdate=$6"

	//
	CHECK_HISTORY_TRANSAKSI = `
	SELECT t.id, t.user_id, t.transaction_type, t.amount, t.category, t.transaction_date, r.name, r.bank_name, r.account_number, r.id
	FROM transactions AS t
	JOIN receivers AS r ON t.user_id = r.user_id
	WHERE t.user_id = $1
	`
)
