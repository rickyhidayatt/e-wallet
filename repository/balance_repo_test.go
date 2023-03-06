package repository

import (
	"database/sql/driver"
	"e-wallet/model"
	"errors"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyBalances = []model.Balances{
	{
		UserId: dummyUsers[0].Id,
		Balance: 100000,
	},
	{
		UserId: dummyUsers[1].Id,
		Balance: 150000,
	},
}

type BalanceRepositoryTestSuite struct {
	suite.Suite
	mockDb *sqlx.DB
	mockSql sqlmock.Sqlmock
}

func (suite *BalanceRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDb, "sqlmock")
	if err != nil {
		log.Fatalln("Error when opening db connection")
	}
	suite.mockDb = sqlxDB
	suite.mockSql = mockSql
}

func (suite *BalanceRepositoryTestSuite) TestSaveNewBalance_Success() {
	dummyBalance := dummyBalances[0]
	query := `INSERT INTO balances \(user_id, balance\) VALUES \(\?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.UserId, dummyBalance.Balance,).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.SaveNewBalance(dummyBalance)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestSaveNewBalance_Failed() {
	dummyBalance := dummyBalances[0]
	query := `INSERT INTO balances \(user_id, balance\) VALUES \(\?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.UserId, dummyBalance.Balance,).WillReturnError(errors.New("Failed"))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.SaveNewBalance(dummyBalance)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestAddBalance_Success() {
	dummyBalance := dummyBalances[0]
	query := `UPDATE balances SET balance = balance + \? WHERE user_id = \?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.Balance, dummyBalance.UserId,).WillReturnResult(driver.RowsAffected(1))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.AddBalance(dummyBalance.UserId, dummyBalance.Balance)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestAddBalance_Failed() {
	dummyBalance := dummyBalances[0]
	query := `UPDATE balances SET balance = balance + \? WHERE user_id = \?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.Balance, dummyBalance.UserId,).WillReturnError(errors.New("Failed"))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.AddBalance(dummyBalance.UserId, dummyBalance.Balance)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestSendBalance_Success() {
	dummyBalance := dummyBalances[0]
	query := `UPDATE balances SET balance = balance - \? WHERE user_id = \?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.Balance, dummyBalance.UserId,).WillReturnResult(driver.RowsAffected(1))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.SendBalance(dummyBalance.UserId, dummyBalance.Balance)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestSendBalance_Failed() {
	dummyBalance := dummyBalances[0]
	query := `UPDATE balances SET balance = balance - \? WHERE user_id = \?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyBalance.Balance, dummyBalance.UserId,).WillReturnError(errors.New("Failed"))
	repo := NewBalanceRepository(suite.mockDb)
	err := repo.SendBalance(dummyBalance.UserId, dummyBalance.Balance)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestGetBalance_Success() {
	dummyBalance := dummyBalances[0]
	rows := sqlmock.NewRows([]string{"user_id", "balance",})
	rows.AddRow(dummyBalance.UserId, dummyBalance.Balance,)
	query := "SELECT * FROM balances WHERE user_id = $1"
	suite.mockSql.ExpectQuery(query).WillReturnRows(rows)
	repo := NewBalanceRepository(suite.mockDb)
	actual, err := repo.GetBalance(dummyBalance.UserId)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TestGetBalance_Failed() {
	dummyBalance := dummyBalances[0]
	rows := sqlmock.NewRows([]string{"user_id", "balance",})
	rows.AddRow(dummyBalance.UserId, dummyBalance.Balance,)
	query := "SELECT * FROM balances WHERE user_id = $1"
	suite.mockSql.ExpectQuery(query).WillReturnError(errors.New("Failed"))
	repo := NewBalanceRepository(suite.mockDb)
	actual, err := repo.GetBalance(dummyBalance.UserId)
	func() {
		defer func() {
			if r := recover(); r == nil {
				assert.Error(suite.T(), err)
			}
		}()
		repo.GetBalance(dummyBalance.UserId)
	}()
	assert.NotEqual(suite.T(), dummyBalance, actual)
	assert.Error(suite.T(), err)
} // This method has passed unit test

func (suite *BalanceRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}