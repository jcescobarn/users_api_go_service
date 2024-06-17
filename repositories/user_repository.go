package repositories

import (
	"users_service_api/config"
	"users_service_api/entities"
	"users_service_api/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db    *config.Database
	utils *utils.Functions
}

func NewUserRepository(db *config.Database, utils *utils.Functions) *UserRepository {
	return &UserRepository{db: db, utils: utils}
}

func (ur *UserRepository) Create(user *entities.User) error {
	var encripted_password string
	var err error
	var result *gorm.DB

	encripted_password, err = ur.utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = encripted_password

	result = ur.db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) GetAll() (*[]entities.User, error) {
	var users *[]entities.User
	var result *gorm.DB

	result = ur.db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	var result *gorm.DB

	result = ur.db.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) Update(user *entities.User) (*entities.User, error) {
	var modified_user *entities.User = user
	var result *gorm.DB
	var encripted_password string
	var err error

	encripted_password, err = ur.utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	modified_user.Password = encripted_password
	result = ur.db.DB.Model(&user).Save(modified_user)
	if result.Error != nil {
		return nil, result.Error
	}

	return modified_user, nil
}

func (ur *UserRepository) Delete(user_id string) (*entities.User, error) {
	var user *entities.User
	var result *gorm.DB

	result = ur.db.DB.Delete(&user, user_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
