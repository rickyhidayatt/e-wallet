package main

import (
	"e-wallet/config"
	"e-wallet/model"
	"e-wallet/repository"
	"e-wallet/utils"
	"fmt"
	"time"
)

func main() {
	config := config.NewConfig()
	db := config.DbConnect()

	// userRepo := repository.NewUserRepository(db)

	// fmt.Println(userRepo.GetUserById("53-331-6070"))
	txRepo := repository.NewTransactionRepository(db)
	// txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo)

	userId := "53-331-6070"
	addBalance := 10000

	saveTransaction := model.Transaction{
		Id:              utils.GenerateId(),
		UserId:          userId,
		TransactionDate: time.Now(),
		TransactionType: "Test transaksi tipe",
		Amount:          addBalance,
		ReciverId:       "b86b9f9bed9e4504a65295f93e9eb23c",
		CategoryId:      "CategoryId1",
	}

	err := txRepo.SaveTransaction(&saveTransaction)
	if err != nil {
		fmt.Println("gagal transaction")
	} else {
		fmt.Println("Berhasil")
	}
	//=====================================
	// saveReciver := model.Receiver{
	// 	Id:            utils.GenerateId(),
	// 	UserId:        userId,
	// 	Name:          "Darwin",
	// 	BankName:      "BCA",
	// 	AccountNumber: "random376464834886",
	// }

	// err := txRepo.SaveReceiver(&saveReciver)
	// if err != nil {
	// 	fmt.Println("gagal transaction")
	// } else {
	// 	fmt.Println("Berhasil")
	// }

	//=====================================

	// err := txRepo.AddBalance(userId, addBalance)
	// if err != nil {
	// 	fmt.Println("gagal")
	// }

	// nilai, err := txUsecase.AddWallet(userId, addBalance)
	// if err != nil {
	// 	fmt.Println("gagal")
	// }
	// fmt.Println("berhasil", nilai)
}
