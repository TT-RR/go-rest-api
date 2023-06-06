package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//設定ファイルの読み込み
	var flag_env = flag.String("GO_ENV", "", "開発環境フラグ")
	flag.Parse()
	if *flag_env == "dev" {
		err := godotenv.Load()
		if err != nil {
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
}
