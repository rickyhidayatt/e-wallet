package usecase

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type transactionRepoMock struct {
	mock.Mock
}

type TransactionUseCaseTestSuite struct {
	repoMock *transactionRepoMock
	suite.Suite
}

func (suite *TransactionUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(transactionRepoMock)
}

func TestTransactionUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseTestSuite))
}