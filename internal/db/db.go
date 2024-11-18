package database

import (
	"fmt"
	getenv "gorm_practice1/internal/env"
	"gorm_practice1/internal/gen"

	// gormModel "gorm_practice1/internal/model" // モデル定義が置かれているパッケージ

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitializeGen(db *gorm.DB) {
    // genのクエリオブジェクトを初期化
    gen.SetDefault(db)
}

// Connect はデータベースとの接続を確立し、オートマイグレーションを実行します
func Connect() error {
    // DSN (Data Source Name) 設定
    dsn := getenv.MYSQL_USER + ":" + getenv.MYSQL_PASS + "@tcp(127.0.0.1:" + getenv.MYSQL_PORT + ")/" + getenv.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
    
    // GORM の設定とデータベース接続
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // ログレベルを設定
    })
    if err != nil {
        return fmt.Errorf("failed to connect database: %v", err)
    }

    // グローバル変数にデータベース接続を設定
    DB = db

    InitializeGen(db)

    return nil
}