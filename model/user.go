package model

import "time"

type User struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserResponse struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
