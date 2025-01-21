package controllers

import (
	"jamie404notfound/go-backend/initializers"
	"jamie404notfound/go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Body struct {
	Username string
	Password string
	Secret   string
}

func UserCreate(c *gin.Context) {
	var body Body

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
	// Bind and validate JSON input
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Query the database for the user
	var user models.User
	err := initializers.DB.Where("username = ?", input.Username).First(&user).Error

	// Handle errors and user not found
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Compare hashed passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Respond with sanitized user details
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"username": user.Username,
		},
	})
}
