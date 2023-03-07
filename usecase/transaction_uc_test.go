package usecase

import (
	"e-wallet/model"
	"e-wallet/utils"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type transactionRepoMock struct {
	mock.Mock
}

func (repo *transactionRepoMock) PrintHistoryTransactions(userId string) ([]model.TransactionReceiver, error) {
	var history []model.TransactionReceiver
	args := repo.Called(userId)
	if args[0] != nil {
		return nil, args.Error(0)
	}
	return history, nil
}

func (repo *transactionRepoMock) SaveTransaction(trx *model.Transaction) error {
	args := repo.Called(trx)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

func (repo *transactionRepoMock) SaveReceiver(trx *model.Receiver) error {
	args := repo.Called(trx)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

type balanceRepoMock struct {
	mock.Mock
}

func (repo *balanceRepoMock) GetBalance(userId string) ([]int, error) {
	var balanceInt []int
	args := repo.Called(userId)
	if args[0] != nil {
		return nil, args.Error(0)
	}
	return balanceInt, nil
}

func (repo *balanceRepoMock) SaveNewBalance(balance model.Balances) error {
	args := repo.Called(balance)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

func (repo *balanceRepoMock) AddBalance(userId string, amount int) error {
	args := repo.Called(userId, amount)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

func (repo *balanceRepoMock) SendBalance(userId string, amount int) error {
	args := repo.Called(userId, amount)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

type receiverRepoMock struct {
	mock.Mock
}

func (repo *receiverRepoMock) GetReceiverById(receiverId string) (model.Receiver, error) {
	args := repo.Called(receiverId)
	if args[0] != nil {
		return model.Receiver{}, args.Error(0)
	}
	return model.Receiver{}, nil
}

func (suite *TransactionUseCaseTestSuite) TestPrintHistoryTransactionById_Success() {
	var dummySliceHistory []model.TransactionReceiver
	var dummyUserHistory model.TransactionReceiver
	var dummyUserId = utils.GenerateId()
	dummyUserHistory.Id = utils.GenerateId()
	dummyUserHistory.UserId = dummyUserId
	dummyUserHistory.TransactionType = ""
	dummyUserHistory.Amount = 10000
	dummyUserHistory.Category = ""
	dummyUserHistory.TransactionDate = time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC)
	dummyUserHistory.ReceiverName = "DummyUser02"
	dummyUserHistory.ReceiverBank = "DummyBank"
	dummyUserHistory.ReceiverAccount = "9009871234"
	dummyUserHistory.ReceiverId = utils.GenerateId()
	dummySliceHistory = append(dummySliceHistory, dummyUserHistory)	
	suite.trxRepoMock.On("PrintHistoryTransactions", dummySliceHistory[0].UserId).Return(nil)
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.PrintHistoryTransactionsById(dummySliceHistory[0].UserId)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionUseCaseTestSuite) TestPrintHistoryTransactionById_Failed() {
	var dummySliceHistory []model.TransactionReceiver
	var dummyUserHistory model.TransactionReceiver
	var dummyUserId = utils.GenerateId()
	dummyUserHistory.Id = utils.GenerateId()
	dummyUserHistory.UserId = dummyUserId
	dummyUserHistory.TransactionType = ""
	dummyUserHistory.Amount = 10000
	dummyUserHistory.Category = ""
	dummyUserHistory.TransactionDate = time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC)
	dummyUserHistory.ReceiverName = "DummyUser02"
	dummyUserHistory.ReceiverBank = "DummyBank"
	dummyUserHistory.ReceiverAccount = "9009871234"
	dummyUserHistory.ReceiverId = utils.GenerateId()
	dummySliceHistory = append(dummySliceHistory, dummyUserHistory)	
	suite.trxRepoMock.On("PrintHistoryTransactions", dummySliceHistory[0].UserId).Return(errors.New("Failed"))
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.PrintHistoryTransactionsById(dummySliceHistory[0].UserId)
	assert.NotNil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionUseCaseTestSuite) TestSendMoney_Success() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
	}
	dummyUser01 := dummyUser.Id
	dummyUser02 := dummyUser.Id
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(nil)
	var dummyUserBalance = model.Balances{
		UserId: dummyUser01,
		Balance: 20000,
	}
	suite.blcRepoMock.On("GetBalance", dummyUserBalance.UserId).Return(nil)
	var dummyReceiver = model.Receiver{
		Id: utils.GenerateId(),
		UserId: dummyUser02,
		Name: "DummyUser02",
		BankName: "DummyBank",
		AccountNumber: "90098761234",
	}
	suite.trxRepoMock.On("SaveReceiver", &dummyReceiver).Return(nil)
	var dummyTransaction = model.Transaction{
		Id: dummyReceiver.Id,
		UserId: dummyUser01,
		TransactionDate: time.Now(),
		TransactionType: "Send E-Money",
		Amount: 10000,
		ReciverId: dummyReceiver.UserId,
		Category: "Direct",
	}
	suite.trxRepoMock.On("SaveTransaction", dummyTransaction).Return(nil)
	var dummyUserBalanceRcv = model.Balances{
		UserId: dummyReceiver.UserId,
		Balance: dummyTransaction.Amount,
	}
	suite.blcRepoMock.On("SendBalance", dummyUserBalanceRcv.UserId, dummyUserBalanceRcv.Balance).Return(nil)
	var dummyUserTrxSend = model.TransactionSend{
		UserId: dummyUser.Id,
		Amount: dummyTransaction.Amount,
		BankName: dummyReceiver.BankName,
		AccountNumber: dummyReceiver.AccountNumber,
		Category: dummyTransaction.Category,
		ReceiverName: dummyReceiver.Name,
	}
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.SendMoney(dummyUserTrxSend)
	assert.Nil(suite.T(), err)
} // This has not passed unit test

func (suite *TransactionUseCaseTestSuite) TestSendMoney_Failed() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
	}
	dummyUser01 := dummyUser.Id
	dummyUser02 := dummyUser.Id
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(nil)
	var dummyUserBalance = model.Balances{
		UserId: dummyUser01,
		Balance: 20000,
	}
	suite.blcRepoMock.On("GetBalance", dummyUserBalance.UserId).Return(nil)
	var dummyReceiver = model.Receiver{
		Id: utils.GenerateId(),
		UserId: dummyUser02,
		Name: "DummyUser02",
		BankName: "DummyBank",
		AccountNumber: "90098761234",
	}
	suite.trxRepoMock.On("SaveReceiver", &dummyReceiver).Return(nil)
	var dummyTransaction = model.Transaction{
		Id: dummyReceiver.Id,
		UserId: dummyUser01,
		TransactionDate: time.Now(),
		TransactionType: "Send E-Money",
		Amount: 10000,
		ReciverId: dummyReceiver.UserId,
		Category: "Direct",
	}
	suite.trxRepoMock.On("SaveTransaction", dummyTransaction).Return(nil)
	var dummyUserBalanceRcv = model.Balances{
		UserId: dummyReceiver.UserId,
		Balance: dummyTransaction.Amount,
	}
	suite.blcRepoMock.On("SendBalance", dummyUserBalanceRcv.UserId, dummyUserBalanceRcv.Balance).Return(errors.New("Failed"))
	var dummyUserTrxSend = model.TransactionSend{
		UserId: dummyUser.Id,
		Amount: dummyTransaction.Amount,
		BankName: dummyReceiver.BankName,
		AccountNumber: dummyReceiver.AccountNumber,
		Category: dummyTransaction.Category,
		ReceiverName: dummyReceiver.Name,
	}
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.SendMoney(dummyUserTrxSend)
	assert.NotNil(suite.T(), err)
} // This has not passed unit test

