package routes

import (
	"users_service_api/handlers"

	"github.com/gin-gonic/gin"
)

type RoleRoutes struct {
	roleHandler *handlers.RoleHandler
}

func NewRoleRoutes(roleHandler *handlers.RoleHandler) *RoleRoutes {
	return &RoleRoutes{roleHandler: roleHandler}
}

func (rr *RoleRoutes) GetRoutes(routes *gin.Engine) {
	var role_routes *gin.RouterGroup = routes.Group("api/v1/role")

	role_routes.POST("", rr.roleHandler.Create)
	role_routes.GET("", rr.roleHandler.GetAll)
	role_routes.GET("/:name", rr.roleHandler.GetByName)
	role_routes.PUT("/:name", rr.roleHandler.Update)
	role_routes.DELETE("/:name", rr.roleHandler.Delete)
}
