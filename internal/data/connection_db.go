package data

import (
	"database/sql"
	"log"
)

func ConnectionDB(driver, url string) (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		log.Panicf("LOG - [DATABASE-PANIC]")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("LOG - [DATABASE-ERROR]: error when try connection on database")
	} else {
		log.Println("LOG - [DATABASE]: connection on database is good")
	}

	return db, err
}
