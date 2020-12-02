package main

import (
	"Bank/db"
	"Bank/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "test")
		if err!=nil{
			fmt.Println("error1", err)
		}
	db.DBInit(database)
	Start(database)

}

func Start(database *sql.DB)(name, surname, gender string, age int64) {
	for {
		login, password,name, surname, gender,age:= services.Authorization(database)
		services.Login(database,login,password)
		return  name,surname,gender,age
	}
}
