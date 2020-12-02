package services

import (
	"Bank/models"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

const AuthorizationOperation = `1.Авторизация 
2.Выйти
3.Регистрация:`

const LoginOperation = `Введите логин и пароль: `


func Authorization(database *sql.DB)(login, password, name, surname, gender string, age int64){
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
		Registration(database, name,surname,gender,login,password, age)
		return

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
		&User.IsAdmin,
		&User.Remove,
		)
	if err != nil{
		fmt.Println(err, `mistake`)
	}
	fmt.Println(User)

	if User.IsAdmin {
		fmt.Println(
			`1.add ATM
2.no, exit'`)
		var number int
		fmt.Scan(&number)
		switch number {
		case 1:
			address := bufio.NewReader(os.Stdin)
			fmt.Println(`enter address: `)
			text, _:= address.ReadString('\n')
			fmt.Scan(&text)
			//fmt.Print(text, ` ` )
			text2 := ""
			fmt.Scanln(&text2)
			//fmt.Print(text2, ` `)
			var ln string
			fmt.Scanln(&ln)

			var Address string
			Address = text+` `+text2+ ` `+ln
			//fmt.Println(text, ``, text2,``, ln)
			fmt.Println(Address)
			//models.AddATM(database, text, text2, ln)
			models.AddATM(database, Address)
		case 2:
			fmt.Println("Goodbye")
			os.Exit(0)
		}
	}
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