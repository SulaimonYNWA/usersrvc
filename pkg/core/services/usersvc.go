package services

import (
	"Bank/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)
const LoginOperation = `Введите логин и пароль:`

const AuthorizationUser = `1. Показать Баланс
2.Перевод Денег
3.Оплата Услуг
4.История транзакций
5.Показать список банкоматов
0.Выход`

const AuthorizedOperation = `1. Показать Баланс
2.Перевод Денег
3.Оплата Услуг
4.История транзакций
5.Показать список банкоматов
6.Добавить Адресс банкомата
7.Показать пользователей
8.Добавить пользователя
0.Выход`


func Login(database *sql.DB)(ok bool, id int64, isAdmin bool) {
	var login, password string
	fmt.Println(LoginOperation)
	fmt.Println(`login: `)
	fmt.Scan(&login)
	fmt.Println(`password: `)
	fmt.Scan(&password)

	var User models.User
	row := database.QueryRow(`select * from users where login = ($1) and password = ($2) `, login, password)
	_ = row.Scan( //err :=
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
	fmt.Println(User)

	if User.ID > 0 {
		return true, User.ID, User.IsAdmin
		//if User.IsAdmin {
		//	fmt.Println(`Вы обладате правами и возможностями админа.`)
		//	Authorization(database, id)
		//}else if !User.IsAdmin {
		//	fmt.Println(`Вы обладате правами и возможностями пользователя`)
		//	UserAuthorization(database, id)
		//}
	} else {return false, User.ID,User.IsAdmin}

	return true, User.ID, User.IsAdmin
	return
}


func UserAuthorization (database *sql.DB, id int64)  {
	fmt.Println(AuthorizationUser)
	fmt.Println(`выберите команду: `)

	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(`Ваш баланс: `)
		CheckBalance(database, id)
	case 2:
		var sum, number2 int64
		fmt.Println(`введите сумму: `)
		fmt.Scanln(&sum)
		fmt.Println(`Введите номер карты, на которую хотите перевести сумму :`)
		fmt.Scanln(&number2)

		Transfer(database, id, sum, number2)


	case 3:
		fmt.Println(`Ops....`)

	case 4:
		fmt.Println("Input number of Account:")
		var number int64
		fmt.Scan(&number)
		fmt.Println(`Ваша история транзакций:`)
		Archive(database, number)
	case 5:
		fmt.Println(`список банкоматов: `)
		ATMs(database)

	case 0:

		os.Exit(0)
	}
}

func Authorization(database *sql.DB, id int64){
	fmt.Println(AuthorizedOperation)
	fmt.Println(`выберите команду: `)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(`Ваш баланс: `)
		CheckBalance(database, id)
	case 2:
		var sum, number2 int64
		fmt.Println(`введите сумму: `)
		fmt.Scan(&sum)

		fmt.Println(`Введите номер карты, на которую хотите перевести сумму :`)
		fmt.Scan(&number2)

		Transfer(database, id, sum, number2)

	case 3:
		//Логика кода такая же как и в предыдущем пункте. Функцию реализую после того как подробно узнаю об оплате услуг.
		fmt.Println(`Wait...`)

	case 4:
		fmt.Println("Input number of Account:")
		var number int64
		fmt.Scan(&number)
		fmt.Println(`Ваша история транзакций:`)
		Archive(database, number)

	case 5:
		fmt.Println(`список банкоматов: `)
		ATMs(database)

	case 6:

		AddNewATM(database)
	case 7:
		fmt.Println(`Список пользователей:`)
		Users(database)

	case 8:

		Registration(database)

	case 0:
		os.Exit(0)
		fmt.Println(`Exit`)

	default:
		fmt.Println("repeat again")
	}
	return
}

//1
func CheckBalance(database *sql.DB, id int64) {
	var amount int64
	row := database.QueryRow(`select amount from accaunts
where user_id = ($1)`, id)
	_ = row.Scan(
		&amount,
	)
	fmt.Println(amount)
}

