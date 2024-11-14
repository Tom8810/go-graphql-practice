package getenv

import (
	"os"

	"github.com/joho/godotenv"
)

var MYSQL_USER, MYSQL_PASS, MYSQL_PORT, DB_NAME string

// 引数に.envファイルへのパスを指定する
func Get(envpath string) error {
	if err := godotenv.Load(envpath); err != nil {
		return err
	}

	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASS = os.Getenv("MYSQL_PASS")
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	DB_NAME = os.Getenv("DB_NAME")

	return nil
}