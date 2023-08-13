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

	web_app = gin.Default()

	//config init
	config := config.NewDatabase()
	config.LoadEnvFile()
	config.ConnectToDB()

	config.DB.AutoMigrate(&entities.User{})

	// User routes init
	utils := utils.NewFunctions()
	user_repository := repositories.NewUserRepository(config, utils)
	user_handler := handlers.NewUserHandler(user_repository, utils)
	user_router := routes.NewUserRoutes(user_handler)
	user_router.GetRoutes(web_app)

	web_app.Run()

}
