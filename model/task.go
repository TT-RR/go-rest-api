package model

import "time"

type Task struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	User     User      `json:"user" gorm:"foreignkey:UserID; constraint:OnDelete:CASCADE"`
	UserID   uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
