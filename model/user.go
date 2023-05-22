package model

import "time"

type User struct {
	Id       uint      `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
