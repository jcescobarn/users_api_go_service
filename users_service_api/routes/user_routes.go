package routes

import (
	"users_service_api/handlers"

	"github.com/gin-gonic/gin"
)

type UserRoutesInterface interface {
	routes(routes *gin.Engine)
}

type UserRoutes struct {
	userHandler *handlers.UserHandler
}

func NewUserRoutes(userHandler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{userHandler: userHandler}
}

func (ur *UserRoutes) GetRoutes(routes *gin.Engine) {
	var user_routes *gin.RouterGroup = routes.Group("api/v1/user")

	user_routes.POST("", ur.userHandler.Create)
	user_routes.GET("", ur.userHandler.GetAll)
	user_routes.GET("/:username", ur.userHandler.GetAll)
	user_routes.PUT("/:username", ur.userHandler.Update)
	user_routes.DELETE("/:username", ur.userHandler.Delete)
}
