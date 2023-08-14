package main

import (
	"github.com/gin-gonic/gin"

	"users_service_api/config"
	"users_service_api/entities"
	"users_service_api/handlers"
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

	web_app = gin.Default()

	//config init
	config_database = config.NewDatabase()
	config_database.LoadEnvFile()
	config_database.ConnectToDB()
	utils := utils.NewFunctions()

	// Run migrations
	config_database.DB.AutoMigrate(&entities.User{}, &entities.Role{})

	// User routes init
	user_repository = repositories.NewUserRepository(config_database, utils)
	user_handler = handlers.NewUserHandler(user_repository, utils)
	user_router = routes.NewUserRoutes(user_handler)
	user_router.GetRoutes(web_app)

	// Role routes init
	role_repository = repositories.NewRoleRepository(config_database)
	role_handler = handlers.NewRoleHandler(role_repository)
	role_router = routes.NewRoleRoutes(role_handler)
	role_router.GetRoutes(web_app)

	web_app.Run()

}
