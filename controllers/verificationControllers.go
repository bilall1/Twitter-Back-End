package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	// PostgreSQL driver
)

func AuthenticateJWT(c *gin.Context) {

	tp := c.GetHeader("ThirdParty")

	b, _ := strconv.ParseBool(tp)

	// Get the JWT token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	c.Next()

	if b == false {

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the same secret key used for signing the token
			return []byte("aurorasolutions"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()

	} else {

		const googleClientId = "753336473496-0m3lj51nmnmfo995oodetpfgrah45icu.apps.googleusercontent.com"

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return

		}

		payload, err := tokenValidator.Validate(context.Background(), tokenString, googleClientId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return

		}
		email := payload.Claims["email"]
		_ = email
		c.Next()

	}

}
