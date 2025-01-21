package main

import (
	"jamie404notfound/go-backend/controllers"
	"jamie404notfound/go-backend/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Configure CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Allow requests from your Vue frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allow specific HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow specific headers
		AllowCredentials: true,                                                // Allow credentials (cookies, authorization headers, etc.)
	}))

	r.GET("/getUsers", controllers.GetUsers)
	r.POST("/login", controllers.LoginCheck)
	r.POST("/createUser", controllers.UserCreate)

	r.Run() // listen and serve on 0.0.0.0:3000
}
