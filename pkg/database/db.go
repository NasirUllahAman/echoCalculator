package database

import (
	"database/sql"

	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func Conc() *sql.DB {
	fmt.Println("CALEED")
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "Nasir@1234567",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "calculator",
		AllowNativePasswords: true,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// db, err := sql.Open("mysql", "nasir:Nasir@1234567@tcp(127.0.0.1:3306)/calculator")
	if err != nil {
		// log.Fatal("ERORR here")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		// log.Fatal("ERORR here")
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	return db

}
