package main

import (
	"jamie404notfound/go-backend-crud/initializers"
	"jamie404notfound/go-backend-crud/initializers/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
