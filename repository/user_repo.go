package repository

import (
	"e-wallet/model"
	"e-wallet/utils"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserById(id string) (*model.User, error)
	SaveUser(newUser *model.User) error
	Update(user *model.User) (*model.User, error)
	DeleteById(id string) error
	FindByEmail(email string) (model.User, error)
	SaveAvatar(user *model.User) (model.User, error)
}
type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) GetUserById(id string) (*model.User, error) {
	var user = &model.User{}
	err := r.db.Get(user, utils.USER_BY_ID, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Get(&user, utils.FIND_BY_EMAIL, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) SaveUser(newUser *model.User) error {
	_, err := u.db.NamedExec(utils.INSERT_NEW_USER, &newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *model.User) (*model.User, error) {
	_, err := u.db.NamedExec(utils.UPDATE_USER_BYID, user)

	if err != nil {
		return nil, err
	}

	updatedUser := &model.User{}
	err = u.db.Get(&updatedUser, utils.USER_BY_ID, user.Id)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userRepository) SaveAvatar(user *model.User) (model.User, error) {
	_, err := u.db.NamedExec(utils.UPDATE_USER_BYID, user)

	if err != nil {
		return model.User{}, err
	}

	updatedUser := model.User{}
	err = u.db.Get(&updatedUser, utils.USER_BY_ID, user.Id)

	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil

}

func (u *userRepository) DeleteById(id string) error {
	_, err := u.db.Exec(utils.DELETE_USER_BYID, id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
