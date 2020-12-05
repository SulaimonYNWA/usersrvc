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
		if err!=nil{
			log.Fatal("error1", err)
		}else {
			fmt.Println("CONNECTION TO DB IS SUCCESS")
		}
	db.DBInit(database)
	Start(database)

}
const AuthorizationOperation = `1.Авторизация
2.Выйти`
var User models.User

func Start(database *sql.DB){
	for {
		fmt.Println(AuthorizationOperation)
		fmt.Println(`Выберите команду:`)
		var cmd int64
		_, err := fmt.Scan(&cmd)
		if err != nil {
			log.Println(`error on line 35 main.go`)
		}
		switch cmd {
		case 1:
			ok, id,isAdmin := services.Login(database)
			row := database.QueryRow(`select id, isAdmin from users where id = ($1) and isAdmin = ($2) `, id, isAdmin)
			_ = row.Scan(
				&User.ID,
				&User.IsAdmin,
				)
			if ok {
				if User.IsAdmin {
					fmt.Println(`Вы обладате правами и возможностями админа.`)
					services.Authorization(database, id)
				}
				if !User.IsAdmin{
					fmt.Println(`Вы обладате правами и возможностями пользователя`)
					services.UserAuthorization(database, id)
				}
				//services.Authorization(database, id)
			fmt.Println(ok)
			}else {fmt.Println(`damn..`)}

		case 2:
			os.Exit(0)
		default:
			fmt.Println(`repeat again`)
		}
		//login, password,name, surname, gender,age:= services.Authorization(database)
		//services.Login(database,login,password)
		//return  name,surname,gender,age
	}
}
