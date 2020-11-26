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

func Start(database *sql.DB)  {
	for {
		login, password := services.Authorization()
		services.Login(database,login,password)
	}
}
