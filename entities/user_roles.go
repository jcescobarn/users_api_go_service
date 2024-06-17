package entities

import (
	"gorm.io/gorm"
)

type UserRoles struct {
	gorm.Model
	Id     uint
	RoleId uint
	UserId uint
}

func (UserRoles) TableName() string {
	return "user_roles"
}
