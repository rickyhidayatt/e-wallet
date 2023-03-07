package controller

import (
	"bytes"
	"e-wallet/model"
	"e-wallet/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyTransaction = []model.Transaction{
	{
		Id: utils.GenerateId(),
		UserId: dummyUsers[0].Id,
		TransactionDate: time.Now(),
		TransactionType: "Top Up",
		Amount: 10000,
		ReciverId: utils.GenerateId(),
		Category: "Direct",
	},
}

var dummyTransactionSend = []model.TransactionSend{
	{
		UserId: dummyUsers[0].Id,
		Amount: dummyTransaction[0].Amount,
		BankName: "Dummy Bank 01",
		AccountNumber: "900987612345",
		Category: dummyTransaction[0].Category,
		ReceiverName: "Dummy Receiver 01",
	},
}

var dummyTransactionRequest = []model.TransactionRequest{
	{
		UserId: dummyUsers[0].Id,
		Amount: dummyTransaction[0].Amount,
		BankName: dummyTransactionSend[0].BankName,
		AccountNumber: dummyTransactionSend[0].AccountNumber,
		Category: dummyTransaction[0].Category,
		ReceiverID: dummyTransaction[0].ReciverId,
	},
}

type TransactionUseCaseMock struct {
	mock.Mock
}

type TransactionControllerTestSuite struct {
	suite.Suite
	routerMock  *gin.Engine
	useCaseMock *TransactionUseCaseMock
}

func (suite *TransactionControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.useCaseMock = new(TransactionUseCaseMock)
}

func (c *TransactionUseCaseMock) TopUp(userId string, addBalance int) (int, error) {
	args := c.Called(userId, addBalance)
	if args.Get(1) != nil {
		return args.Int(0), args.Get(1).(error)
	}
	return args.Int(1), nil
}

func (c *TransactionUseCaseMock) SendMoney(req model.TransactionSend) (*model.Transfer, error) {
	args := c.Called(req)
	if args.Get(0) != nil {
		return &model.Transfer{}, args.Get(0).(error)
	}
	return &model.Transfer{}, nil
}

func (c *TransactionUseCaseMock) PrintHistoryTransactionsById(userId string) ([]model.TransactionReceiver, error) {
	history := []model.TransactionReceiver{}
	args := c.Called(userId)
	if args.Get(0) != nil {
		return history, args.Get(0).(error)
	}
	return history, nil
}

func (c *TransactionUseCaseMock) RequestMoney(req model.TransactionRequest) (model.Transaction, error) {
	args := c.Called(req)
	if args.Get(0) != nil {
		return model.Transaction{}, args.Get(0).(error)
	}
	return model.Transaction{}, nil
}

func (c *TransactionUseCaseMock) GetBonus(userId string) error {
	args := c.Called(userId)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (suite *TransactionControllerTestSuite) TestTopUp_Success() {
	topUp := dummyTransaction[0]
	user := dummyUsers[0]
	suite.useCaseMock.On("TopUp", &user.Id, &topUp.Amount).Return(nil)
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(topUp)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/topup", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualTransaction model.Transaction
	json.Unmarshal([]byte(response), &actualTransaction)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), topUp.Amount, actualTransaction.Amount)
}

func (suite *TransactionControllerTestSuite) TestTopUp_FailedUseCase() {
	topUp := dummyTransaction[0]
	user := dummyUsers[0]
	suite.useCaseMock.On("TopUp", &user.Id, &topUp.Amount).Return(errors.New("failed"))
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(topUp)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/topup", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	var errorResponse struct{ Err string }
	json.Unmarshal([]byte(response), &errorResponse)
	assert.Equal(suite.T(), "failed", errorResponse.Err)
}

func (suite *TransactionControllerTestSuite) TestTransfer_Success() {
	transfer := dummyTransactionSend[0]
	suite.useCaseMock.On("SendMoney", &transfer).Return(nil)
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(transfer)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/transfer", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualTransaction model.Transaction
	json.Unmarshal([]byte(response), &actualTransaction)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), transfer.UserId, actualTransaction.UserId)
}

func (suite *TransactionControllerTestSuite) TestTransfer_FailedUseCase() {
	transfer := dummyTransactionSend[0]
	suite.useCaseMock.On("SendMoney", &transfer).Return(errors.New("failed"))
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(transfer)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/transfer", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	var errorResponse struct{ Err string }
	json.Unmarshal([]byte(response), &errorResponse)
	assert.Equal(suite.T(), "failed", errorResponse.Err)
}

func (suite *TransactionControllerTestSuite) TestRequestMoney_Success() {
	trxRequest := dummyTransactionRequest[0]
	suite.useCaseMock.On("RequestMoney", &trxRequest).Return(nil)
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(trxRequest)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/request", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualTransaction model.Transaction
	json.Unmarshal([]byte(response), &actualTransaction)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), trxRequest.ReceiverID, actualTransaction.ReciverId)
}

func (suite *TransactionControllerTestSuite) TestRequestMoney_FailedUseCase() {
	trxRequest := dummyTransactionRequest[0]
	suite.useCaseMock.On("RequestMoney", &trxRequest).Return(errors.New("failed"))
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(trxRequest)
	request, _ := http.NewRequest(http.MethodPost, "api/transaction/request", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	var errorResponse struct{ Err string }
	json.Unmarshal([]byte(response), &errorResponse)
	assert.Equal(suite.T(), "failed", errorResponse.Err)
}

func (suite *TransactionControllerTestSuite) TestPrintTransactionHistory_Success() {
	transaction := dummyTransaction[0]
	suite.useCaseMock.On("PrintHistoryTransactionById", &transaction.UserId).Return(nil)
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(transaction)
	request, _ := http.NewRequest(http.MethodGet, "api/transaction/:id/history", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualTransaction model.Transaction
	json.Unmarshal([]byte(response), &actualTransaction)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), transaction.UserId, actualTransaction.UserId)
}

func (suite *TransactionControllerTestSuite) TestPrintTransactionHistory_FailedUseCase() {
	transaction := dummyTransaction[0]
	suite.useCaseMock.On("PrintHistoryTransactionById", &transaction.UserId).Return(errors.New("failed"))
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(transaction)
	request, _ := http.NewRequest(http.MethodGet, "api/transaction/:id/history", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	var errorResponse struct{ Err string }
	json.Unmarshal([]byte(response), &errorResponse)
	assert.Equal(suite.T(), "failed", errorResponse.Err)
}

func (suite *TransactionControllerTestSuite) TestGiftBirthDay_Success() {
	user := dummyUsers[0]
	suite.useCaseMock.On("GetBonus", &user.Id).Return(nil)
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodGet, "api/transaction/:id/gift", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualUser model.User
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.BirthDate, actualUser.BirthDate)
}

func (suite *TransactionControllerTestSuite) TestGiftBirthDay_FailedUseCase() {
	user := dummyUsers[0]
	suite.useCaseMock.On("GetBonus", &user.Id).Return(errors.New("failed"))
	NewTransactionController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodGet, "api/transaction/:id/history", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	var errorResponse struct{ Err string }
	json.Unmarshal([]byte(response), &errorResponse)
	assert.Equal(suite.T(), "failed", errorResponse.Err)
}

func TestTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}