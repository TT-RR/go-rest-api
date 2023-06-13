package controller

import (
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {

}

func (tc *taskController) GetTaskById(c echo.Context) error {

}

func (tc *taskController) CreateTask(c echo.Context) error {

}

func (tc *taskController) UpdateTask(c echo.Context) error {

}

func (tc *taskController) DeleteTask(c echo.Context) error {

}
