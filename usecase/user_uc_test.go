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
		return args.Error(0)
	}
	return nil
}

func (repo *userRepoMock) Update(user *model.User) (model.User, error) {
	args := repo.Called(user)
	if args[0] != nil {
		return *user, args.Error(0)
	}
	return *user, nil
}

func (repo *userRepoMock) FindByEmail(email string) (model.User, error) {
	args := repo.Called(email)
	if args[0] != nil {
		return model.User{}, args.Error(0)
	}
	return model.User{}, nil
}

func (repo *userRepoMock) SaveAvatar(user *model.User) (model.User, error) {
	args := repo.Called(user)
	if args[0] != nil {
		return model.User{}, args.Error(0)
	}
	return model.User{}, nil
}

type UserUseCaseTestSuite struct {
	repoMock *userRepoMock
	suite.Suite
}

func (suite *UserUseCaseTestSuite) TestRegisterUser_Success() {
	var dummyUserInput = model.UserRegister{
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
		BirthDate: time.Date(1990, time.January, 1, 1, 1, 1, 1, time.UTC),
	}
	var dummyNewUser = model.User{
		Id: utils.GenerateId(),
		Name: dummyUserInput.Name,
		Email: dummyUserInput.Email,
		PhoneNumber: dummyUserInput.PhoneNumber,
		Password: dummyUserInput.Password,
		Address: dummyUserInput.Address,
		BirthDate: dummyUserInput.BirthDate,
		ProfilePicture: "DummyImage",
		CreatedAt: time.Now(),
		UpdateAt: time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC),
	}
	suite.repoMock.On("SaveUser", &dummyNewUser).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.RegisterUser(&dummyUserInput)
	assert.Nil(suite.T(), err)
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestRegisterUser_Failed() {
	var dummyUserInput = model.UserRegister{
		Name: "DummyUser01",
		Email: "dummy-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "DummyPass*123",
		Address: "Jl. Dummy No. 25",
		BirthDate: time.Date(1990, time.January, 1, 1, 1, 1, 1, time.UTC),
	}
	var dummyNewUser = model.User{
		Id: utils.GenerateId(),
		Name: dummyUserInput.Name,
		Email: dummyUserInput.Email,
		PhoneNumber: dummyUserInput.PhoneNumber,
		Password: dummyUserInput.Password,
		Address: dummyUserInput.Address,
		BirthDate: dummyUserInput.BirthDate,
		ProfilePicture: "DummyImage",
		CreatedAt: time.Now(),
		UpdateAt: time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC),
	}
	suite.repoMock.On("SaveUser", &dummyNewUser).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.RegisterUser(&dummyUserInput)
	assert.NotNil(suite.T(), err)
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
		Email: "dummy-email@mail.com",
		Password: "DummyPass*123",
	}
	suite.repoMock.On("FindByEmail", dummyUser.Email).Return(nil)
	var dummyUserLogin = model.UserLogin{
		Email: dummyUser.Email,
		Password: dummyUser.Password,
	}
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.Login(dummyUserLogin)
	assert.Nil(suite.T(), err)
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestLogin_Failed() {
	var dummyUser = model.User{
		Id: utils.GenerateId(),
		Email: "dummy-email@mail.com",
		Password: "DummyPass*123",
	}
	suite.repoMock.On("FindByEmail", dummyUser.Email).Return(errors.New("Failed"))
	var dummyUserLogin = model.UserLogin{
		Email: dummyUser.Email,
		Password: dummyUser.Password,
	}
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.Login(dummyUserLogin)
	assert.NotNil(suite.T(), err)
} // This method has passed unit test

func (suite *UserUseCaseTestSuite) TestUpdateUser_Success() {
	var dummyLastUser = model.User{
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
	suite.repoMock.On("GetUserById", dummyLastUser.Id).Return(nil)
	var dummyUpdatedUser = model.UserUpdate{
		Id: dummyLastUser.Id,
		Name: dummyLastUser.Name,
		Email: dummyLastUser.Email,
		PhoneNumber: dummyLastUser.PhoneNumber,
		Password: dummyLastUser.Password,
		Address: dummyLastUser.Address,
	}
	suite.repoMock.On("Update", &dummyLastUser).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.UpdateUser(&dummyUpdatedUser)
	assert.Nil(suite.T(), err)	
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestUpdateUser_Failed() {
	var dummyLastUser = model.User{
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
	suite.repoMock.On("GetUserById", dummyLastUser.Id).Return(nil)
	var dummyUpdatedUser = model.UserUpdate{
		Id: dummyLastUser.Id,
		Name: dummyLastUser.Name,
		Email: dummyLastUser.Email,
		PhoneNumber: dummyLastUser.PhoneNumber,
		Password: dummyLastUser.Password,
		Address: dummyLastUser.Address,
	}
	suite.repoMock.On("Update", &dummyLastUser).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.UpdateUser(&dummyUpdatedUser)
	assert.NotNil(suite.T(), err)	
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestIsEmailAvailable_Success() {
	var dummyUserEmail = model.CheckEmail{
		Email: "dummy-email@mail.com",
	}
	var dummyUserEmailRepo = model.User{
		Email: dummyUserEmail.Email,
	}
	suite.repoMock.On("FindByEmail", dummyUserEmailRepo.Email).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.IsEmailAvailable(dummyUserEmail)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *UserUseCaseTestSuite) TestIsEmailAvailable_Failed() {
	var dummyUserEmail = model.CheckEmail{
		Email: "dummy-email@mail.com",
	}
	var dummyUserEmailRepo = model.User{
		Email: dummyUserEmail.Email,
	}
	suite.repoMock.On("FindByEmail", dummyUserEmailRepo.Email).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.IsEmailAvailable(dummyUserEmail)
	assert.NotNil(suite.T(), err)
} // This method has passed unit test

func (suite *UserUseCaseTestSuite) TestSaveAvatar_Success() {
	var dummyUserAvatar = model.User{
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
	suite.repoMock.On("GetUserById", dummyUserAvatar.Id).Return(nil)
	suite.repoMock.On("SaveAvatar", &dummyUserAvatar).Return(nil)
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.SaveAvatar(dummyUserAvatar.Id, dummyUserAvatar.ProfilePicture)
	assert.Nil(suite.T(), err)
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) TestSaveAvatar_Failed() {
	var dummyUserAvatar = model.User{
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
	suite.repoMock.On("GetUserById", dummyUserAvatar.Id).Return()
	suite.repoMock.On("SaveAvatar", &dummyUserAvatar).Return(errors.New("Failed"))
	userUseCase := NewUserUseCase(suite.repoMock)
	_, err := userUseCase.SaveAvatar(dummyUserAvatar.Id, dummyUserAvatar.ProfilePicture)
	assert.NotNil(suite.T(), err)
} // This method has not passed unit test

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(userRepoMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}