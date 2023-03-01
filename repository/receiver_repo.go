package repository

import (
	"e-wallet/model"
	"e-wallet/utils"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type ReceiverRepository interface {
	GetReceiverById(receiverId string) (model.Receiver, error)
}

type receiverRepository struct {
	db *sqlx.DB
}

func (r *receiverRepository) GetReceiverById(receiverId string) (model.Receiver, error) {

	var receivers model.Receiver
	err := r.db.Get(&receivers, utils.GET_RECEIVER_BY_ID, receiverId)
	if err != nil {
		fmt.Println("Masuk sini gak")
		log.Fatal(err)
		return receivers, err
	}

	return receivers, nil
}

func NewReceiverRepository(dbArg *sqlx.DB) ReceiverRepository {
	return &receiverRepository{
		db: dbArg,
	}
}
