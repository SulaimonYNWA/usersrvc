package services

import (
	"Bank/db"
	"Bank/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

const LoginOperation = `Enter login and password:`

const AuthorizationUser = `1. Show balance
2.Transfer cash
3.Archive
4.List of ATMs
0.Exit`

const AuthorizedOperation = `1. Show balance
2.Transfer cash
3.Archive
4.List of ATMs
5.Add address of ATM
6.Show users
7.Add user
0.Exit`

func Login(database *sql.DB) (ok bool, id int64, isAdmin bool) {
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
	} else {
		return false, User.ID, User.IsAdmin
	}

	return true, User.ID, User.IsAdmin
	return
}

func UserAuthorization(database *sql.DB, id int64) {
	fmt.Println(AuthorizationUser)
	fmt.Println(`Select command: `)

	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(`Your balance: `)
		CheckBalance(database, id)
	case 2:
		var sum, number2 int64
		fmt.Println(`Input amount: `)
		fmt.Scanln(&sum)
		fmt.Println(`Input number of Receiver :`)
		fmt.Scanln(&number2)

		Transfer(database, id, sum, number2)

	case 3:
		fmt.Println("Input number of Account:")
		var number int64
		fmt.Scan(&number)
		fmt.Println(`Your transactions:`)
		Archive(database, number)
	case 4:
		fmt.Println(`list of ATMs: `)
		ATMs(database)

	case 0:

		os.Exit(0)
	}
}

func Authorization(database *sql.DB, id int64) {
	fmt.Println(AuthorizedOperation)
	fmt.Println(`select command: `)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println(`your balance: `)
		CheckBalance(database, id)
	case 2:
		var sum, number2 int64
		fmt.Println(`input amount: `)
		fmt.Scan(&sum)

		fmt.Println(` Input number of Receiver:`)
		fmt.Scan(&number2)

		Transfer(database, id, sum, number2)

	case 3:
		fmt.Println("Input number of Account:")
		var number int64
		fmt.Scan(&number)
		fmt.Println(`Your transactions:`)
		Archive(database, number)

	case 4:
		fmt.Println(`list of ATMs: `)
		ATMs(database)

	case 5:

		AddNewATM(database)
	case 6:
		fmt.Println(`List of users:`)
		Users(database)
	case 7:

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
	var amount, number int64
	row := database.QueryRow(`select amount, number from accaunts where user_id =($1)`, id)
	_ = row.Scan(
		&amount,
		&number,
	)
	newAmount := amount - sum
	if amount < sum {
		log.Println(`not enough money on your account`)
		return
	} else {
		fmt.Println(`number:`, number, `amount: `, amount, `updated amount:`, newAmount)
		_, err := database.Exec(db.UpdateSenderAmount, newAmount, number)
		if err != nil {
			log.Println(`cannot transfer money`, err)
			return
		}
	}

	row = database.QueryRow(`select amount from accaunts where number = ($2) `, number2)
	_ = row.Scan(
		&amount,
	)
	newReceiverAmount := amount + sum
	fmt.Println(`number:`, number2, `amount: `, amount, `amount+sum: `, newReceiverAmount)

	_, err := database.Exec(db.UpdateReceiverAmount, newReceiverAmount, number2)
	if err != nil {
		log.Println(`cannot receive money`, err)
		return
	}
	models.AddTransaction(database, number, number2, sum, newAmount)
}

//3 Оплата услуг

//4 история транзакций
func Archive(database *sql.DB, sender int64) {
	transaction := models.Transaction{}

	rows, err := database.Query(`select id, date,time,amount, sender_number, receiver_number, available_limit from archive  where sender_number = ($1)`, sender)
	if err != nil {
		log.Println(err, `users are not selected`)
	}
	defer rows.Close()
	fmt.Println(`ID 	Date         Time      Amount    Your number      Receiver      Rest `)

	for rows.Next() {
		err := rows.Scan(
			&transaction.ID,
			&transaction.Date,
			&transaction.Time,
			&transaction.Amount,
			&transaction.SenderNumber,
			&transaction.ReceiverNumber,
			&transaction.AvailableLimit,
		)
		if err != nil {
			log.Println(err, ` not selected archive`)
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
		log.Println(err, `not selected`)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&id,
			&address,
		)
		if err != nil {
			log.Println(err, `not selected next`)
		}
		fmt.Println(id, address)
	}
	log.Println(id, address)

}

//6 добавть банкомат
func AddNewATM(database *sql.DB) (ok bool) {

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
			log.Printf(`cant read command: %v`, err)
		}
		fmt.Println(s)
		sprintf := fmt.Sprintf(`%s %s`, s, Address)
		fmt.Println(sprintf)
		_, err = models.AddATM(database, sprintf)
		if err != nil {
			fmt.Println(`not ok`, err)
		} else {
			fmt.Println(`ok`)
		}
	case 2:
		fmt.Println("Goodbye")
		os.Exit(0)
	}
	return
}

//7 список пользователей
func Users(database *sql.DB) {
	var id, age int64
	var name, surname string

	rows, err := database.Query(`select id, age, name, surname from users `)
	if err != nil {
		log.Println(err, `users are not selected`)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
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
func Registration(database *sql.DB) (err error) {
	//var User models.User
	var name, surname, gender, login, password string
	var age int64
	fmt.Println(`name and surname:`)
	fmt.Scan(&name)
	fmt.Scan(&surname)
	fmt.Println(`Age:`)
	fmt.Scan(&age)
	fmt.Println(`gender: `)
	fmt.Scan(&gender)
	fmt.Println(LoginOperation)
	fmt.Println(`login: `)
	fmt.Scan(&login)
	fmt.Println(`password: `)
	fmt.Scan(&password)
	_, s := database.Exec(`insert into users(name, surname,age, gender, login, password) values (($1),($2),($3),($4),($5),($6))`, name, surname, age, gender, login, password)
	if s != nil {
		fmt.Println(`cannot register the new user`, s)
		return s
	}
	fmt.Println(`new user is registered!`)
	return err
}