//2
func Transfer(database *sql.DB, id, sum, number2 int64) {
	var amount,number int64
	row := database.QueryRow(`select amount, number from accaunts where user_id =($1)`, id)
	_ = row.Scan(
		&amount,
		&number,
	)
	newAmount:=amount-sum
	if amount < sum {
		log.Println(`недостаточно средств`)
		return
	}else {
	fmt.Println(`number:`,number,`amount: `,amount, `amount - sum :`, newAmount)
	_, err := database.Exec(`update accaunts set amount = ($1) where number = ($2)`, newAmount, number)
	if err != nil {
		log.Println(`cannot transfer money`, err)
		return
	}
	}

	row = database.QueryRow(`select amount from accaunts where number = ($2) `, number2)
	_ = row.Scan(
		&amount,
	)
	newReceiverAmount := amount+sum
	fmt.Println(`number:`,number2,`amount: `,amount, `amount+sum: `, newReceiverAmount)

	_, err := database.Exec(`update accaunts set amount = ($1) where number = ($2)`, newReceiverAmount, number2)
	if err != nil {
		log.Println(`cannot receive money`, err)
		return
	}
	models.AddTransaction(database, number, number2, sum, newAmount)
}

//3 Оплата услуг


//4 история транзакций
func Archive(database *sql.DB, sender int64)  {
	transaction := models.Transaction{}

	rows, err := database.Query(`select id, date,time,amount, sender_number, receiver_number, available_limit from archive  where sender_number = ($1)`, sender)
	if err != nil {
		log.Fatal(err, `users are not selected`)
	}
	defer rows.Close()
	fmt.Println(`ID 	Date         Time      Amount    Your number      Receiver      Rest `)

	for rows.Next(){
		err:= rows.Scan(
			&transaction.ID,
			&transaction.Date,
			&transaction.Time,
			&transaction.Amount,
			&transaction.SenderNumber,
			&transaction.ReceiverNumber,
			&transaction.AvailableLimit,
		)
		if err != nil {
			log.Fatal(err, ` not selected archive`)
		}
	fmt.Println(transaction.ID, `  `, transaction.Date, `  `, transaction.Time, `     `, transaction.Amount, `         `, transaction.SenderNumber, `         `, transaction.ReceiverNumber, `     `, transaction.AvailableLimit)
	}
}

//5 список банкоматов
func ATMs(database *sql.DB) {
	var id int64
	var address string
	rows, err := database.Query(`select id, address from ATMs `)
	if err != nil {
		log.Fatal(err, `not selected`)
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(
			&id,
			&address,
		)
		if err != nil {
			log.Fatal(err, `not selected next`)
		}
	fmt.Println(id, address)
	}
		log.Println(id, address)

}


//6 добавть банкомат
func AddNewATM(database *sql.DB)(ok bool) {

		fmt.Println(
			`1.add ATM
2.no, exit'`)
		var number int
		fmt.Scan(&number)
		switch number {
		case 1:
			fmt.Println(`enter a new address: `)
			var s string
			fmt.Scan(&s)
			reader := bufio.NewReader(os.Stdin)
			Address, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf(`cant read command: %v`,err)
			}
			fmt.Println(s)
			sprintf := fmt.Sprintf(`%s %s`, s, Address)
			fmt.Println(sprintf)
			_, err = models.AddATM(database, sprintf)
			if err != nil {
				fmt.Println(`vse ploxo`, err)
			} else {
				fmt.Println(`vse ok`)
		}
		case 2:
			fmt.Println("Goodbye")
			os.Exit(0)
		}
	return
}

//7 список пользователей
func Users(database *sql.DB)  {
	var id, age int64
	var name, surname string

	rows, err := database.Query(`select id, age, name, surname from users `)
	if err != nil {
		log.Fatal(err, `users are not selected`)
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(
			&id,
			&age,
			&name,
			&surname,
		)
		if err != nil {
			log.Fatal(err, ` not selected next`)
		}
		fmt.Println(id, name, surname, age)
	}
	return
}




//8 добавить пользователя
func Registration(database *sql.DB ) (err error) {
	//var User models.User
	var name,surname,gender, login, password string
	var age int64
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
	_, s := database.Exec(`insert into users(name, surname,age, gender, login, password) values (($1),($2),($3),($4),($5),($6))`, name, surname, age, gender, login,password)
	if s!=nil {
		fmt.Println(`cannot registre the new user`, s)
		return s
	}
	fmt.Println(`new user is registered!`)
	return err
}