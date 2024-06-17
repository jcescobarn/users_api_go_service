package handlers

import (
	"net/http"
	"strconv"
	"users_service_api/entities"
	"users_service_api/repositories"
	"users_service_api/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
	utils          *utils.Functions
}

type UserRequestBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ModifyUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUserHandler(userRepository *repositories.UserRepository, utils *utils.Functions) *UserHandler {
	return &UserHandler{userRepository: userRepository, utils: utils}
}

func (uh *UserHandler) Create(context *gin.Context) {
	var request_body UserRequestBody
	var err error

	if err = context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *entities.User = &entities.User{
		Username: request_body.Username,
		Name:     request_body.Name,
		Password: request_body.Password,
		Email:    request_body.Email,
	}

	var result error = uh.userRepository.Create(user)
	if result != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User Created", "user": user})
	return
}

func (uh *UserHandler) GetAll(context *gin.Context) {
	var users *[]entities.User
	var err error

	users, err = uh.userRepository.GetAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &users)
	return
}

func (uh *UserHandler) GetByUsername(context *gin.Context) {
	var username string
	var err error
	var user *entities.User

	username = context.Param("username")

	user, err = uh.userRepository.GetUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &user)
	return
}

func (uh *UserHandler) Update(context *gin.Context) {
	var username string
	var err error
	var user *entities.User
	var body ModifyUserRequest
	var modified_user *entities.User

	username = context.Param("username")
	user, err = uh.userRepository.GetUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.BindJSON(&body)
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Password != "" {
		user.Password, err = uh.utils.HashPassword(body.Password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}

	modified_user, err = uh.userRepository.Update(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated", "user": modified_user})
}

func (uh *UserHandler) Delete(context *gin.Context) {
	var username string
	var err error
	var user *entities.User
	var status *entities.User

	username = context.Param("username")
	user, err = uh.userRepository.GetUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	status, err = uh.userRepository.Delete(strconv.Itoa(int(user.Id)))

	context.JSON(http.StatusOK, gin.H{"message": "User deleted", "user": status})
}
