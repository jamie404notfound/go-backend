package controllers

import (
	"jamie404notfound/go-backend/initializers"
	"jamie404notfound/go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
	var body models.Body

	// Bind the incoming JSON body to 'body'
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user with the hashed password
	user := models.User{
		Username: body.Username,
		Password: string(hashedPassword), // Store the hashed password
		Secret:   body.Secret,
	}

	// Save the user to the database
	result := initializers.DB.Create(&user)

	// Check for database errors
	if result.Error != nil {
		c.Status(400) // Bad request
		return
	}

	// Respond with the created user (excluding password for security reasons)
	c.JSON(200, gin.H{
		"user": gin.H{
			"username": user.Username,
			"secret":   user.Secret, // You can return other user details if needed
		},
	})
}

func GetUsers(c *gin.Context) {
	var users models.Users
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"user": users.Username,
	})
}

func LoginCheck(c *gin.Context) {
	// Bind and validate JSON input
	var input models.LoginInput
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
