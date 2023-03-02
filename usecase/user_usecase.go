package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(input *model.User) (model.User, error)
	Login(email string, password string) (model.User, error)
	UpdateUser(update *model.User) (model.User, error)
	DeleteUserById(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) RegisterUser(input *model.User) (model.User, error) {
	var user = model.User{}
	user.Id = utils.GenerateId()

	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	user.Address = input.Address

	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(Password)

	err = u.userRepo.SaveUser(&user)
	if err != nil {
		return user, errors.New("failed register user")
	}

	return user, nil
}

func (u *userUseCase) Login(email string, password string) (model.User, error) {
	var usernil = model.User{}
	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return usernil, err
	}

	if user.Id == "" {
		return usernil, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return usernil, errors.New("password incorrect")
	}

	if user.Email != email {
		return usernil, errors.New("invalid email")
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(update *model.User) (model.User, error) {
	var updatedUser model.User
	user, err := u.userRepo.GetUserById(update.Id)
	if err != nil {
		return updatedUser, err
	}

	user.Name = update.Name
	user.Email = update.Email
	user.Address = update.Address

	newPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return updatedUser, err
	}
	user.Password = string(newPass)

	err = u.userRepo.Update(user)
	if err != nil {
		return updatedUser, err
	}

	updatedUser = *user

	return updatedUser, nil
}

func (u *userUseCase) DeleteUserById(id string) error {
	return u.userRepo.DeleteById(id)
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}
