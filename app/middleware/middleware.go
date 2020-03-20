package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

func Panic(c *gin.Context){
	if err := recover(); err != nil {
		data := fmt.Sprintf("Recovering from panic in printAllOperations error is: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": data,
		})
		c.Abort()
   }
}

func PanicHandler(c *gin.Context) {
	//panic handler
	defer Panic(c)
	c.Next()
}

func AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		// Pass on to the next-in-chain
		c.Next()
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}