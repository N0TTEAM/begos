package model1

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `validate:"required"`
	Email     string `validate:"required"`
	GoogleID  string `gorm:"type:varchar(255);unique;not null"`
	Role      string `gorm:"type:enum('admin','customer');default:'customer'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
