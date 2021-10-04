package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	res, err := d.db.Exec("UPDATE user SET name=?, password=COALESCE(NULLIF(?,''), password), updated=? WHERE id=?",
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

func (d *Database) GetMerchants() ([]models.User, error) {
	var merchant models.User
	var merchants []models.User
	usertype := "merchant"
	r, err := d.db.Query("SELECT * FROM user WHERE type=?", usertype)
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}
	for r.Next() {
		if err := r.Scan(&merchant.ID, &merchant.Username, &merchant.Name, &merchant.Password, &merchant.UserType, &merchant.Created, &merchant.Updated); err != nil {
			log.Warning.Println(err)
			return nil, err
		}
		merchants = append(merchants, merchant)
	}
	return merchants, nil
}

func (d *Database) Validate(um, pw string) (models.User, error) {
	var user models.User
	err := d.db.QueryRow("SELECT * FROM user WHERE username=?", um).Scan(&user.ID,
		&user.Username, &user.Name, &user.Password, &user.UserType, &user.Created, &user.Updated)
	if err != nil {
		log.Warning.Println(err)
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
	if err != nil {
		log.Info.Println(err)
		return user, sql.ErrNoRows
	}
	return user, nil
}
