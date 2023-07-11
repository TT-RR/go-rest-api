package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

// 渡されたユーザIDと一致するユーザIDのタスク一覧を取得
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:       v.ID,
			Title:    v.Title,
			CreateAt: v.CreateAt,
			UpdateAt: v.UpdateAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

// 渡されたユーザIDと一致するユーザIDのタスク一覧の抽出
func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskByID(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:       task.ID,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		UpdateAt: task.UpdateAt,
	}
	return resTask, nil
}

// タスクの作成
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タイトルは1文字以上10文字以下であること
	if err := tu.tv.TaskValidator(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:       task.ID,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		UpdateAt: task.UpdateAt,
	}
	return resTask, nil
}

// タスクの更新
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidator(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:       task.ID,
		Title:    task.Title,
		CreateAt: task.CreateAt,
		UpdateAt: task.UpdateAt,
	}
	return resTask, nil
}

// タスクの削除
func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
