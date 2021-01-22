package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "root:newpass@tcp(127.0.0.1:3306)/sampledb")
	checkErr(err)

	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
