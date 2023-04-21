package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rganes5/go-jwt-auth/initializers"
	"github.com/rganes5/go-jwt-auth/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Checking Middleware")

	//RECOVER PANIC
	defer func() {

		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{})
			fmt.Printf("\n\nRecovered from panic. %s\n\n", err)
		}
	}()

	//Get the cookie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie")
		// c.AbortWithStatus(http.StatusUnauthorized)
		c.HTML(http.StatusOK, "index.html", nil)
		return
	}
	fmt.Println("Cookie present")
	fmt.Println(tokenString)
	//Decode/validate it
	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("LINE1")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("LINE2")

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("LINE3")

		return []byte(os.Getenv("SECRET")), nil

	})
	fmt.Println("LINE4")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("LINE5")

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("LINE6")

		//Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Token tampered")
		}
		fmt.Println("LINE8")

		//Find the user with token sub
		var user models.User
		fmt.Println("LINE9")

		id := claims["ID"]
		fmt.Println("LINE10")

		fmt.Println(id)
		result := initializers.DB.Where("id = ?", id).Find(&user)
		if result.Error != nil {
			fmt.Println("Failed to find user")
		}
		if user.Role == "user" {
			c.HTML(http.StatusOK, "welcomeUser.html", user.Name)
		} else {
			c.HTML(http.StatusOK, "welcomeAdmin.html", nil)
		}

		//Attach to req
		// c.Set("user", user)

		//Continue
		c.Next()
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		// c.AbortWithStatus(http.StatusUnauthorized)
	}

}
