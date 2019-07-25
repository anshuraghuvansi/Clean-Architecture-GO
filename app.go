package main

import (
	"os"
	"user/routes"
	"user/utils/database"

	"github.com/anshuraghuvansi/env"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("environment", "development")
	env.Configure()

	db, err := database.CreateAndConnect()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	app := gin.New()
	routes.Init(app, db)
	app.Run(os.Getenv("PORT"))
}
