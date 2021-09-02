package main

import (
	"FPproject/Backend/database"
	"fmt"
)

func main() {
	database.InitDB()
	fmt.Println("Backend")
}
