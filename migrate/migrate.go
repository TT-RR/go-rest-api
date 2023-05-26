package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	//NewDB関数を実行し、アドレスを取得
	dbConn := db.NewDB()
	defer fmt.Println("Success migration!!")
	defer db.CloseDB(dbConn)
	//渡したい構造体情報を引数に指定し、DBにテーブルを作成
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
