package service

import "github.com/rudirahardian/go_env/app/repository"
import "github.com/gin-gonic/gin"
import "github.com/rudirahardian/go_env/app/models"
import "time"
import "os"
import "strconv"
import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func FindUser(username string, password string) ([]repository.User, error){
	//add repository insert User
	return repository.LoginQuery(username, password)
}

func InsertUser(c *gin.Context) (models.User, error){
	var user models.User
	
	file, err := c.FormFile("foto")
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	fileName := strconv.FormatInt(expirationTime.Unix(),10) + file.Filename
	path := "images/" + fileName
	if file.Header.Get("Content-Type") == "image/jpeg" || file.Header.Get("Content-Type") == "image/png"{
		if err := c.SaveUploadedFile(file, path); err != nil {
			return nil, err
		}
	}else{
		return nil, err
	}

	user = &models.Users{
		Name: c.PostForm("name"),
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Foto: fileName,
	}


	if _, err := user.InsertData(); err != nil{
		err := os.Remove(path)

		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func ExtractClaims(tokenStr string) (*Claims, error) {
	tknStr := tokenStr
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