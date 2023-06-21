package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskByID(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// 渡されたユーザIDと一致するユーザIDのタスク一覧を取得
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("create_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// 渡されたユーザIDと一致するユーザIDのタスク一覧の抽出
func (tr *taskRepository) GetTaskByID(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// タスクの作成
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// タスクの更新
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	//Model…taskのポインターを渡す　Clauses…ポインターの先の値を書き換える　Where…条件を指定　Update…更新するカラムを指定
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	//RowsAffected…更新された行数を返す
	//0の場合はそもそも存在していない
	if result.RowsAffected < 1 {
		return fmt.Errorf("object not found")
	}
	return nil
}

// タスクの削除
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object not found")
	}
	return nil
}
