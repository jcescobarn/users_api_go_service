package routes

import (
	"users_service_api/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	loginHandler *handlers.LoginHandler
}

func (ar *AuthRoutes) AuthRouter(router *gin.Engine) {

	auth_routes := router.Group("api/v1/auth")

	auth_routes.POST("/login", ar.loginHandler.Login)

}
