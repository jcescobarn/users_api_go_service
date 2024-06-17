package routes

import (
	"users_service_api/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	loginHandler *handlers.LoginHandler
}

func NewAuthRoutes(loginHandler *handlers.LoginHandler) *AuthRoutes {
	return &AuthRoutes{loginHandler: loginHandler}
}

func (ar *AuthRoutes) GetRoutes(router *gin.Engine) {

	auth_routes := router.Group("api/v1/auth")

	auth_routes.POST("/login", ar.loginHandler.Login)

}
