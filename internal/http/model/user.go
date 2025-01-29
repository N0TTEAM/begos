package model

import "time"

type User struct {
	Id        uint
	Username  string `validate:"required"`
	Email     string `validate:"required"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
