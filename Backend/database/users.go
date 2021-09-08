package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"fmt"
	"time"
)

type DBuser struct {
	*models.Users
}

func GetAllUser() []DBuser {
	var allUsers []DBuser
	users := DBuser{
		&models.Users{},
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&users.ID, &users.Username, &users.Name, &users.Password,
			&users.UserType, &users.Created_at, &users.Updated_at)
		allUsers = append(allUsers, users)
	}
	fmt.Println(allUsers)
	return allUsers
}

func (u *DBuser) Get(id int) {
	if err := db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&u.ID, &u.Username, &u.Name, &u.Password, &u.UserType, &u.Created_at, &u.Updated_at); err != nil {
		log.Warning.Println(err)
	}
}

func (u *DBuser) Delete(id int) {
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func (u *DBuser) Create() {
	_, err := db.Exec("INSERT INTO users(username, full_name, password, type, created_at) VALUES (?,?,?,?,?)",
		u.Username, u.Name, u.Password, u.UserType, time.Now())
	if err != nil {
		log.Warning.Println(err)
	}
}

func (u *DBuser) Update(id int) {
	_, err := db.Exec("UPDATE users SET full_name=?, password=?, type=?, updated_at=? WHERE id=?",
		u.Name, u.Password, u.UserType, time.Now(), id)
	if err != nil {
		log.Warning.Println(err)
	}
}
