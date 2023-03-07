package controller

import (
	"bytes"
	"e-wallet/model"
	"e-wallet/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyUsers = []model.User{
	{
		Id: utils.GenerateId(),
		Name: "DummyUser01",
		Email: "dummy01-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "Dum01*123",
		Address: "Jl. Dummy No. 01",
		BirthDate: time.Date(1990, time.June, 1, 1, 1, 1, 1, time.UTC),
		ProfilePicture: "DummyImage01",
		CreatedAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
		UpdateAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
	},
}

var dummyUserLogin = []model.UserLogin{
	{
		Email: dummyUsers[0].Email,
		Password: dummyUsers[0].Password,
	},
}

var dummyUserRegister = []model.UserRegister{
	{
		Name: dummyUsers[0].Name,
		Email: dummyUsers[0].Email,
		Password: dummyUsers[0].Password,
		PhoneNumber: dummyUsers[0].PhoneNumber,
		Address: dummyUsers[0].Address,
		BirthDate: dummyUsers[0].BirthDate,
	},
}

var dummyUserUpdate = []model.UserUpdate{
	{
		Id: dummyUsers[0].Id,
		Name: dummyUsers[0].Name,
		Email: dummyUsers[0].Email,
		Password: dummyUsers[0].Password,
		PhoneNumber: dummyUsers[0].PhoneNumber,
		Address: dummyUsers[0].Address,
	},
}

var dummyCheckEmail = []model.CheckEmail{
	{
		Email: dummyUsers[0].Email,
	},
}

type UserUseCaseMock struct {
	mock.Mock
}

type UserControllerTestSuite struct {
	suite.Suite
	routerMock *gin.Engine
	useCaseMock *UserUseCaseMock
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.useCaseMock = new(UserUseCaseMock)
}

func (c *UserUseCaseMock) RegisterUser(input *model.UserRegister) (model.User, error) {
	args := c.Called(input)
	if args.Get(0) != nil {
		return model.User{}, args.Get(0).(error)
	}
	return model.User{}, nil
}

func (c *UserUseCaseMock) Login(input model.UserLogin) (model.User, error) {
	args := c.Called(input)
	if args.Get(1) != nil {
		return model.User{}, args.Get(1).(error)
	}
	return args.Get(0).(model.User), nil	
}

func (c *UserUseCaseMock) UpdateUser(update *model.UserUpdate) (model.User, error) {
	args := c.Called(update)
	if args.Get(0) != nil {
		return model.User{}, args.Get(0).(error)
	}
	return model.User{}, nil
}

func (c *UserUseCaseMock) IsEmailAvailable(input model.CheckEmail) (bool, error) {
	args := c.Called(input)
	if args.Get(1) != nil {
		return false, args.Get(1).(error)
	}
	return true, nil	
}

func (c *UserUseCaseMock) SaveAvatar(id string, fileLocation string) (model.User, error) {
	args := c.Called(id, fileLocation)
	if args.Get(0) != nil {
		return model.User{}, args.Get(0).(error)
	}
	return model.User{}, nil
}

func (suite *UserControllerTestSuite) TestLogin_Success() {
	user := dummyUsers
	login := dummyUserLogin
	suite.useCaseMock.On("Login", login).Return(&user, nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/login", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualUser []model.User
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), 1, len(actualUser))
	assert.Equal(suite.T(), user[0].Name, actualUser[0].Name)
	assert.Nil(suite.T(), err)
}

func (suite *UserControllerTestSuite) TestRegisterUser_Success() {
	newUser := dummyUserRegister[0]
	user := dummyUsers[0]
	suite.useCaseMock.On("RegisterUser", &newUser).Return(nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(newUser)
	request, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualUser model.User
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.Name, actualUser.Name)
}

func (suite *UserControllerTestSuite) TestUpdateUser_Success() {
	userUpdate := dummyUserUpdate[0]
	user := dummyUsers[0]
	suite.useCaseMock.On("UpdateUser", &userUpdate).Return(&user, nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(userUpdate)
	request, _ := http.NewRequest(http.MethodPut, "api/user/update", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualUser model.User
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.Name, actualUser.Name)
}

func (suite *UserControllerTestSuite) TestCheckEmail_Success() {
	user := dummyUsers
	email := dummyCheckEmail
	suite.useCaseMock.On("CheckEmail", email).Return(true, nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/check-email", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualUser []model.User
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), 1, len(actualUser))
	assert.Equal(suite.T(), user[0].Email, actualUser[0].Email)
	assert.Nil(suite.T(), err)
}

func (suite *UserControllerTestSuite) TestUploadAvatar_Success() {
	id := dummyUsers[0].Id
	saveFormat := dummyUsers[0].ProfilePicture
	user := dummyUsers[0]
	suite.useCaseMock.On("UploadAvatar", &id, &saveFormat).Return(nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(&user)
	request, _ := http.NewRequest(http.MethodPost, "api/user/avatar/:id", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualUser model.User
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.ProfilePicture, actualUser.ProfilePicture)
}

func (suite *UserControllerTestSuite) TestLogout_Success() {
	NewUserController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "api/user/logout", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualUser []model.User
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}