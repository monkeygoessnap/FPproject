package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"

	"github.com/google/uuid"
)

func (d *Database) InsertUser(user models.User) (string, error) {
	id := uuid.New().String()
	res, err := d.db.Exec("INSERT INTO user(id, username, name, password, type, created, updated) VALUES(?,?,?,?,?,?,?)",
		id, user.Username, user.Name, user.Password, user.UserType, time.Now(), time.Now())
	if err != nil {
		log.Warning.Println(err)
		return "", err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Warning.Println(err)
		return "", err
	} else if affected < 1 {
		log.Warning.Println(ErrNoRowsAffected)
		return "", ErrNoRowsAffected
	}
	return id, nil
}

func (d *Database) DelUser(id string) (string, error) {
	res, err := d.db.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
		return "", err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Warning.Println(err)
		return "", err
	} else if affected < 1 {
		log.Warning.Println(ErrNoRowsAffected)
		return "", ErrNoRowsAffected
	}
	return id, nil
}

func (d *Database) UpdateUser(user models.User) (string, error) {
	res, err := d.db.Exec("UPDATE user SET name=?, password=?, updated=? WHERE id=?",
		user.Name, user.Password, time.Now(), user.ID)
	if err != nil {
		log.Warning.Println(err)
		return "", err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Warning.Println(err)
		return "", err
	} else if affected < 1 {
		log.Warning.Println(ErrNoRowsAffected)
		return "", ErrNoRowsAffected
	}
	return user.ID, nil
}

func (d *Database) GetUser(id string) (models.User, error) {
	var user models.User
	err := d.db.QueryRow("SELECT * FROM user WHERE id=?", id).Scan(&user.ID,
		&user.Username, &user.Name, &user.Password, &user.UserType, &user.Created, &user.Updated)
	if err != nil {
		log.Warning.Println(err)
		return user, err
	}
	return user, nil
}

// func GetAllUser() []DBuser {
// 	var allUsers []DBuser
// 	var user DBuser
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		log.Warning.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password,
// 			&user.UserType, &user.Created_at, &user.Updated_at)
// 		allUsers = append(allUsers, user)
// 	}
// 	return allUsers
// }
