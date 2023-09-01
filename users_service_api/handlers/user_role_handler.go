package handlers

import (
	"net/http"
	"strconv"
	"users_service_api/entities"
	"users_service_api/repositories"

	"github.com/gin-gonic/gin"
)

type UserRoleHandler struct {
	userRoleRepository *repositories.UserRoleRepository
}

type UserRoleRequestBody struct {
	RoleId uint `json:"role_id"`
	UserId uint `json:"user_id"`
}

func NewUserRoleHandler(userRoleRepository *repositories.UserRoleRepository) *UserRoleHandler {
	return &UserRoleHandler{userRoleRepository: userRoleRepository}
}

func (rh *UserRoleHandler) Create(context *gin.Context) {
	var body UserRoleRequestBody
	var err error

	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userRole *entities.UserRoles = &entities.UserRoles{
		UserId: body.UserId,
		RoleId: body.RoleId,
	}

	var result error = rh.userRoleRepository.Create(userRole)
	if result != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Role assigned", "user": userRole})
	return
}

func (urh *UserRoleHandler) GetAll(context *gin.Context) {
	var roles *[]entities.UserRoles
	var err error

	roles, err = urh.userRoleRepository.GetAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &roles)
	return
}

func (urh *UserRoleHandler) GetByUserId(context *gin.Context) {
	var user_id string
	var err error
	var role *[]entities.UserRoles

	user_id = context.Param("user_id")

	role, err = urh.userRoleRepository.GetByUserId(user_id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &role)
	return
}

func (urh *UserRoleHandler) Delete(context *gin.Context) {
	var user_id uint64
	var role_id uint64
	var err error
	var status *entities.UserRoles

	user_id, err = strconv.ParseUint(context.Param("user_id"), 10, 64)
	role_id, err = strconv.ParseUint(context.Param("role_id"), 10, 64)
	status, err = urh.userRoleRepository.Delete(uint(role_id), uint(user_id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Role deleted", "role": status})
}
