package main

import (
	"jamie404notfound/go-backend/controllers"
	"jamie404notfound/go-backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/getUsers", controllers.GetUsers)
	r.GET("/login", controllers.LoginCheck)
	r.POST("/createUser", controllers.UserCreate)

	r.Run() // listen and serve on 0.0.0.0:3000
}
