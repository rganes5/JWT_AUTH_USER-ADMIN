package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rganes5/go-jwt-auth/initializers"
	"github.com/rganes5/go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// Handler for creating a new user
func AdminSubmitHandler(c *gin.Context) {
	userName := c.Request.FormValue("Name")
	userEmail := c.Request.FormValue("Email")
	userPassword := []byte(c.Request.FormValue("Password"))
	userNumber := c.Request.FormValue("Number")
	userRole := c.Request.FormValue("Role")
	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(userPassword, bcrypt.DefaultCost)
	if err != nil {
		panic("Failed To Hash Password")
	}
	//Inputting into the database
	user := models.User{Name: userName, Email: userEmail, Password: string(hashedPassword), Number: userNumber, Role: userRole}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "createUser.html", gin.H{
			"error": "Email / Contact Number already exist!",
		})
		return
	}
	c.Redirect(http.StatusFound, "/AdminPanel")
}

// Handler to load the welcome page of admin
func AdminPanelHandler(c *gin.Context) {
	var user []models.User
	if err := initializers.DB.Where("role!=?", "admin").Find(&user).Error; err != nil {
		fmt.Println("Error loading the database")
	}
	c.HTML(http.StatusOK, "welcomeAdmin.html", gin.H{
		"data": user,
	})
	for _, v := range user {
		fmt.Println(v.ID, v.Name, v.Email, v.Password, v.Number, v.Role)
	}
}

// Handler to get to load new user page
func AdminCreateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "createUser.html", nil)
}

// Handler to edit the user details
func EditAdminHandler(c *gin.Context) {
	var user []models.User
	id := c.Query("id")
	if err := initializers.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		fmt.Println("Failed to Edit User")
	}
	c.HTML(http.StatusOK, "editAdmin.html", gin.H{
		"data": user,
	})
}

// Handler to delete the user
func DeleteHandler(c *gin.Context) {
	var user []models.User
	id := c.Query("id")
	result := initializers.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		panic("failed to delete user")
	} else {
		c.Redirect(http.StatusFound, "/AdminPanel")
	}
}

// Handler to Update the user data
func UpdateHandler(c *gin.Context) {
	id := c.Query("id")
	var user models.User
	user.Name = c.Request.FormValue("Username")
	user.Email = c.Request.FormValue("Usermail")
	editedpass := []byte(c.Request.FormValue("Password"))
	// Hash the password with bcrypt
	hashed, err := bcrypt.GenerateFromPassword(editedpass, bcrypt.DefaultCost)
	if err != nil {
		panic("Failed To Hash Password")
	}
	user.Password = string(hashed)
	// user.Password = c.Request.FormValue("Password")
	user.Number = c.Request.FormValue("Number")
	user.Role = c.Request.FormValue("Role")
	result := initializers.DB.Where("id=?", id).Updates(&user)
	if result.Error != nil {
		panic("failed to update user")
	} else {
		c.Redirect(http.StatusFound, "/AdminPanel")
	}
}

// Handler to search the user
func SearchHandler(c *gin.Context) {
	var user []models.User
	name := c.Query("Name")
	fmt.Println(name)
	result := initializers.DB.Where("Name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Println("User not found")
		c.Redirect(http.StatusFound, "/AdminPanel")
		// c.HTML(http.StatusNotFound, "welcomeAdmin.html", gin.H{
		// 	"errormessage": "User not found",
		// })
	} else {
		c.HTML(http.StatusOK, "welcomeAdmin.html", gin.H{
			"data": user,
		})
	}
}

// var user []models.User
// id := c.Query("id")
// if err := initializers.DB.Where("id = ?", id).Find(&user).Error; err != nil {
// 	fmt.Println("Failed to Edit User")
// }
// c.HTML(http.StatusOK, "editAdmin.html", gin.H{
// 	"data": user,
