package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"
)

func (d *Database) InsertUH(id string, h models.UserHealth) (string, error) {
	res, err := d.db.Exec("INSERT INTO userhealth(id, gender, height, weight, dob, active, target, created, updated) VALUES(?,?,?,?,?,?,?,?,?)",
		id, h.Gender, h.Height, h.Weight, h.DOB, h.Active, h.Target, time.Now(), time.Now())
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

func (d *Database) DelUH(id string) (string, error) {
	res, err := d.db.Exec("DELETE FROM userhealth WHERE id=?", id)
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

func (d *Database) UpdateUH(h models.UserHealth) (string, error) {
	res, err := d.db.Exec("UPDATE userhealth SET gender=?, height=?, weight=?, dob=?, active=?, target=?, updated=? WHERE id=?",
		h.Gender, h.Height, h.Weight, h.DOB, h.Active, h.Target, time.Now(), h.ID)
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
	return h.ID, nil
}

func (d *Database) GetUH(id string) (models.UserHealth, error) {
	var h models.UserHealth
	err := d.db.QueryRow("SELECT * FROM userhealth WHERE id=?", id).Scan(&h.ID,
		&h.Gender, &h.Height, &h.Weight, &h.DOB, &h.Active, &h.Target, &h.Created, &h.Updated)
	if err != nil {
		log.Warning.Println(err)
		return h, err
	}
	return h, nil
}
