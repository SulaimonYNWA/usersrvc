package models

import (
	"Bank/db"
	"database/sql"
	"log"
)

type ATMs struct {
	ID int64
	Address string
	Works bool
}

func AddATM(database *sql.DB, Address string)(ok bool, err error)  {
	//rows, err := database.Query(`select * from ATMs`)
	_, err = database.Exec(db.AddNewATM, Address)
	if err!=nil{
		log.Println(`cant insert my db`, err)
	}
	return
}
