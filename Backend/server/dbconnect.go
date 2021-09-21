package server

import (
	"FPproject/Backend/log"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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

func connectDB() *sql.DB {
	db, _ := sql.Open("mysql", dsnStr())
	if err := db.Ping(); err != nil {
		log.Error.Fatal(err)
	} else {
		log.Info.Println("Success DB Connection")
	}
	return db
}
