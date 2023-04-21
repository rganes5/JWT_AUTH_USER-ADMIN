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

	defer func() {

		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{})
			fmt.Printf("\n\nRecovered from panic. %s\n\n", err)
		}
	}()

	//Get the cookie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode/validate it
	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && token != nil {

		//Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		//Find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["ID"])

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attach to req
		c.Set("user", user)

		//Continue

		c.Next()
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		// c.AbortWithStatus(http.StatusUnauthorized)
	}

}
