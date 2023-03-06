package controller

import (
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
	users := dummyUsers
	suite.useCaseMock.On("Login").Return(users, nil)
	NewUserController(suite.routerMock, suite.useCaseMock)
	// ini baru kondisikan HTTP Status
	r := httptest.NewRecorder()
	// request test yang sesuai
	request, err := http.NewRequest(http.MethodGet, "/customer", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualUsers []model.User
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUsers)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), 1, len(actualUsers))
	assert.Equal(suite.T(), users[0].Name, actualUsers[0].Name)
	assert.Nil(suite.T(), err)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}