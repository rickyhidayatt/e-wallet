package usecase

import (
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"errors"
	"log"
	"time"
)

type TransactionUseCase interface {
	TopUp(userId string, addBalance int) (int, error)
	SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) (*model.Transfer, error)
	PrintHistoryTransactionsById(userId string) (error, []model.TransactionReceiver)
	RequestMoney(userId string, amount int, bankName string, accountNumber string, category string, receiverId string) (model.Transaction, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	balanceRepo     repository.BalanceRepository
	receiverRepo    repository.ReceiverRepository
}

func (tx *transactionUseCase) RequestMoney(userId string, amount int, bankName string, accountNumber string, category string, receiverId string) (model.Transaction, error) {
	// cek receiverId ada atau tidak
	var transaction model.Transaction
	_, err := tx.receiverRepo.GetReceiverById(receiverId)
	if err != nil {
		return transaction, errors.New("receiver account id not found")
	}
	// cek userId ada di db ada apa enggak
	_, err = tx.userRepo.GetUserById(userId)
	if err != nil {
		return transaction, errors.New("user id, not found")
	}

	if amount < 500 {
		return transaction, errors.New("the minimum amount to request money is IDR 500")
	} else if receiverId == "" {
		return transaction, errors.New("please fill in the receiver id")
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
		return transaction, errors.New("failed save transaction")
	}

	return transactions, nil
}

func (tx *transactionUseCase) TopUp(userId string, addBalance int) (int, error) {

	checkId, err := tx.userRepo.GetUserById(userId)
	if checkId == nil {
		return 0, errors.New("id not found")
	}

	if addBalance < 10000 {
		return 0, errors.New("the minimum amount for top up is IDR 10.000")
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

func (tx *transactionUseCase) SendMoney(userId string, amount int, bankName string, category string, accountNumber string, receiverName string) (*model.Transfer, error) {

	user, err := tx.userRepo.GetUserById(userId)
	if user == nil {
		return nil, errors.New("failed to get user by id")
	}

	if amount < 5000 {
		return nil, errors.New("the minimum amount is IDR 5.000")
	} else if receiverName == "" {
		return nil, errors.New("please fill in the receiver name")
	}

	balances, err := tx.balanceRepo.GetBalance(userId)
	if err != nil {
		return nil, errors.New("failed to get user balances")
	}

	for _, balance := range balances {
		if balance < amount {
			return nil, errors.New("there are insufficient funds on your account")
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
		return nil, errors.New("failed to save receiver")
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
		return nil, err
	}

	err = tx.balanceRepo.SendBalance(userId, amount)

	if err != nil {
		log.Fatal("failed to send balance", err)
		return nil, err
	}

	return &model.Transfer{
		UserId:        userId,
		Amount:        amount,
		Category:      category,
		BankName:      bankName,
		ReceiverName:  receiverName,
		AccountNumber: accountNumber,
	}, nil
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
