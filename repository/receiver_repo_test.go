package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyReceivers = []model.Receiver{
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[0].Id,
		Name: "DummyUser01",
		BankName: "Dummy Bank 1",
		AccountNumber: "900123456789",
	},
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[1].Id,
		Name: "DummyUser02",
		BankName: "Dummy Bank 2",
		AccountNumber: "900987654321",
	},
}

type ReceiverRepositoryTestSuite struct {
	suite.Suite
	mockDb *sqlx.DB
	mockSql sqlmock.Sqlmock
}

func (suite *ReceiverRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDb, "sqlmock")
	if err != nil {
		log.Fatalln("Error when opening db connection")
	}
	suite.mockDb = sqlxDB
	suite.mockSql = mockSql
}

func (suite *ReceiverRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func (suite *ReceiverRepositoryTestSuite) TestGetReceiverById_Success() {
	dummyReceiver := dummyReceivers[0]
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "bank_name", "account_number"})
	rows.AddRow(dummyReceiver.Id, dummyReceiver.UserId, dummyReceiver.Name, dummyReceiver.BankName, dummyReceiver.AccountNumber)
	suite.mockSql.ExpectQuery("SELECT * FROM receivers WHERE id = $1").WithArgs(dummyReceiver.Id).WillReturnRows(rows)
	repo := NewReceiverRepository(suite.mockDb)
	actual, err := repo.GetReceiverById(dummyReceiver.Id)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has not passed unit test

func (suite *ReceiverRepositoryTestSuite) TestGetReceiverById_Failed() {
	dummyReceiver := dummyReceivers[0]
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "bank_name", "account_number"})
	rows.AddRow(dummyReceiver.Id, dummyReceiver.UserId, dummyReceiver.Name, dummyReceiver.BankName, dummyReceiver.AccountNumber)
	suite.mockSql.ExpectQuery("SELECT * FROM receivers WHERE id = $1").WithArgs(dummyReceiver.Id).WillReturnError(errors.New("failed"))
	repo := NewReceiverRepository(suite.mockDb)
	_, err := repo.GetReceiverById(dummyReceiver.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("id receiver not found"), err)
} // This method has passed unit test

func TestReceiverRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ReceiverRepositoryTestSuite))
}