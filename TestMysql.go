package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/yibi")
	checkErr(err)

	err = db.Ping()
	if err == nil {
		fmt.Println("ping is ok")
	}
	checkErr(err)
	defer db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
