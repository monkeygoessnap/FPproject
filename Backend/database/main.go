package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	db, err := sql.Open("mysql", "root:password@tcp(0.0.0.0:55241)/FP")
	defer db.Close()
	if err != nil {
		fmt.Println("DB Err")
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Ping Err")
	}

	c, err := ioutil.ReadFile("D:\\Projects\\Go\\src\\FPproject\\Backend\\database\\FPsql.sql")
	if err != nil {
		fmt.Println(err)
	}
	sql := string(c)
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successful InitDB")
}
