package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id          uint
	RoleName    string
	Description string
}

func (Role) TableName() string {
	return "roles"
}
