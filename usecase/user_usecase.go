package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
)

type UserUseCase interface {
	ViewUserById(id string) (model.User, error)
	ViewAllUser(page int, totalRows int) ([]model.User, error)
	CreateNewUser(newUser *model.User) error
	UpdateUser(user model.User) error
	DeleteUserById(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) ViewUserById(id string) (model.User, error) {
	return u.userRepo.ViewById(id)
}

func (u *userUseCase) ViewAllUser(page int, totalRows int) ([]model.User, error) {
	return u.userRepo.ViewAll(page, totalRows)
}

func (u *userUseCase) CreateNewUser(newUser *model.User) error {
	newUser.Id = utils.GenerateId()
	return u.userRepo.CreateNew(newUser)
}

func (u *userUseCase) UpdateUser(user model.User) error {
	if len(user.Name) < 3 || len(user.Name) > 20 {
		return errors.New("nama Minimal 3 Sampai 20 karakter")
	}
	return nil
}

func (u *userUseCase) DeleteUserById(id string) error {
	return u.userRepo.DeleteById(id)
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}
