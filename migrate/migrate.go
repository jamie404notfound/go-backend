package main

import (
	"jamie404notfound/go-backend/initializers"
	"jamie404notfound/go-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
