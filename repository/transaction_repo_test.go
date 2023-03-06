package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"errors"
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyTransactions = []model.Transaction{
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[0].Id,
		TransactionDate: time.Now(),
		TransactionType: "Send",
		Amount: 20000,
		ReciverId: dummyReceivers[0].Id,
		Category: "Direct",
	},
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[1].Id,
		TransactionDate: time.Now(),
		TransactionType: "Top Up",
		Amount: 50000,
		ReciverId: dummyReceivers[1].Id,
		Category: "Direct",
	},
}

var dummyTransactionReceivers = []model.TransactionReceiver{
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[0].Id,
		TransactionType: dummyTransactions[0].TransactionType,
		Amount: dummyTransactions[0].Amount,
		Category: dummyTransactions[0].Category,
		TransactionDate: dummyTransactions[0].TransactionDate,
		ReceiverName: dummyReceivers[0].Name,
		ReceiverBank: dummyReceivers[0].BankName,
		ReceiverAccount: dummyReceivers[0].AccountNumber,
		ReceiverId: dummyReceivers[0].Id,
	},
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[1].Id,
		TransactionType: dummyTransactions[1].TransactionType,
		Amount: dummyTransactions[1].Amount,
		Category: dummyTransactions[1].Category,
		TransactionDate: dummyTransactions[1].TransactionDate,
		ReceiverName: dummyReceivers[1].Name,
		ReceiverBank: dummyReceivers[1].BankName,
		ReceiverAccount: dummyReceivers[1].AccountNumber,
		ReceiverId: dummyReceivers[1].Id,
	},
}

type TransactionRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sqlx.DB
	mockSql sqlmock.Sqlmock
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDb, "sqlmock")
	if err != nil {
		log.Fatalln("Error when opening db connection")
	}
	suite.mockDb = sqlxDB
	suite.mockSql = mockSql
}

func (suite *TransactionRepositoryTestSuite) TestSaveTransaction_Success() {
	dummyTransaction := dummyTransactions[0]
	query := `INSERT INTO transactions \(id, user_id, transaction_type, amount, receiver_id, category, transaction_date\) VALUES \(\?, \?, \?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyTransaction.Id, dummyTransaction.UserId, dummyTransaction.TransactionType, dummyTransaction.Amount, dummyTransaction.ReciverId, dummyTransaction.Category, dummyTransaction.TransactionDate,).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewTransactionRepository(suite.mockDb)
	err := repo.SaveTransaction(&dummyTransaction)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TestSaveTransaction_Failed() {
	dummyTransaction := dummyTransactions[0]
	query := `INSERT INTO transactions \(id, user_id, transaction_type, amount, receiver_id, category, transaction_date\) VALUES \(\?, \?, \?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WillReturnError(errors.New("Failed"))
	repo := NewTransactionRepository(suite.mockDb)
	err := repo.SaveTransaction(&dummyTransaction)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TestSaveReceiver_Success() {
	dummyReceiver := dummyReceivers[0]
	query := `INSERT INTO receivers \(id, user_id, name, bank_name, account_number\) VALUES \(\?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyReceiver.Id, dummyReceiver.UserId, dummyReceiver.Name, dummyReceiver.BankName, dummyReceiver.AccountNumber,).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewTransactionRepository(suite.mockDb)
	err := repo.SaveReceiver(&dummyReceiver)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TestSaveReceiver_Failed() {
	dummyReceiver := dummyReceivers[0]
	query := `INSERT INTO receivers \(id, user_id, name, bank_name, account_number\) VALUES \(\?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyReceiver.Id, dummyReceiver.UserId, dummyReceiver.Name, dummyReceiver.BankName, dummyReceiver.AccountNumber,).WillReturnError(errors.New("Failed"))
	repo := NewTransactionRepository(suite.mockDb)
	err := repo.SaveReceiver(&dummyReceiver)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TestPrintHistoryTransactions_Success() {
	dummyTransactionReceiver := dummyTransactionReceivers[0]
	rows := sqlmock.NewRows([]string{"id", "user_id", "transaction_type", "amount", "category", "transaction_date", "name", "bank_name", "account_number", "receiver_id"})
	rows.AddRow(dummyTransactionReceiver.Id, dummyTransactionReceiver.UserId, dummyTransactionReceiver.TransactionType, dummyTransactionReceiver.Amount, dummyTransactionReceiver.Category, dummyTransactionReceiver.TransactionDate, dummyTransactionReceiver.ReceiverName, dummyTransactionReceiver.ReceiverBank, dummyTransactionReceiver.ReceiverAccount, dummyTransactionReceiver.ReceiverId)
	query := `SELECT transactions.id, transactions.user_id, transactions.transaction_type, transactions.amount, transactions.category, transactions.transaction_date, receivers.name, receivers.bank_name, receivers.account_number, receivers.id as receiver_id
	FROM transactions JOIN receivers ON transactions.receiver_id = receivers.id WHERE transactions.user_id = $1`
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewTransactionRepository(suite.mockDb)
	actual, err := repo.PrintHistoryTransactions(dummyTransactionReceiver.UserId)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TestPrintHistoryTransactions_Failed() {
	dummyTransactionReceiver := dummyTransactionReceivers[0]
	rows := sqlmock.NewRows([]string{"id", "user_id", "transaction_type", "amount", "category", "transaction_date", "name", "bank_name", "account_number", "receiver_id"})
	rows.AddRow(dummyTransactionReceiver.Id, dummyTransactionReceiver.UserId, dummyTransactionReceiver.TransactionType, dummyTransactionReceiver.Amount, dummyTransactionReceiver.Category, dummyTransactionReceiver.TransactionDate, dummyTransactionReceiver.ReceiverName, dummyTransactionReceiver.ReceiverBank, dummyTransactionReceiver.ReceiverAccount, dummyTransactionReceiver.ReceiverId)
	query := `SELECT transactions.id, transactions.user_id, transactions.transaction_type, transactions.amount, transactions.category, transactions.transaction_date, receivers.name, receivers.bank_name, receivers.account_number, receivers.id as receiver_id
	FROM transactions JOIN receivers ON transactions.receiver_id = receivers.id WHERE transactions.user_id = $1`
	suite.mockSql.ExpectQuery(query).WillReturnError(errors.New("Failed"))
	repo := NewTransactionRepository(suite.mockDb)
	actual, err := repo.PrintHistoryTransactions(dummyTransactionReceiver.UserId)
	func() {
		defer func() {
			if r := recover(); r == nil {
				assert.Error(suite.T(), err)
			}
		}()
		repo.PrintHistoryTransactions(dummyTransactionReceiver.UserId)
	}()
	assert.NotEqual(suite.T(), dummyTransactionReceiver, actual)
	assert.Error(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}