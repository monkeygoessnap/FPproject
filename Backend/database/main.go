package database

import (
	"FPproject/Backend/log"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dsnStr() string {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PW")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	//formats the connection string
	dsnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		user, pw, port, dbName)

	return dsnStr
}

// func loadSQL() {
// 	c, err := ioutil.ReadFile("./database/FPsql.sql")
// 	if err != nil {
// 		log.Info.Println(err)
// 	}
// 	sql := strings.Split(strings.TrimSpace(string(c)), ";")
// 	for i, v := range sql {
// 		_, err := db.Exec(v + ";")
// 		if err != nil {
// 			log.Info.Println(err)
// 		} else {
// 			fmt.Println("number:", i, v)
// 		}
// 	}
// }

func InitDB() {
	db, _ = sql.Open("mysql", dsnStr())
	if err := db.Ping(); err != nil {
		log.Info.Println(err)
	} else {
		log.Info.Println("Success DB Connection")
	}
	//loadSQL()
}
