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

// BUG Masih Ada
type TransactionUseCase interface {
	TopUp(userId string, addBalance int) (int, error)
	SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) error
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	balanceRepo     repository.BalanceRepository
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
		CategoryId:      category,
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

// func (tx *transactionUseCase) AddWallet(userId string, addBalance int) (int, error) {

// 	checkId, err := tx.userRepo.GetUserById(userId)

// 	if checkId == nil {
// 		fmt.Println("error id gak ada")
// 		return 0, err
// 	}
// 	fmt.Println("di Usecase", checkId)

// 	if addBalance < 10000 {
// 		fmt.Println("gagal balance")
// 		return 0, errors.New("specify a Rp.10,000 minimum balance")
// 	}

// 	err = tx.transactionRepo.AddBalance(userId, addBalance)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return addBalance, nil
// }

// func (tx *transactionUseCase) SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) error {
// 	user, err := tx.userRepo.GetUserById(userId)
// 	if user == nil {
// 		log.Fatal("failed to get user by id", err)
// 		return err
// 	}

// 	if amount < 5000 {
// 		log.Fatal("The minimum amount is 5.000")
// 		return err
// 	}

// 	balances, err := tx.transactionRepo.GetBalance(userId)
// 	if err != nil {
// 		log.Fatal("failed to get user balances", err)
// 		return err
// 	}

// 	for _, balance := range balances {
// 		if balance < amount {
// 			log.Fatal(" There are insufficient funds on your account; ")
// 			return err
// 		}
// 	}

// 	receiver := model.Receiver{
// 		Id:            utils.GenerateId(),
// 		UserId:        userId,
// 		Name:          receiverName,
// 		BankName:      bankName,
// 		AccountNumber: accountNumber,
// 	}

// 	err = tx.transactionRepo.SaveReceiver(&receiver)
// 	if err != nil {
// 		return err
// 	}

// 	transaction := model.Transaction{
// 		Id:              utils.GenerateId(),
// 		UserId:          userId,
// 		TransactionDate: time.Now(),
// 		TransactionType: bankName,
// 		Amount:          amount,
// 		ReciverId:       receiver.Id,
// 		CategoryId:      category,
// 	}
// 	err = tx.transactionRepo.SaveTransaction(&transaction)

// 	if err != nil {
// 		log.Fatal("failed to save transaction", err)
// 		return err
// 	}

// 	err = tx.transactionRepo.SendBalance(userId, amount)

// 	if err != nil {
// 		log.Fatal("failed to send balance", err)
// 		return err
// 	}

// 	return nil
// }

func NewTransactionUseCase(txRepoArg repository.TransactionRepository, userArg repository.UserRepository, balanceArg repository.BalanceRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: txRepoArg,
		userRepo:        userArg,
		balanceRepo:     balanceArg,
	}
}
