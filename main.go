package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	db := db.NewDB() //dbの初期化
	//NewUserRepositoryの起動
	userRepository := repository.NewUserRepository(db)
	//NewUserUsecaseの起動
	userUsecase := usecase.NewUserUsecase(userRepository)
	//NewUserControllerの起動
	userContoller := controller.NewUserController(userUsecase)
	//NewRouterの起動
	e := router.NewRouter(userContoller)
	//サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
