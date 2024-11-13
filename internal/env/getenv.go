package getenv

import (
	"os"

	"github.com/joho/godotenv"
)

var MYSQL_USER, MYSQL_PASS, MYSQL_PORT, DB_NAME string

func Get() error {
	if err := godotenv.Load("./internal/env/.env"); err != nil {
		return err
	}

	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASS = os.Getenv("MYSQL_PASS")
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	DB_NAME = os.Getenv("DB_NAME")

	return nil
}