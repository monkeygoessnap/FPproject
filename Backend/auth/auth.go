/*
Package auth provides the authentication middleware for the mux router
*/
package auth

import (
	"FPproject/Backend/log"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	Name string
	ID   string
	jwt.StandardClaims
}

func GenerateJWT(name, id string) (map[string]string, error) {

	key := os.Getenv("JWT_SECRETKEY") //JWT SECRET KEY

	claims := &jwtCustomClaims{
		name,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "FPproject",
			IssuedAt:  time.Now().Unix(),
		},
	}
	//access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}

	return map[string]string{
		"access_token": tokenString,
		"expire":       strconv.Itoa(int(claims.StandardClaims.ExpiresAt)),
	}, nil
}

//JWT Auth Middleware
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("JWT_SECRETKEY")
		tokenString := c.GetHeader("access_token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New("unexpected signing method")
				log.Info.Println(err)
				return nil, err
			}
			return []byte(key), nil
		})
		if err != nil {
			log.Info.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			c.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)["ID"]
			c.Set("ID", claims)
			c.Next()
		}
	}
}
