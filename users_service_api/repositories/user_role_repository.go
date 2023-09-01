package repositories

import (
	"users_service_api/config"
	"users_service_api/entities"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	db *config.Database
}

func NewUserRoleRepository(db *config.Database) *UserRoleRepository {
	return &UserRoleRepository{db: db}
}

func (urr *UserRoleRepository) Create(user_role *entities.UserRoles) error {

	var result *gorm.DB
	result = urr.db.DB.Create(user_role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (urr *UserRoleRepository) GetAll() (*[]entities.UserRoles, error) {
	var user_roles *[]entities.UserRoles
	var result *gorm.DB

	result = urr.db.DB.Find(&user_roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return user_roles, nil

}

func (urr *UserRoleRepository) GetByUserId(user_id string) (*[]entities.UserRoles, error) {
	var roles *[]entities.UserRoles
	var result *gorm.DB

	result = urr.db.DB.Where(" UserId = ?", user_id).First(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil

}

func (urr *UserRoleRepository) Delete(role_id uint, user_id uint) (*entities.UserRoles, error) {
	var deleted_role *entities.UserRoles
	var result *gorm.DB

	result = urr.db.DB.Where("role_id = ? AND role_user = ?", role_id, user_id).Delete(&deleted_role)
	if result.Error != nil {
		return nil, result.Error
	}

	return deleted_role, nil
}
