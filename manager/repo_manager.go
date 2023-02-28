package manager

import "e-wallet/repository"

type RepositoryManager interface {
	BalanceRepository() repository.BalanceRepository
	ReceiverRepository() repository.ReceiverRepository
	TransactionRepository() repository.TransactionRepository
	UserRepository() repository.UserRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) BalanceRepository() repository.BalanceRepository {
	return repository.NewBalanceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ReceiverRepository() repository.ReceiverRepository {
	return repository.NewReceiverRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.SqlDb())
}

func (r *repositoryManager) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func NewRepositoryManager(inf InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: inf,
	}
}
