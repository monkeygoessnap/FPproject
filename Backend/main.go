package main

import (
	"FPproject/Backend/database"
	"FPproject/Backend/env"
	"FPproject/Backend/log"
	"fmt"
)

func main() {
	log.InitLog()
	env.InitEnv()
	database.InitDB()
	fmt.Println("Backend")
}
