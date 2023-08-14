package repositories

import (
	"users_service_api/config"
	"users_service_api/entities"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *config.Database
}

func NewRoleRepository(db *config.Database) *RoleRepository {
	return &RoleRepository{db: db}
}

func (rr *RoleRepository) Create(role *entities.Role) error {

	var result *gorm.DB
	result = rr.db.DB.Create(role)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rr *RoleRepository) GetAll() (*[]entities.Role, error) {
	var roles *[]entities.Role
	var result *gorm.DB

	result = rr.db.DB.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil

}

func (rr *RoleRepository) GetByName(rolename string) (*entities.Role, error) {
	var role *entities.Role
	var result *gorm.DB

	result = rr.db.DB.Where(" rolename = ?", rolename).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil

}

func (rr *RoleRepository) Update(role *entities.Role) (*entities.Role, error) {
	var modified_role *entities.Role
	var result *gorm.DB

	result = rr.db.DB.Model(&role).Save(modified_role)
	if result.Error != nil {
		return nil, result.Error
	}

	return modified_role, nil
}

func (rr *RoleRepository) Delete(role_id string) (*entities.Role, error) {
	var deleted_role *entities.Role
	var result *gorm.DB

	result = rr.db.DB.Delete(&deleted_role, role_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return deleted_role, nil
}
