package main

import (
	"Bank/db"
	"Bank/models"
	"Bank/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	database, err := sql.Open("sqlite3", "test")
	if err != nil {
		log.Fatal("error1", err)
	} else {
		fmt.Println("CONNECTION TO DB IS SUCCESS")
	}
	db.DBInit(database)
	Start(database)

}

const AuthorizationOperation = `1.Authorization
0.Exit`

func Start(database *sql.DB) {
	intro(database)

}
func intro(database *sql.DB) {

	for {
		fmt.Println(AuthorizationOperation)
		fmt.Println(`select a command:`)
		var cmd int64
		_, err := fmt.Scan(&cmd)
		if err != nil {
			log.Println(`error on line 35 main.go`)
		}
		var User models.User
		switch cmd {

		case 1:
			ok, id, isAdmin := services.Login(database)
			row := database.QueryRow(`select id, isAdmin from users where id = ($1) and isAdmin = ($2) `, id, isAdmin)
			_ = row.Scan(
				&User.ID,
				&User.IsAdmin,
			)
			if ok {
				if User.IsAdmin {
					fmt.Println(`You are admin.`)
					services.Authorization(database, id)
				}
				if !User.IsAdmin {
					fmt.Println(`You are user`)
					services.UserAuthorization(database, id)
				}
				fmt.Println(ok)
			} else {
				fmt.Println(`not ok`)
			}

		case 0:
			os.Exit(0)
		default:
			fmt.Println(`repeat again`)
		}

	}
}
