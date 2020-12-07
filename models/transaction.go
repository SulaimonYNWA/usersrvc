package models

import (
	"Bank/db"
	"database/sql"
	"time"
)

type Transaction struct {
	ID int64
	Date string
	Time string
	Amount int64
	SenderNumber int64
	ReceiverNumber int64
	AvailableLimit int64
}

func AddTransaction(Db *sql.DB, myAccount, receiverAccount int64, operationAmount, newAmount int64) (err error) {
	var check Transaction
	data := time.Now()
	check.Date = data.Format("02-Jan-2006")
	check.Time = data.Format("15:40")
	check.Amount = operationAmount
	check.SenderNumber = myAccount
	check.AvailableLimit = newAmount
	check.ReceiverNumber = receiverAccount
	_, err = Db.Exec(db.AddTransaction, check.Date, check.Time, check.Amount, check.SenderNumber, check.ReceiverNumber, check.AvailableLimit)
	if err != nil {
		panic(err)
	}
	return
}