func (suite *TransactionUseCaseTestSuite) TestTopUp_Success() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
	}
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(nil)
	var dummyUserBalance = model.Balances{
		UserId: dummyUser.Id,
		Balance: 20000,
	}
	suite.blcRepoMock.On("GetBalance", dummyUserBalance.UserId).Return(nil)
	suite.blcRepoMock.On("SaveNewBalance", dummyUserBalance).Return(nil)
	suite.blcRepoMock.On("AddBalance", dummyUserBalance.UserId, dummyUserBalance.Balance).Return(nil)
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.TopUp(dummyUserBalance.UserId, dummyUserBalance.Balance)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionUseCaseTestSuite) TestTopUp_Failed() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
	}
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(errors.New("Failed"))
	var dummyUserBalance = model.Balances{
		UserId: dummyUser.Id,
		Balance: 20000,
	}
	suite.blcRepoMock.On("GetBalance", dummyUserBalance.UserId).Return(errors.New("Failed"))
	suite.blcRepoMock.On("SaveNewBalance", dummyUserBalance).Return(errors.New("Failed"))
	suite.blcRepoMock.On("AddBalance", dummyUserBalance.UserId, dummyUserBalance.Balance).Return(errors.New("Failed"))
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.TopUp(dummyUserBalance.UserId, dummyUserBalance.Balance)
	assert.NotNil(suite.T(), err)
} // This method has passed unit test

