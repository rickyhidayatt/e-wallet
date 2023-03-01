package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
)

type UserUseCase interface {
	RegisterUser(input *model.UserRegister) (model.User, error)
	Login(input model.UserLogin) (model.User, error)
	UpdateUser(update *model.UserUpdate) (model.User, error)
	DeleteUserById(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) RegisterUser(input *model.UserRegister) (model.User, error) {
	var user = model.User{}
	user.Id = utils.GenerateId()

	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	user.Address = input.Address
	user.Password = input.Password

	err := u.userRepo.SaveUser(&user)
	if err != nil {
		return user, errors.New("failed register user")
	}

	return user, nil
}

func (u *userUseCase) Login(input model.UserLogin) (model.User, error) {
	email := input.Email
	user := model.User{}

	user.Email = input.Email
	user.Password = input.Password

	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.Id == "" {
		return user, errors.New("no user found on that email")
	}

	if user.Password != input.Password {
		return user, errors.New("password incorrect")
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(update *model.UserUpdate) (model.User, error) {
	var users = model.User{}
	users.Id = update.Id

	user, err := u.userRepo.GetUserById(users.Id)
	if err != nil {
		return users, errors.New("id not found")
	}

	users.Name = update.Name
	users.Email = update.Email
	users.Address = update.Address
	users.Password = update.Password
	users.Address = update.Address

	err = u.userRepo.Update(user)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userUseCase) DeleteUserById(id string) error {
	return u.userRepo.DeleteById(id)
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}
