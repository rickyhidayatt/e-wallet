package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
	"fmt"
	"log"
	"time"
)

type TransactionUseCase interface {
	TopUp(userId string, addBalance int) (int, error)
	SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) error
	PrintHistoryTransactionsById(userId string) (error, []model.TransactionReceiver)
	RequestMoney(userId string, amount int, bankName string, accountNumber string, category string, receiverId string) error
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	balanceRepo     repository.BalanceRepository
	receiverRepo    repository.ReceiverRepository
}

func (tx *transactionUseCase) RequestMoney(userId string, amount int, bankName string, accountNumber string, category string, receiverId string) error {
	// cek receiverId ada atau tidak
	_, err := tx.receiverRepo.GetReceiverById(receiverId)
	if err != nil {
		return errors.New("receiver account id not found")
	}
	// cek userId ada di db ada apa enggak
	_, err = tx.userRepo.ViewById(userId)
	if err != nil {
		return errors.New("user id, not found")
	}

	transactions := model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          userId,
		TransactionDate: time.Now(),
		TransactionType: "Request Money",
		Amount:          amount,
		ReciverId:       receiverId,
		Category:        category,
	}

	// simpan request transaksi
	err = tx.transactionRepo.SaveTransaction(&transactions)
	if err != nil {
		return err
	}
	// Notifikasi on progress
	fmt.Println("Your request has been sent successfully")

	return nil
}

func (tx *transactionUseCase) TopUp(userId string, addBalance int) (int, error) {

	checkId, err := tx.userRepo.ViewById(userId)

	if checkId.Id == userId {
		fmt.Println("error id gak ada")
		return 0, err
	}

	if addBalance < 10000 {
		fmt.Println("gagal balance")
		return 0, errors.New("specify a Rp.10,000 minimum balance")
	}

	id, err := tx.balanceRepo.GetBalance(userId)
	if err != nil {
		log.Fatal(err)
	}

	//check id exist
	if len(id) == 0 {
		balances := model.Balances{
			UserId:  userId,
			Balance: addBalance,
		}
		err = tx.balanceRepo.SaveNewBalance(balances)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = tx.balanceRepo.AddBalance(userId, addBalance)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
	}

	return addBalance, nil
}

func (tx *transactionUseCase) SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) error {
	user, err := tx.userRepo.ViewById(userId)
	if user.Id == userId {
		log.Fatal("failed to get user by id", err)
		return err
	}

	if amount < 5000 {
		log.Fatal("The minimum amount is 5.000")
		return err
	}

	balances, err := tx.balanceRepo.GetBalance(userId)
	if err != nil {
		log.Fatal("failed to get user balances", err)
		return err
	}

	for _, balance := range balances {
		if balance < amount {
			log.Fatal(" There are insufficient funds on your account; ")
			return err
		}
	}

	receiver := model.Receiver{
		Id:            utils.GenerateId(),
		UserId:        userId,
		Name:          receiverName,
		BankName:      bankName,
		AccountNumber: accountNumber,
	}

	err = tx.transactionRepo.SaveReceiver(&receiver)
	if err != nil {
		return err
	}

	transaction := model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          userId,
		TransactionDate: time.Now(),
		TransactionType: bankName,
		Amount:          amount,
		ReciverId:       receiver.Id,
		Category:        category,
	}
	err = tx.transactionRepo.SaveTransaction(&transaction)

	if err != nil {
		log.Fatal("failed to save transaction", err)
		return err
	}

	err = tx.balanceRepo.SendBalance(userId, amount)

	if err != nil {
		log.Fatal("failed to send balance", err)
		return err
	}

	return nil
}

func (tx *transactionUseCase) PrintHistoryTransactionsById(userId string) (error, []model.TransactionReceiver) {
	var transactionsHistory []model.TransactionReceiver
	trxHistory, err := tx.transactionRepo.PrintHistoryTransactions(userId)
	if err != nil {
		log.Fatal(err)
	}
	transactionsHistory = append(transactionsHistory, trxHistory...)

	return nil, transactionsHistory
}

func NewTransactionUseCase(tx repository.TransactionRepository, usr repository.UserRepository, blc repository.BalanceRepository, rcv repository.ReceiverRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tx,
		userRepo:        usr,
		balanceRepo:     blc,
		receiverRepo:    rcv,
	}
}
