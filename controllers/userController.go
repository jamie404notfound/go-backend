package controllers

import (
	"jamie404notfound/go-backend/initializers"
	"jamie404notfound/go-backend/models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var body struct {
		Username string
		Password string
		Secret   string
	}

	c.Bind(&body)

	user := models.User{Username: body.Username, Password: body.Password, Secret: body.Secret}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"user": users,
	})
}
