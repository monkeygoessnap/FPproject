package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"
)

type DBuser struct {
	models.Users
}

func GetAllUser() []DBuser {
	var allUsers []DBuser
	var user DBuser
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password,
			&user.UserType, &user.Created_at, &user.Updated_at)
		allUsers = append(allUsers, user)
	}
	return allUsers
}

func GetUser(id int) DBuser {
	var user DBuser
	if err := db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&user.ID,
		&user.Username, &user.Name, &user.Password, &user.UserType, &user.Created_at, &user.Updated_at); err != nil {
		log.Warning.Println(err)
	}
	return user
}

func DelUser(id int) {
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func AddUser(user DBuser) {
	_, err := db.Exec("INSERT INTO users(username, full_name, password, type, created_at) VALUES (?,?,?,?,?)",
		user.Username, user.Name, user.Password, user.UserType, time.Now())
	if err != nil {
		log.Warning.Println(err)
	}
}

func UpdateUser(user DBuser) {
	_, err := db.Exec("UPDATE users SET full_name=?, password=?, type=?, updated_at=? WHERE id=?",
		user.Name, user.Password, user.UserType, time.Now(), user.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
