package usecase

import (
	"e-wallet/model"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)
type repoMock struct {
	mock.Mock
}

func (uc *repoMock) GetUserById(id string) (*model.User, error) {
	return &model.User{}, nil
}

func (uc *repoMock) ViewAll() ([]model.User, error) {
	return nil, nil
}

func (uc *repoMock) SaveUser(newUser *model.User) error {
	return nil
}

func (uc *repoMock) Update(user *model.User) error {
	return nil
}

func (uc *repoMock) DeleteById(id string) error {
	return nil
}

func (uc *repoMock) FindByEmail(email string) (model.User, error) {
	return model.User{}, nil
}

type UserUseCaseTestSuite struct {
	repoMock *repoMock
	suite.Suite
}

func (suite *UserUseCaseTestSuite) TestGetUserById_Success() {
	
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}