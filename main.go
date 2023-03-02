package main

import (
	"e-wallet/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Server().Run()

	// db, err := sqlx.Connect("postgres", "user=postgres password=Repsol12 dbname=wallet_app sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // test koneksi database
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Database connected successfully.")

	// // close koneksi database setelah selesai
	// defer db.Close()

	// repo := repository.NewUserRepository(db)
	// use := usecase.NewUserUseCase(repo)

	// id := "d58f7bef0e1b4deb9552ea01a205c8a61"

	// _, err = use.SaveAvatar(id, "/percobaan.img")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {

	// 	fmt.Println("berhasil simpan")
	// }
}
