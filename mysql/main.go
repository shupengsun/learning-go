package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(192.168.162.108:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("SELECT app_id,name FROM alarm_app WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var app_id, name string
	err = stmt.QueryRow(21).Scan(&app_id, &name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("result: %s , %s", app_id, name)

}
