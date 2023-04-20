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

func main() {
	// Returns a new instance of the Gin engine with the default middleware already attached.
	r := gin.Default()

	// For accessing all the html files.
	r.LoadHTMLGlob("templates/*.html")
	//For accessing all the assets like css and images.
	r.Static("/assets", "./assets")

	//Welcome page
	r.GET("/", controllers.IndexHandler)

	//User Welcome page
	r.POST("/welcomeUser", controllers.UserLoginHandler)

	//Logout handler
	r.GET("/logout", controllers.LogoutHandler)

	//Signup page
	r.GET("/userSignup", controllers.UserSignupHandler)

	//User Submit on Signup page
	r.POST("/user-Submit", controllers.UserSubmitHandler)

	// r.POST("/signup", controllers.Signup)
	// r.POST("/login", controllers.Login)
	// r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080}
}
