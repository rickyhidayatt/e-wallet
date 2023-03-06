package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type BalanceRepository interface {
	AddBalance(userId string, amount int) error
	GetBalance(userId string) ([]int, error)
	SendBalance(userId string, amount int) error
	SaveNewBalance(balance model.Balances) error
}

type balanceRepository struct {
	db *sqlx.DB
}

func (b *balanceRepository) SaveNewBalance(balance model.Balances) error {
	_, err := b.db.NamedExec(utils.INSERT_BALANCE, &balance)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func (b *balanceRepository) AddBalance(userId string, amount int) error {
	var err error

	//rollback
	err = b.runTransaction(func(tx *sqlx.Tx) error {
		if err = b.checkUserExists(tx, userId); err != nil {
			log.Fatal(err)
			return err
		}

		balance := model.Balances{
			UserId:  userId,
			Balance: amount,
		}

		if _, err = tx.NamedExec(utils.ADD_BALANCE, &balance); err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	})

	return err
}

func (b *balanceRepository) SendBalance(userId string, amount int) error {
	var err error

	//rollback
	err = b.runTransaction(func(tx *sqlx.Tx) error {
		if err = b.checkUserExists(tx, userId); err != nil {
			return err
		}

		balance := model.Balances{
			UserId:  userId,
			Balance: amount,
		}

		if _, err = tx.NamedExec(utils.SEND_BALANCE, &balance); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (b *balanceRepository) GetBalance(userId string) ([]int, error) {
	var balances []model.Balances
	var balanceInt []int

	err := b.db.Select(&balances, utils.CHECK_BALANCE_BY_ID, userId)
	if err != nil {
		return nil, err
	}

	for _, v := range balances {
		balanceInt = append(balanceInt, v.Balance)
	}

	return balanceInt, nil
}

func (r *balanceRepository) checkUserExists(tx *sqlx.Tx, userId string) error {
	var user []model.User
	if err := tx.Select(&user, utils.USER_BY_ID, userId); err != nil {
		return err
	}

	if len(user) == 0 {
		return fmt.Errorf("user %s not found", userId)
	}

	return nil
}

func (r *balanceRepository) runTransaction(f func(*sqlx.Tx) error) error {
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

func NewBalanceRepository(dbArg *sqlx.DB) BalanceRepository {
	return &balanceRepository{
		db: dbArg,
	}
}
