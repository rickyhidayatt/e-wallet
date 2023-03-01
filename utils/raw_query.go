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

	CHECK_HISTORY_TRANSAKSI = `
	SELECT transactions.id, transactions.user_id, transactions.transaction_type, transactions.amount, transactions.category, transactions.transaction_date, receivers.name, receivers.bank_name, receivers.account_number, receivers.id as receiver_id
	FROM transactions JOIN receivers ON transactions.receiver_id = receivers.id WHERE transactions.user_id = $1
	`
	INSERT_RECEIVER = "INSERT INTO receivers (id, user_id, name, bank_name, account_number) VALUES (:id, :user_id, :name, :bank_name, :account_number)"

	//USERS
	USER_BY_ID      = "SELECT * FROM users WHERE id=$1"
	SELECT_ALL_USER = "SELECT * FROM users"
	INSERT_NEW_USER = `
	INSERT INTO users (id, name, email, phone_number, password, address, birth_date, profile_picture, created_at, update_at)
	VALUES (:id, :name, :email, :phone_number, :password, :address, :birth_date, :profile_picture, :created_at, :update_at)
	`
	UPDATE_USER_BYID = `
	UPDATE users SET name=:name, email=:email, phone_number=:phone_number, password=:password, address=:address, birth_date=:birth_date, profile_picture=:profile_picture, update_at=:update_at 
	WHERE id=:id
	`
	DELETE_USER_BYID = "DELETE FROM users WHERE id = $1"
	FIND_BY_EMAIL    = "SELECT * FROM users WHERE email = $1"

	//RECEIVER
	GET_RECEIVER_BY_ID = "SELECT * FROM receivers WHERE id = $1"
)
