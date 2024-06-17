package routes

import (
	"users_service_api/handlers"
	middleware "users_service_api/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler    *handlers.UserHandler
	authMidlleware *middleware.AuthMiddleware
}

func NewUserRoutes(userHandler *handlers.UserHandler, authMiddleware *middleware.AuthMiddleware) *UserRoutes {
	return &UserRoutes{userHandler: userHandler, authMidlleware: authMiddleware}
}

func (ur *UserRoutes) GetRoutes(routes *gin.Engine) {
	var user_routes *gin.RouterGroup = routes.Group("api/v1/user")

	user_routes.POST("", ur.authMidlleware.Validate(), ur.userHandler.Create)
	user_routes.GET("", ur.authMidlleware.Validate(), ur.userHandler.GetAll)
	user_routes.GET("/:username", ur.authMidlleware.Validate(), ur.userHandler.GetAll)
	user_routes.PUT("/:username", ur.authMidlleware.Validate(), ur.userHandler.Update)
	user_routes.DELETE("/:username", ur.authMidlleware.Validate(), ur.userHandler.Delete)

}
