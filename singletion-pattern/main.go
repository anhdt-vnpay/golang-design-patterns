package main

import (
	"database/sql"
	"fmt"
	"sync"
)

var db *sql.DB
var once sync.Once
var err error

func Init() (*sql.DB, error) {
	once.Do(func() {

		host := "localhost"
		port := 5432
		user := "postgres"
		password := "1"
		dbname := "postgres"

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)
	})
	return db, err
}

func main() {

	db, err := Init()

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully DB connected!")
}
