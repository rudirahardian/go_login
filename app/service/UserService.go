package service

import "github.com/rudirahardian/go_env/app/repository"
import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func FindUser(username string, password string) ([]repository.User, error){
	//add repository insert User
	return repository.LoginQuery(username, password)
}

func InsertUser(user repository.User){
	repository.InsertQuery(user)
}

func ExtractClaims(tokenStr string) (*Claims, error) {
	// Get the JWT string from the cookie
	tknStr := tokenStr

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims,err
		}
		return claims,err
	}
	if !tkn.Valid {
		return claims,err
	}

	return claims,err
}