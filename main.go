package main

import (
	"e-wallet/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Server().Run()
}
