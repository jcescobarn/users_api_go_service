package routes

import (
	"users_service_api/handlers"

	middleware "users_service_api/middlewares"

	"github.com/gin-gonic/gin"
)

type RoleRoutes struct {
	roleHandler    *handlers.RoleHandler
	authMidlleware *middleware.AuthMiddleware
}

func NewRoleRoutes(roleHandler *handlers.RoleHandler, authMiddleware *middleware.AuthMiddleware) *RoleRoutes {
	return &RoleRoutes{roleHandler: roleHandler, authMidlleware: authMiddleware}
}

func (rr *RoleRoutes) GetRoutes(routes *gin.Engine) {
	var role_routes *gin.RouterGroup = routes.Group("api/v1/role")

	role_routes.POST("", rr.authMidlleware.Validate(), rr.roleHandler.Create)
	role_routes.GET("", rr.authMidlleware.Validate(), rr.roleHandler.GetAll)
	role_routes.GET("/:name", rr.authMidlleware.Validate(), rr.roleHandler.GetByName)
	role_routes.PUT("/:name", rr.authMidlleware.Validate(), rr.roleHandler.Update)
	role_routes.DELETE("/:name", rr.authMidlleware.Validate(), rr.roleHandler.Delete)
}
