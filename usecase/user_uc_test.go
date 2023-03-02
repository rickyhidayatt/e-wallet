package usecase

import (
	"e-wallet/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)
type userRepoMock struct {
	mock.Mock
}

func (repo *userRepoMock) GetUserById(id string) (*model.User, error) {
	args := repo.Called(id)
	if args[0] != nil {
		return nil, args.Error(0)
	}
	return &model.User{}, nil
}

func (repo *userRepoMock) SaveUser(newUser *model.User) error {
	args := repo.Called(newUser)
	if args[0] != nil {
		return errors.New("Failed")
	}
	return nil
}

func (repo *userRepoMock) Update(user *model.User) (*model.User, error) {
	args := repo.Called(user)
	if args[0] != nil {
		return nil, args.Error(0)
	}
	return &model.User{}, nil
}

func (repo *userRepoMock) DeleteById(id string) error {
	args := repo.Called(id)
	if args[0] != nil {
		return args.Error(0)
	}
	return nil
}

func (repo *userRepoMock) FindByEmail(email string) (model.User, error) {
	return model.User{}, nil
}

func (repo *userRepoMock) SaveAvatar(user *model.User) (model.User, error) {
	return model.User{}, nil
}

type UserUseCaseTestSuite struct {
	repoMock *userRepoMock
	suite.Suite
}

func (suite *UserUseCaseTestSuite) TestGetUserById_Update_Success() {
	var dummyUpdateUser = &model.UserUpdate{
		Id: "DU001",
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
	}
	suite.repoMock.On("GetUserById", dummyUpdateUser.Id).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.UpdateUser(dummyUpdateUser)
	assert.Nil(suite.T(), err)	
}

func (suite *UserUseCaseTestSuite) TestGetUserById_Update_Failed() {
	var dummyUpdateUser = &model.UserUpdate{
		Id: "DU001",
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
	}
	suite.repoMock.On("GetUserById", dummyUpdateUser.Id).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.UpdateUser(dummyUpdateUser)
	assert.NotNil(suite.T(), err)	
}

func (suite *UserUseCaseTestSuite) TestSaveUser_Success() {
	var dummyUser = &model.UserRegister{
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
	}
	suite.repoMock.On("SaveUser", dummyUser).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.RegisterUser(dummyUser)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestSaveUser_Failed() {
	var dummyUser = &model.UserRegister{
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
	}
	suite.repoMock.On("SaveUser", dummyUser).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.RegisterUser(dummyUser)
	assert.Error(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(userRepoMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}