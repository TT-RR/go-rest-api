package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskByID(userId uint, taskId uint) ([]model.TaskResponse, error)
	CreateTask(task model.Task) ([]model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) ([]model.TaskResponse, error)
	DeleteTask(uesrId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUseCase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {

}

func (tu *taskUsecase) GetTaskByID(userId uint, taskId uint) ([]model.TaskResponse, error) {

}

func (tu *taskUsecase) CreateTask(task model.Task) ([]model.TaskResponse, error) {

}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) ([]model.TaskResponse, error) {

}

func (tu *taskUsecase) DeleteTask(uesrId uint, taskId uint) error {

}
