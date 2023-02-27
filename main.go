package main

import (
	"e-wallet/config"
	"e-wallet/repository"
	"fmt"
)

func main() {
	config := config.NewConfig()
	db := config.DbConnect()

	// userRepo := repository.NewUserRepository(db)
	// balanceRepo := repository.NewBalanceRepository(db)

	// test := model.Balances{
	// 	UserId:  "92-363-9555",
	// 	Balance: 250000,
	// }
	// balanceRepo.SaveBalance(test)

	// // fmt.Println(userRepo.GetUserById("53-331-6070"))
	txRepo := repository.NewTransactionRepository(db)
	// txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo, balanceRepo)

	fmt.Println(txRepo.PrintHistoryTransactions("2321"))

	// userId := "2321"
	// addBalance := 10000

	// saldo, err := txUsecase.TopUp(userId, addBalance)
	// // SendMoney(userId, addBalance, "test aja", "menabung", "80280239298", "diki")

	// if err != nil {
	// 	fmt.Println("Gagal")
	// } else {
	// 	fmt.Println("berhasil, saldo kamu :", saldo)
	// }

	//save transaksi
	//=====================
	// saveTransaction := model.Transaction{
	// 	Id:              utils.GenerateId(),
	// 	UserId:          userId,
	// 	TransactionDate: time.Now(),
	// 	TransactionType: "Test transaksi tipe",
	// 	Amount:          addBalance,
	// 	ReciverId:       "b86b9f9bed9e4504a65295f93e9eb23c",
	// 	Category:        "CategoryId1",
	// }

	// err := txUsecase.SendMoney(userId, addBalance, "BCA", "Jajan Pizza Online 2", "2090990909", "reza")
	// if err != nil {
	// 	fmt.Println("gagal transaction")
	// } else {
	// 	fmt.Println("Berhasil")
	// }
	//=====================

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
