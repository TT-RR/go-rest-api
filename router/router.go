package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	//エンドポイントの追加
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	//CSRFTokenを格納するCookieを作成する
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60,
	}))

	//SingUpのエンドポイントにリクエストが来たら、SignUpメソッドを実行する
	e.POST("/signup", uc.SignUp)
	//Loginのエンドポイントにリクエストが来たら、Loginメソッドを実行する
	e.POST("/login", uc.Login)
	//Logoutのエンドポイントにリクエストが来たら、Logoutメソッドを実行する
	e.POST("/logout", uc.Logout)
	//csrfのエンドポイントにリクエストが来たら、CsrfTokenメソッドを実行する
	e.GET("/csrf", uc.CsrfToken)
	//tasksのエンドポイントにリクエストが来たら、echojwt.WithConfigを実行する
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
