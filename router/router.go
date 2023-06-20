package router

import (
	"go-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	//エンドポイントの追加
	e.POST("/signup", uc.SignUp)
	//SingUpのエンドポイントにリクエストが来たら、SignUpメソッドを実行する
	e.POST("/login", uc.Login)
	//Loginのエンドポイントにリクエストが来たら、Loginメソッドを実行する
	e.POST("/logout", uc.Logout)
	//Logoutのエンドポイントにリクエストが来たら、Logoutメソッドを実行する

	//tasksのエンドポイントにリクエストが来たら、echojwt.WithConfigを実行する
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	return e
}
