package main

import (
	"github.com/gin-gonic/gin"

	"users_service_api/config"
	"users_service_api/entities"
	"users_service_api/handlers"
	middleware "users_service_api/middlewares"
	"users_service_api/repositories"
	"users_service_api/routes"
	"users_service_api/utils"
)

func init() {
}

func main() {
	var web_app *gin.Engine
	var config_database *config.Database
	var user_repository *repositories.UserRepository
	var user_handler *handlers.UserHandler
	var user_router *routes.UserRoutes
	var role_repository *repositories.RoleRepository
	var role_handler *handlers.RoleHandler
	var role_router *routes.RoleRoutes
	var user_role_repository *repositories.UserRoleRepository
	var user_role_handler *handlers.UserRoleHandler
	var user_role_routes *routes.UserRoleRoutes
	var auth_middleware *middleware.AuthMiddleware
	var login_handler *handlers.LoginHandler
	var login_routes *routes.AuthRoutes

	web_app = gin.Default()

	//config init
	config_database = config.NewDatabase()
	config_database.LoadEnvFile()
	config_database.ConnectToDB()
	utils := utils.NewFunctions()

	// Run migrations
	config_database.DB.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.UserRoles{})

	//Middleware init

	// User routes init
	user_repository = repositories.NewUserRepository(config_database, utils)
	user_handler = handlers.NewUserHandler(user_repository, utils)
	user_router = routes.NewUserRoutes(user_handler, auth_middleware)
	user_router.GetRoutes(web_app)

	// Role routes init
	role_repository = repositories.NewRoleRepository(config_database)
	role_handler = handlers.NewRoleHandler(role_repository)
	role_router = routes.NewRoleRoutes(role_handler, auth_middleware)
	role_router.GetRoutes(web_app)

	// User Role routes init
	user_role_repository = repositories.NewUserRoleRepository(config_database)
	user_role_handler = handlers.NewUserRoleHandler(user_role_repository)
	user_role_routes = routes.NewUserRoleRoutes(user_role_handler, auth_middleware)
	user_role_routes.GetRoutes(web_app)

	// Login Routes init
	login_handler = handlers.NewLoginHandler(user_repository, utils)
	login_routes = routes.NewAuthRoutes(login_handler)
	login_routes.GetRoutes(web_app)

	web_app.Run()

}
