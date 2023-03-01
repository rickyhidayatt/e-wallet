package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
	"net/mail"
	"strconv"
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
	passCheck := utils.Password(newUser.Password) 
	if len(newUser.Name) < 4 || len(newUser.Name) > 50 {
		return errors.New("nama berisi 4-50 karakter")
	} else if _, err := mail.ParseAddress(newUser.Email); err != nil{
		return errors.New("format email salah, ex: user@mail.com")
 	} else if len(newUser.PhoneNumber) < 10 || len(newUser.PhoneNumber) > 13 {
		return errors.New("nomor telepon berisi 10-13 karakter")
	} else if _, err := strconv.Atoi(newUser.PhoneNumber); err != nil {
		return errors.New("nomor telepon hanya berisikan angka")
	} else if !passCheck {
		return errors.New("password harus kombinasi uppercase, angka, dan simbol")
	}
	return u.userRepo.CreateNew(newUser)
}

func (u *userUseCase) UpdateUser(user model.User) error {
	passCheck := utils.Password(user.Password) 
	if len(user.Name) < 4 || len(user.Name) > 50 {
		return errors.New("nama berisi 4-50 karakter")
	} else if _, err := mail.ParseAddress(user.Email); err != nil{
		return errors.New("format email salah, ex: user@mail.com")
 	} else if len(user.PhoneNumber) < 10 || len(user.PhoneNumber) > 13 {
		return errors.New("nomor telepon berisi 10-13 karakter")
	} else if _, err := strconv.Atoi(user.PhoneNumber); err != nil {
		return errors.New("nomor telepon hanya berisikan angka")
	} else if !passCheck {
		return errors.New("password harus kombinasi uppercase, angka, dan simbol")
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
