package utils

const (

	// TRANSACTIONS
	INSERT_BALANCE      = "INSERT INTO balance (user_id, balance) VALUES (:user_id, :balance)"
	ADD_BALANCE         = "UPDATE balances SET balance = balance + :balance WHERE user_id = :user_id"
	SEND_BALANCE        = "UPDATE balances SET balance = balance - :balance WHERE user_id = :user_id"
	CHECK_BALANCE_BY_ID = "SELECT * FROM balances WHERE user_id = $1"

	//USERS
	USER_BY_ID       = "SELECT * FROM users WHERE id=$1"
	SELECT_ALL_USER  = "SELECT * FROM users"
	INSERT_NEW_USER  = "INSERT INTO users (name, email, phonenumber, password, address, birthDate) VALUES ($1, $2, $3, $4, $5, $6)"
	UPDATE_USER_BYID = "UPDATE users SET name=$1, email=$2, phonenumber=$3, password=$4, address=$5, birthdate=$6"
)
