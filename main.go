package main

import (
	"backend/config"
	"backend/database"
	"backend/routes"
)

func main() {

	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3002"))
}