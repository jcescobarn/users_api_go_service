package routes

import (
	"users_service_api/handlers"
	middleware "users_service_api/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoleRoutes struct {
	userRoleHandler *handlers.UserRoleHandler
	authMidlleware  *middleware.AuthMiddleware
}

func NewUserRoleRoutes(userRoleHandler *handlers.UserRoleHandler, authMiddleware *middleware.AuthMiddleware) *UserRoleRoutes {
	return &UserRoleRoutes{userRoleHandler: userRoleHandler, authMidlleware: authMiddleware}
}

func (urr *UserRoleRoutes) GetRoutes(routes *gin.Engine) {
	var user_role_routes *gin.RouterGroup = routes.Group("api/v1/user_role")

	user_role_routes.POST("", urr.authMidlleware.Validate(), urr.userRoleHandler.Create)
	user_role_routes.GET("", urr.authMidlleware.Validate(), urr.userRoleHandler.GetAll)
	user_role_routes.GET("/:user_id", urr.authMidlleware.Validate(), urr.userRoleHandler.GetByUserId)
	user_role_routes.DELETE("", urr.authMidlleware.Validate(), urr.userRoleHandler.Delete)

}
