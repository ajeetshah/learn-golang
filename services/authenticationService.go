package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func BasicAuth(c *gin.Context) {
	gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	})
}

func GetToken(c *gin.Context) {
	token, _ := GenerateJWT()
	tokens = append(tokens, token)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: "username",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	reqTokenArray := strings.Split(bearerToken, " ")
	if len(reqTokenArray) < 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	reqToken := reqTokenArray[1]
	if reqToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
}
