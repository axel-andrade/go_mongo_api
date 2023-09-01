package main

import (
	"go_mongo_api/src/infra/bootstrap"
	"go_mongo_api/src/infra/database"
	"go_mongo_api/src/infra/http/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
*
A função init por padrão é a primeira a ser executada pelo go.
Utilizada para configurar ou fazer um pré carregamento.
*
*/
func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	database.ConnectDB()
	database.ConnectRedisDB()

	// dependecies := bootstrap.LoadDependenciesV2()

	// logoutCtrl := dependecies.Get("logoutCtrl").(*controllers.LogoutController)

	// print(logoutCtrl)
}

func main() {
	dependecies := bootstrap.LoadDependencies()

	server := server.NewServer(os.Getenv("PORT"))
	server.AddRoutes(dependecies)
	server.Run()
}
