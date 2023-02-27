package main

import (
	"e-wallet/config"
	"e-wallet/repository"
	"e-wallet/usecase"
	"fmt"
)

func main() {
	config := config.NewConfig()
	db := config.DbConnect()

	userRepo := repository.NewUserRepository(db)

	// fmt.Println(userRepo.GetUserById("53-331-6070"))
	txRepo := repository.NewTransactionRepository(db)
	txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo)

	userId := "53-331-6070"
	addBalance := 10000

	// err := txRepo.AddBalance(userId, addBalance)
	// if err != nil {
	// 	fmt.Println("gagal")
	// }

	nilai, err := txUsecase.AddWallet(userId, addBalance)
	if err != nil {
		fmt.Println("gagal")
	}
	fmt.Println("berhasil", nilai)
}
