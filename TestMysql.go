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

	/*查询*/
	hql := "select * from yibi.user where id = 1"
	rows, err := db.Query(hql)
	fmt.Println(rows)
	defer rows.Close()
	/*增加*/
	_, err = db.Exec("insert into yibi.user(id, account) values(null, '20')")
	/*删除*/
	_, err = db.Exec("delete from yibi.user where id = 38")
	for rows.Next() {
		var fileId1 string
		var fileId2 string
		err := rows.Scan(&fileId1, &fileId2)
		fmt.Println(fileId1, fileId2)
		checkErr(err)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
