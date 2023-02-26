package main

import (
	"e-wallet/config"
	"e-wallet/repository"
	"fmt"
)

func main() {
	config := config.NewConfig()
	db := config.DbConnect()

	userRepo := repository.NewUserRepository(db)

	fmt.Println(userRepo.GetUserById("53-331-6070"))
	// txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo)

	txRepo := repository.NewTransactionRepository(db)
	userId := "53-331-6070"
	addBalance := 30000

	err := txRepo.AddBalance(userId, addBalance)
	if err != nil {
		fmt.Println("gagal")
	}
	fmt.Println("berhasil")
}
