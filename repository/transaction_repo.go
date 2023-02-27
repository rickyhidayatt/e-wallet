package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	SaveTransaction(trx *model.Transaction) error
	SaveReceiver(trx *model.Receiver) error
	PrintHistoryTransactions(userId string) error
}

type transactionRepository struct {
	db *sqlx.DB
}

// BUG nilai belum ke ambil
func (tx *transactionRepository) PrintHistoryTransactions(userId string) error {
	nilai, err := tx.db.Exec(utils.CHECK_HISTORY_TRANSAKSI, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nilai)
	return nil
}

func (tx *transactionRepository) SaveTransaction(trx *model.Transaction) error {
	_, err := tx.db.NamedExec(utils.INSERT_TRANSACTION, &trx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (tx *transactionRepository) SaveReceiver(trx *model.Receiver) error {
	_, err := tx.db.NamedExec(utils.INSERT_RECEIVER, &trx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func NewTransactionRepository(dbArg *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: dbArg,
	}
}
