package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	//usecaseのインターフェースに依存させる
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	//usecaseのインターフェースを受け取る
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	//リクエストボディを取得構造体に変換(Bind)
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//usecaseのSignUpメソッドを実行(ユーザの新規作成)
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	//リクエストボディを取得構造体に変換(Bind)
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)                      //cookieの生成
	cookie.Name = "token"                           //cookieの名前
	cookie.Value = tokenString                      //cookieの値
	cookie.Expires = time.Now().Add(24 * time.Hour) //cookieの有効期限
	cookie.Path = "/"                               //cookieのパス
	cookie.Domain = os.Getenv("API_DOMAIN")         //cookieのドメイン
	cookie.Secure = true                            //cookieのセキュア属性
	cookie.HttpOnly = true                          //cookieのHttpOnly属性
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) //cookieをレスポンスに含める
	//cookieを含めたレスポンス
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)              //cookieの生成
	cookie.Name = "token"                   //cookieの名前
	cookie.Value = ""                       //cookieの値
	cookie.Expires = time.Now()             //cookieの有効期限
	cookie.Path = "/"                       //cookieのパス
	cookie.Domain = os.Getenv("API_DOMAIN") //cookieのドメイン
	cookie.Secure = true                    //cookieのセキュア属性
	cookie.HttpOnly = true                  //cookieのHttpOnly属性
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) //cookieをレスポンスに含める
	//cookieを含めたレスポンス
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
