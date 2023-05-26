package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//設定ファイルの読み込み
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err == nil {
			log.Fatalln(err)
		}
	}
	//書式の通りにURLを作成
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	//URLを使ってDBに接続
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB Connected Success")
	return db
}

func CloseDB(db *gorm.DB) {
	sqldb, _ := db.DB()
	if err := sqldb.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB Closed Success")
}
