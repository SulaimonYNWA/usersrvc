package services

import (
	"Bank/models"
	"database/sql"
	"fmt"
)

const AuthorizationOperation = `1.Авторизация, 
2.Выйти`

const LoginOperation = `Введите логин и пароль: `


func Authorization()(login, password string){
	fmt.Println(AuthorizationOperation)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
	fmt.Println(LoginOperation)
		fmt.Println(`login: `)
	fmt.Scan(&login)
		fmt.Println(`password: `)
		fmt.Scan(&password)
	return password, login
	case 2:
		fmt.Println("Goodbye")
	default:
		fmt.Println("repeat again")
	}
	return
}

func Login(database *sql.DB, login, password string){ //(ok bool)
	var User models.User
	_ = database.QueryRow(`select * from users where login = ($1) and password = ($2) `, login, password).Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Remove,
		)
	//if User.ID >0{
	//	return true
	//}
	//return false
	fmt.Println(User)
}