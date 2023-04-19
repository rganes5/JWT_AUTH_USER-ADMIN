package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rganes5/go-jwt-auth/initializers"
	"github.com/rganes5/go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// This will recieve the Gin context
func Signup(c *gin.Context) {
	// Get the email and password off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash password",
		})
		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{})
}
