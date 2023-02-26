package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Db *sqlx.DB
}

func (c *Config) initDb() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")

	connectDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open(dbDriver, connectDB)
	// defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Berhasil konnek")

	c.Db = db

}

func (c *Config) DbConnect() *sqlx.DB {
	return c.Db
}

func NewConfig() Config {
	config := Config{}
	config.initDb()
	return config
}
