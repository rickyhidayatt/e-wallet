package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
	"time"
)

type UserUseCase interface {
	RegisterUser(input *model.UserRegister) (model.User, error)
	Login(input model.UserLogin) (model.User, error)
	UpdateUser(update *model.UserUpdate) (model.User, error)
	IsEmailAvailable(input model.CheckEmail) (bool, error)
	SaveAvatar(id string, fileLocation string) (model.User, error)
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
	user.CreatedAt = time.Now()
	user.BirthDate = input.BirthDate

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

	user, err := u.userRepo.GetUserById(update.Id)

	if err != nil {
		return *user, err
	}

	/* if user == nil {
		return nil, errors.New("user not found")
	} */

	if update.Name != "" {
		user.Name = update.Name
	}
	if update.Email != "" {
		user.Email = update.Email
	}
	if update.Address != "" {
		user.Address = update.Address
	}
	if update.Password != "" {
		user.Password = update.Password
	}

	user.UpdateAt = time.Now()

	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (u *userUseCase) IsEmailAvailable(input model.CheckEmail) (bool, error) {
	email := input.Email
	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return false, errors.New("email not found")
	}

	if user.Id == "" {
		return true, nil
	}

	return false, nil
}

func (u *userUseCase) SaveAvatar(id string, fileLocation string) (model.User, error) {

	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	user.ProfilePicture = fileLocation
	updateUser, err := u.userRepo.SaveAvatar(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}
