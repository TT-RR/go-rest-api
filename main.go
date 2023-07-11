package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB() //dbの初期化
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	//NewUserRepositoryの起動
	userRepository := repository.NewUserRepository(db)
	//NewTaskRepositoryの起動
	taskRepository := repository.NewTaskRepository(db)

	//NewUserUsecaseの起動
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	//NewTaskUsecaseの起動
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	//NewUserControllerの起動
	userContoller := controller.NewUserController(userUsecase)
	//NewTaskControllerの起動
	taskContoller := controller.NewTaskController(taskUsecase)

	//NewRouterの起動
	e := router.NewRouter(userContoller, taskContoller)
	//サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
