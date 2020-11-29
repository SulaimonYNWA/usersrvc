package services

import (
	"Bank/models"
	"database/sql"
	"fmt"
	"os"
)

const AuthorizationOperation = `1.Авторизация 
2.Выйти
3.Регистрация:`

const LoginOperation = `Введите логин и пароль: `


func Authorization()(login, password, name, surname, gender string, age int64){
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
	return
	case 2:
		fmt.Println("Goodbye")
		os.Exit(0)
	case 3:
		fmt.Println(`Ваше имя и фамилия:`)
		fmt.Scan(&name)
		fmt.Scan(&surname)
		fmt.Println(`Ваш возраст:`)
		fmt.Scan(&age)
		fmt.Println(`Пол: `)
		fmt.Scan(&gender)
		fmt.Println(LoginOperation)
		fmt.Println(`login: `)
		fmt.Scan(&login)
		fmt.Println(`password: `)
		fmt.Scan(&password)

	default:
		fmt.Println("repeat again")
	}
	return
}

func Login(database *sql.DB, login, password string){ //(ok bool)
	var User models.User
	row := database.QueryRow(`select * from users where login = ($1) and password = ($2) `, login, password)
	err := row.Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Remove,
		)
	if err != nil{
		fmt.Println(err, `mistake`)
	}
	//if User.ID >0{
	//	return true
	//}
	//return false
	fmt.Println(User)
}

func Registration(database *sql.DB, name,surname,gender, login, password string, age int64) (err error) {
	//var User models.User
	_, s := database.Exec(`insert into users(name, surname,age, gender, login, password) values (($1),($2),($3),($4),($5),($6))`, name, surname, age, gender, login,password)
	if s!=nil {
		fmt.Println(`cannot registre the new user`, s)
		return s
	}
	fmt.Println(`new user is registered!`)
	return err
}