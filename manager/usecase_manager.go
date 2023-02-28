package manager

import "e-wallet/usecase"

type UseCaseManager interface {
	TransactionUseCase() usecase.TransactionUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(
		u.repoManager.TransactionRepository(),
		u.repoManager.UserRepository(),
		u.repoManager.BalanceRepository(),
		u.repoManager.ReceiverRepository(),
	)
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return u.UserUseCase()
}

func NewUseCaseManager(rm RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: rm,
	}
}
