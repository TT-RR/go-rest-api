package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	//エンドポイントの追加
	e.POST("/signup", uc.SignUp)
	//SingUpのエンドポイントにリクエストが来たら、SignUpメソッドを実行する
	e.POST("/login", uc.Login)
	//Loginのエンドポイントにリクエストが来たら、Loginメソッドを実行する
	e.POST("/logout", uc.Logout)
	//Logoutのエンドポイントにリクエストが来たら、Logoutメソッドを実行する
	return e
}