func (suite *TransactionUseCaseTestSuite) TestRequestMoney_Success() {
	var dummyReceiver = model.TransactionRequest{
		UserId: utils.GenerateId(),
		Amount: 10000,
		BankName: "DummyBank",
		AccountNumber: "90012345678",
		Category: "Direct",
		ReceiverID: utils.GenerateId(),
	}
	suite.rcvRepoMock.On("GetReceiverById", dummyReceiver.ReceiverID).Return(nil)
	suite.usrRepoMock.On("GetUserById", dummyReceiver.UserId).Return(nil)
	var dummyTransaction = &model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          dummyReceiver.UserId,
		TransactionDate: time.Now(),
		TransactionType: "Request Money",
		Amount:          dummyReceiver.Amount,
		ReciverId:       dummyReceiver.ReceiverID,
		Category:        dummyReceiver.Category,
	}
	suite.trxRepoMock.On("SaveTransaction", dummyTransaction).Return(nil)
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.RequestMoney(dummyReceiver)
	assert.Nil(suite.T(), err)
} // This method has not passed unit test

func (suite *TransactionUseCaseTestSuite) TestRequestMoney_Failed() {
	var dummyReceiver = model.TransactionRequest{
		UserId: utils.GenerateId(),
		Amount: 10000,
		BankName: "DummyBank",
		AccountNumber: "90012345678",
		Category: "Direct",
		ReceiverID: utils.GenerateId(),
	}
	suite.rcvRepoMock.On("GetReceiverById", dummyReceiver.ReceiverID).Return(nil)
	suite.usrRepoMock.On("GetUserById", dummyReceiver.UserId).Return(nil)
	var dummyTransaction = &model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          dummyReceiver.UserId,
		TransactionDate: time.Now(),
		TransactionType: "Request Money",
		Amount:          dummyReceiver.Amount,
		ReciverId:       dummyReceiver.ReceiverID,
		Category:        dummyReceiver.Category,
	}
	suite.trxRepoMock.On("SaveTransaction", dummyTransaction).Return(errors.New("Failed"))
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	_, err := transactionUseCase.RequestMoney(dummyReceiver)
	assert.NotNil(suite.T(), err)
} // This method has not passed unit test

func (suite *TransactionUseCaseTestSuite) TestGetBonus_Success() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
		BirthDate: time.Date(1990, time.January, 1, 1, 1, 1, 1, time.UTC),
		ProfilePicture: "DummyImage",
		CreatedAt: time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
		UpdateAt: time.Now(),
	}
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(nil)
	var dummyUserBalance = model.Balances{
		UserId: dummyUser.Id,
		Balance: 20000,
	}
	suite.blcRepoMock.On("AddBalance", dummyUserBalance.UserId, dummyUserBalance.Balance).Return(nil)
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	err := transactionUseCase.GetBonus(dummyUser.Id)
	assert.Nil(suite.T(), err)
} // This method has not passed unit test

func (suite *TransactionUseCaseTestSuite) TestGetBonus_Failed() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
		BirthDate: time.Date(1990, time.January, 1, 1, 1, 1, 1, time.UTC),
		ProfilePicture: "DummyImage",
		CreatedAt: time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
		UpdateAt: time.Now(),
	}
	suite.usrRepoMock.On("GetUserById", dummyUser.Id).Return(nil)
	var dummyUserBalance = model.Balances{
		UserId: dummyUser.Id,
		Balance: 20000,
	}
	suite.blcRepoMock.On("AddBalance", dummyUserBalance.UserId, dummyUserBalance.Balance).Return(nil)
	transactionUseCase := NewTransactionUseCase(suite.trxRepoMock, suite.usrRepoMock, suite.blcRepoMock, suite.rcvRepoMock)
	err := transactionUseCase.GetBonus(dummyUser.Id)
	assert.NotNil(suite.T(), err)
} // This method has passed unit test

type TransactionUseCaseTestSuite struct {
	trxRepoMock *transactionRepoMock
	usrRepoMock *userRepoMock
	blcRepoMock *balanceRepoMock
	rcvRepoMock *receiverRepoMock
	suite.Suite
}

func (suite *TransactionUseCaseTestSuite) SetupTest() {
	suite.trxRepoMock = new(transactionRepoMock)
	suite.usrRepoMock = new(userRepoMock)
	suite.blcRepoMock = new(balanceRepoMock)
	suite.rcvRepoMock = new(receiverRepoMock)
}

func TestTransactionUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseTestSuite))
}