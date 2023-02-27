package usecase

import (
	"e-wallet/repository"
	"fmt"
)

// BUG Masih Ada
type TransactionUseCase interface {
	AddWallet(userId string, addBalance int) (int, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

func (tx *transactionUseCase) AddWallet(userId string, addBalance int) (int, error) {

	checkId, err := tx.userRepo.GetUserById(userId)

	if checkId == nil {
		fmt.Println("error id gak ada")
		return 0, err
	}
	fmt.Println("di Usecase", checkId)

	if addBalance < 10000 {
		fmt.Println("gagal balance")
		return 0, fmt.Errorf("addBalance amount (%d) exceeds maximum limit of 10000", addBalance)
	}

	// var balance model.Balances
	// fmt.Println("Jumlahnya :", balance.Balance)

	err = tx.transactionRepo.AddBalance(userId, addBalance)
	if err != nil {
		return 0, err
	}

	return addBalance, nil
}

func NewTransactionUseCase(txRepoArg repository.TransactionRepository, userArg repository.UserRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: txRepoArg,
		userRepo:        userArg,
	}
}
