package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rganes5/go-jwt-auth/controllers"
	"github.com/rganes5/go-jwt-auth/initializers"
)

// Init will work first
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// This is dockerized now
// Main func
func main() {
	// Returns a new instance of the Gin engine with the default middleware already attached.
	r := gin.Default()

	// For accessing all the html files.
	r.LoadHTMLGlob("templates/*.html")
	//For accessing all the assets like css and images.
	r.Static("/assets", "./assets")

	//Welcome page
	r.GET("/", controllers.IndexHandler)
	// r.GET("/", middleware.RequireAuth, controllers.IndexHandler)

	//USER

	//Common Login page Route
	r.POST("/welcome", controllers.LoginHandler)

	//Signup page Route
	r.GET("/userSignup", controllers.UserSignupHandler)

	//User Submit on Signup page
	r.POST("/user-Submit", controllers.UserSubmitHandler)

	//Logout Route
	r.GET("/logout", controllers.LogoutHandler)

	//ADMIN

	//Admin Welcome
	r.GET("/AdminPanel", controllers.AdminPanelHandler)

	//Admin Create new user
	r.GET("/createUser", controllers.AdminCreateHandler)

	//Admin submit action on new user
	r.POST("/admin-Submit", controllers.AdminSubmitHandler)

	//Admin edit
	r.GET("/AdminPanel/editAdmin", controllers.EditAdminHandler)

	//Admin delete
	r.GET("/AdminPanel/delete", controllers.DeleteHandler)

	//Admin update
	r.POST("/update", controllers.UpdateHandler)

	//Admin Search
	r.GET("/searchUser", controllers.SearchHandler)

	//Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	r.Run() // listen and serve on 0.0.0.0:8080}
}
