package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rganes5/go-jwt-auth/initializers"
	"github.com/rganes5/go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
)

////////////HANDLERS/////////////

// Handler for user login/index page
func IndexHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	c.HTML(http.StatusOK, "index.html", nil)
}

// Handler for loading user sign up page

func UserSignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "userSignup.html", nil)
}

// Handler for User sign up
func UserSubmitHandler(c *gin.Context) {
	userName := c.Request.FormValue("Name")
	userEmail := c.Request.FormValue("Email")
	userPassword := []byte(c.Request.FormValue("Password"))
	userNumber := c.Request.FormValue("Number")
	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(userPassword, bcrypt.DefaultCost)
	if err != nil {
		panic("Failed To Hash Password")
	}
	//Inputting into the database
	user := models.User{Name: userName, Email: userEmail, Password: string(hashedPassword), Number: userNumber}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "userSignup.html", gin.H{
			"error": "Email / Contact Number already exist!",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")

}

// Handler for user Login page
func LoginHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	userEmail := c.Request.FormValue("Email")
	userPassword := []byte(c.Request.FormValue("Password"))
	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", userEmail)

	if user.ID == 0 {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Invalid Email",
		})
		return
	}
	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), userPassword)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Invalid Password",
		})
		return
	}
	//Generate a jwt token for user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.ID,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println("Error")
	}
	//Send in back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	//Redirecting to welcome page
	fmt.Println("--", user.Role, user.Name, "SECRET", tokenString)
	//Check whether it is a user or admin
	if user.Role == "user" {
		c.HTML(http.StatusOK, "welcomeUser.html", user.Name)
		// c.Redirect(http.StatusFound, "/UserPanel")
	} else {
		c.Redirect(http.StatusFound, "/AdminPanel")
		// c.HTML(http.StatusOK, "welcomeAdmin.html", user.Name)
	}
}

// Handler for common logout
func LogoutHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, -1, "", "", false, true)
	fmt.Println("Token cleared")
	c.Redirect(http.StatusFound, "/")

}

// Authorization middlewares
func Validate(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "I'm logged in",
	user, _ := c.Get("user")
	//if i wanted to do something with user
	// user.(models.User).Email
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
