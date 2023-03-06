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
	SendMoney(req model.TransactionSend) (*model.Transfer, error)
	PrintHistoryTransactionsById(userId string) ([]model.TransactionReceiver, error)
	RequestMoney(req model.TransactionRequest) (model.Transaction, error)
	GetBonus(userId string) error
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	balanceRepo     repository.BalanceRepository
	receiverRepo    repository.ReceiverRepository
}

func (tx *transactionUseCase) TopUp(userId string, addBalance int) (int, error) {

	checkId, err := tx.userRepo.GetUserById(userId)
	if err != nil {
		return 0, err
	} else if checkId == nil {
		return 0, errors.New("id not found")
	}

	if addBalance < 10000 {
		return 0, errors.New("the minimum amount for top up is IDR 10.000")
	}

	id, err := tx.balanceRepo.GetBalance(userId)
	if err != nil {
		log.Fatal(err)
	}

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

func (tx *transactionUseCase) SendMoney(req model.TransactionSend) (*model.Transfer, error) {
	user, err := tx.userRepo.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.New("failed to get user by id")
	}

	if req.Amount < 5000 {
		return nil, errors.New("the minimum amount is IDR 5.000")
	} else if req.ReceiverName == "" {
		return nil, errors.New("please fill in the receiver name")
	}

	balances, err := tx.balanceRepo.GetBalance(req.UserId)
	if err != nil {
		return nil, errors.New("failed to get user balances")
	}

	for _, balance := range balances {
		if balance < req.Amount {
			return nil, errors.New("your money is not enough")
		}
	}

	receiver := model.Receiver{
		Id:            utils.GenerateId(),
		UserId:        req.UserId,
		Name:          req.ReceiverName,
		BankName:      req.BankName,
		AccountNumber: req.AccountNumber,
	}

	err = tx.transactionRepo.SaveReceiver(&receiver)
	if err != nil {
		return nil, errors.New("failed to save receiver")
	}

	transaction := model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          req.UserId,
		TransactionDate: time.Now(),
		TransactionType: req.BankName,
		Amount:          req.Amount,
		ReciverId:       receiver.Id,
		Category:        req.Category,
	}

	err = tx.transactionRepo.SaveTransaction(&transaction)
	if err != nil {
		log.Fatal("failed to save transaction", err)
		return nil, err
	}

	err = tx.balanceRepo.SendBalance(req.UserId, req.Amount)
	if err != nil {
		log.Fatal("failed to send balance", err)
		return nil, err
	}

	transfer := &model.Transfer{
		UserId:        req.UserId,
		Amount:        req.Amount,
		Category:      req.Category,
		BankName:      req.BankName,
		ReceiverName:  req.ReceiverName,
		AccountNumber: req.AccountNumber,
	}

	return transfer, nil
}

func (tx *transactionUseCase) PrintHistoryTransactionsById(userId string) ([]model.TransactionReceiver, error) {
	var transactionsHistory []model.TransactionReceiver
	if userId == "" {
		return nil, errors.New("fill in your user id")
	}

	trxHistory, err := tx.transactionRepo.PrintHistoryTransactions(userId)
	if err != nil {
		return nil, errors.New("id not found")
	}

	// Don't need second validation, already check valid by id
	/* if len(trxHistory) == 0 {
		return nil, errors.New("not found transaction data for user")
	} */
	
	transactionsHistory = append(transactionsHistory, trxHistory...)
	return transactionsHistory, nil
}

func (tx *transactionUseCase) RequestMoney(req model.TransactionRequest) (model.Transaction, error) {
	var transaction model.Transaction

	_, err := tx.receiverRepo.GetReceiverById(req.ReceiverID)
	if err != nil {
		return transaction, errors.New("receiver account id not found")
	}

	_, err = tx.userRepo.GetUserById(req.UserId)
	if err != nil {
		return transaction, errors.New("user id not found")
	}

	if req.Amount < 500 {
		return transaction, errors.New("the minimum amount to request money is IDR 500")
	} else if req.ReceiverID == "" {
		return transaction, errors.New("please fill in the receiver id")
	}

	transaction = model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          req.UserId,
		TransactionDate: time.Now(),
		TransactionType: "Request Money",
		Amount:          req.Amount,
		ReciverId:       req.ReceiverID,
		Category:        req.Category,
	}

	err = tx.transactionRepo.SaveTransaction(&transaction)
	if err != nil {
		return transaction, errors.New("failed to save transaction")
	}

	return transaction, nil
}

func (tx *transactionUseCase) GetBonus(userId string) error {
	user, err := tx.userRepo.GetUserById(userId)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New("id not found")
	}

	now := time.Now()
	if user.BirthDate.Month() == now.Month() && user.BirthDate.Day() == now.Day() {
		tx.balanceRepo.AddBalance(userId, 25000)
	} else {
		return errors.New("we're sorry, but it's not your birthday yet. Please be patient :)")
	}
	return nil
}

func NewTransactionUseCase(tx repository.TransactionRepository, usr repository.UserRepository, blc repository.BalanceRepository, rcv repository.ReceiverRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tx,
		userRepo:        usr,
		balanceRepo:     blc,
		receiverRepo:    rcv,
	}
}
