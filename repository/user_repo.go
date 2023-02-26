package repository

import (
	"e-wallet/model"
	"e-wallet/utils"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserById(id string) ([]model.User, error)
	ViewAll(page int, totalRows int) ([]model.User, error)
	CreateNew(newUser *model.User) error
	Update(user model.User) error
}
type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) GetUserById(id string) ([]model.User, error) {
	var user []model.User
	err := r.db.Select(&user, utils.USER_BY_ID, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) ViewAll(page int, totalRows int) ([]model.User, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var users []model.User
	err := u.db.Select(&users, utils.SELECT_ALL_USER, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) CreateNew(newUser *model.User) error {
	_, err := u.db.Exec(utils.INSERT_NEW_USER, newUser.Name, newUser.Email, newUser.PhoneNumber, newUser.Password, newUser.Address, newUser.BirthDate)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user model.User) error {
	_, err := u.db.NamedExec(utils.UPDATE_USER_BYID, user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
