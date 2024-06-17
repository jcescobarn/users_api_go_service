package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"users_service_api/entities"
	"users_service_api/repositories"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleRepository *repositories.RoleRepository
}

type RoleRequestBody struct {
	Description string `json:"description"`
	RoleName    string `json:"role_name"`
}

type ModifyRoleRequest struct {
	Description string `json:"description"`
}

func NewRoleHandler(roleRepository *repositories.RoleRepository) *RoleHandler {
	return &RoleHandler{roleRepository: roleRepository}
}

func (rh *RoleHandler) Create(context *gin.Context) {
	var body RoleRequestBody
	var err error

	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Description == "" || body.RoleName == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "role_name or description empty"})
		return
	}

	var role *entities.Role = &entities.Role{
		Description: body.Description,
		RoleName:    body.RoleName,
	}

	var result error = rh.roleRepository.Create(role)
	if result != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Role Created", "user": role})
	return
}

func (rh *RoleHandler) GetAll(context *gin.Context) {
	var roles *[]entities.Role
	var err error

	roles, err = rh.roleRepository.GetAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &roles)
	return
}

func (rh *RoleHandler) GetByName(context *gin.Context) {
	var name string
	var err error
	var role *entities.Role

	name = context.Param("name")

	role, err = rh.roleRepository.GetByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, &role)
	return
}

func (rh *RoleHandler) Update(context *gin.Context) {
	var name string
	var err error
	var role *entities.Role
	var body ModifyRoleRequest
	var modified_role *entities.Role

	name = context.Param("name")
	role, err = rh.roleRepository.GetByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.BindJSON(&body)
	if body.Description != "" {
		role.Description = body.Description
	}

	modified_role, err = rh.roleRepository.Update(role)
	fmt.Println(modified_role)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"sserror": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Role updated", "role": modified_role})
}

func (rh *RoleHandler) Delete(context *gin.Context) {
	var name string
	var err error
	var role *entities.Role
	var status *entities.Role

	name = context.Param("name")
	role, err = rh.roleRepository.GetByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	status, err = rh.roleRepository.Delete(strconv.Itoa(int(role.Id)))

	context.JSON(http.StatusOK, gin.H{"message": "Role deleted", "role": status})
}
