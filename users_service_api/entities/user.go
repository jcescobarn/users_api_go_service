package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint
	Username string
	Name     string
	Password string
	Email    string
}

func (User) TableName() string {
	return "users"
}
