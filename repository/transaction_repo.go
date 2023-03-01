package repository

import (
	"e-wallet/model"
	"e-wallet/utils"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	SaveTransaction(trx *model.Transaction) error
	SaveReceiver(trx *model.Receiver) error
	PrintHistoryTransactions(userId string) ([]model.TransactionReceiver, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func (tx *transactionRepository) PrintHistoryTransactions(userId string) ([]model.TransactionReceiver, error) {

	var transactions []model.TransactionReceiver

	err := tx.db.Select(&transactions, utils.CHECK_HISTORY_TRANSAKSI, userId)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tx *transactionRepository) SaveTransaction(trx *model.Transaction) error {
	_, err := tx.db.NamedExec(utils.INSERT_TRANSACTION, &trx)
	if err != nil {
		return err
	}

	return nil
}

func (tx *transactionRepository) SaveReceiver(trx *model.Receiver) error {
	_, err := tx.db.NamedExec(utils.INSERT_RECEIVER, &trx)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionRepository(dbArg *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: dbArg,
	}
}
