package main

import (
	"os"

	"github.com/danhngocphh/fptda2/api"
	"github.com/danhngocphh/fptda2/database"
	"github.com/danhngocphh/fptda2/lib/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
