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

	r.POST("/signup", controllers.Signup)

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080}
}
