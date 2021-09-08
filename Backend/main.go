package main

import (
	"FPproject/Backend/database"
	"FPproject/Backend/env"
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"fmt"
)

func main() {
	log.InitLog()
	env.InitEnv()
	database.InitDB()
	fmt.Println("Backend")
	newUser := database.DBuser{
		&models.Users{
			Username: "user10",
			Name:     "username2",
			Password: "pass",
			UserType: "admin",
		},
	}
	fmt.Println(newUser.Username, "\n\n")
	//newUser.Create()
	for _, v := range database.GetAllUser() {
		fmt.Println(v.Username)
	}
	fmt.Println(newUser.Username)
	fmt.Println(newUser.Created_at)
	// var newUser2 database.DBuserm
	// newUser2.Get(1)
	// fmt.Println(&newUser2)
}
