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
	//NewTaskRepositoryの起動
	taskRepository := repository.NewTaskRepository(db)

	//NewUserUsecaseの起動
	userUsecase := usecase.NewUserUsecase(userRepository)
	//NewTaskUsecaseの起動
	taskUsecase := usecase.NewTaskUsecase(taskRepository)

	//NewUserControllerの起動
	userContoller := controller.NewUserController(userUsecase)
	//NewTaskControllerの起動
	taskContoller := controller.NewTaskController(taskUsecase)

	//NewRouterの起動
	e := router.NewRouter(userContoller, taskContoller)
	//サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
