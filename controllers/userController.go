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

func LoginCheck(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON body to the 'input' struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Initialize a user model and query the database
	var user models.User
	result := initializers.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user)

	// Check if the user was found
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Optionally, validate the password here
	if input.Password != user.Password {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// Respond with the user details (not including password or secret if not needed)
	c.JSON(200, gin.H{
		"user": user.Username,
	})
}
