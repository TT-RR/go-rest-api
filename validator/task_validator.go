package validator

import "go-rest-api/model"

type ITaskValidator interface {
	TaskValidator(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidator(task model.Task) error
