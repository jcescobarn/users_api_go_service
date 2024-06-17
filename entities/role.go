package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id          uint
	RoleName    string
	Description string
	UserPerRole []UserRoles
}

func (Role) TableName() string {
	return "roles"
}
