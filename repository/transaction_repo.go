package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	AddBalance(userId string, amount int) error
	GetBalance(userId string) ([]int, error)
	SendBalance(userId string, amount int) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (r *transactionRepository) AddBalance(userId string, amount int) error {
	var err error

	//rolback jika ada kesalahan
	err = r.runTransaction(func(tx *sqlx.Tx) error {
		if err = r.checkUserExists(tx, userId); err != nil {
			return fmt.Errorf("failed to add balance: %v", err)
		}

		balance := model.Balances{
			UserId:  userId,
			Balance: amount,
		}

		if _, err = tx.NamedExec(utils.ADD_BALANCE, &balance); err != nil {
			return fmt.Errorf("failed to add balance: %v", err)
		}

		return nil
	})

	return err
}

func (r *transactionRepository) SendBalance(userId string, amount int) error {
	var err error

	//rolback jika ada kesalahan
	err = r.runTransaction(func(tx *sqlx.Tx) error {
		if err = r.checkUserExists(tx, userId); err != nil {
			return fmt.Errorf("failed to send balance: %v", err)
		}

		balance := model.Balances{
			UserId:  userId,
			Balance: amount,
		}

		if _, err = tx.NamedExec(utils.SEND_BALANCE, &balance); err != nil {
			return fmt.Errorf("failed to send balance: %v", err)
		}

		return nil
	})

	return err
}

func (r *transactionRepository) GetBalance(userId string) ([]int, error) {
	var balances []model.Balances
	var balanceInt []int
	var err error

	err = r.db.Select(&balances, utils.CHECK_BALANCE_BY_ID, userId)
	if err != nil {
		return nil, err
	}

	for _, v := range balances {
		balanceInt = append(balanceInt, v.Balance)
	}

	return balanceInt, nil
}

func (r *transactionRepository) checkUserExists(tx *sqlx.Tx, userId string) error {
	var user []model.User
	if err := tx.Select(&user, utils.USER_BY_ID, userId); err != nil {
		return err
	}

	if len(user) == 0 {
		return fmt.Errorf("user %s not found", userId)
	}

	return nil
}

func (r *transactionRepository) runTransaction(f func(*sqlx.Tx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	return f(tx)
}

func NewTransactionRepository(dbArg *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: dbArg,
	}
}
