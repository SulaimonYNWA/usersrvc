package db

import (
	"database/sql"
	"log"
)

func DBInit(database *sql.DB) {
	DDLs := []string{CreateUsersAccaunt}
	for _, ddl:= range DDLs{
		_, err := database.Exec(ddl)
		if err!= nil{
			log.Fatalf("cant init ,bro %s err is %e", ddl, err)
		}
	}
}
