package main

import (
	app "campfire/internal"
	"campfire/internal/database"
	"campfire/pkg/config"
)

func main() {
	config.LoadEnv(".env")

	database.CreateMongodbConnection()
	database.CreatePostgresConnection()

	defer database.CloseMongodbConnection()
	defer database.ClosePostgresConnection()

	r := app.NewApp()
	r.Start()
}
