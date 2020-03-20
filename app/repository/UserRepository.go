package repository

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import helper "github.com/rudirahardian/go_env/config"

type User struct {
	Id int
	Username string
	Password string
	Name string
	Foto string
	Count int
}

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql", helper.DotEnvVariable("user")+":"+helper.DotEnvVariable("password")+"@tcp("+helper.DotEnvVariable("DB_HOST")+":3306)/"+helper.DotEnvVariable("database"))
    if err != nil {
        return nil, err
    }

    return db, nil
}

func LoginQuery(username string, password string) ([]User,error) {
    db, err := connect()
    if err != nil {
        return nil, err
    }

    rows, err := db.Query("select id, name, username, password, count(*) as count from user where username = ? AND password = ? group by id ASC ", username, password)

	if err != nil {
        return nil, err
	}
	
    defer rows.Close()

    var result []User

    for rows.Next() {
        var each = User{}
        var err = rows.Scan(&each.Id, &each.Name, &each.Username, &each.Password, &each.Count)

        if err != nil {
            return result, err
        }

        result = append(result, each)
	}

	if err = rows.Err(); err != nil {
        return result, err
    }
	
	db.Close()
    return result, nil
}

func InsertQuery(user User){
	db, err := connect()
    if err != nil {
        fmt.Println(err)
    }

    rows, err := db.Query("insert into user (name, username, password, foto) values (?, ?, ?, ?)", user.Name, user.Username, user.Password, user.Foto)

	if err != nil {
        fmt.Println(err)
	}
	
    defer rows.Close()
}