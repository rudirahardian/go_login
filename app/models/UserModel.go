package models

import userRepository "github.com/rudirahardian/go_env/app/repository"
import "fmt"

type User interface {
	InsertData() (string, error)
}

type Users struct {
	Id int
	Username string
	Password string
	Name string
	Foto string
	Count int
}

func (user *Users) InsertData() (string, error){
	db, err := userRepository.Connect()
	if err != nil {
		fmt.Println(err)
	}
	
	_, err = db.Query("insert into user (name, username, password, foto) values (?, ?, ?, ?)", user.Name, user.Username, user.Password, user.Foto)	
	if err != nil {
		fmt.Printf("DB Error: ")
		fmt.Println(err)
		return "error", err
	}
		
	db.Close()
	
	return "success", err
}