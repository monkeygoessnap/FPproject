package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"
)

func (d *Database) InsertAdd(id string, add models.Address) (string, error) {
	res, err := d.db.Exec("INSERT INTO address(id, postal, floor, unit, created, updated) VALUES(?,?,?,?,?,?)",
		id, add.Postal, add.Floor, add.Unit, time.Now(), time.Now())
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

func (d *Database) DelAdd(id string) (string, error) {
	res, err := d.db.Exec("DELETE FROM address WHERE id=?", id)
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

func (d *Database) UpdateAdd(add models.Address) (string, error) {
	res, err := d.db.Exec("UPDATE address SET postal=?, floor=?, unit=?, updated=? WHERE id=?",
		add.Postal, add.Floor, add.Unit, time.Now(), add.ID)

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
	return add.ID, nil
}

func (d *Database) GetAdd(id string) (models.Address, error) {
	var add models.Address
	err := d.db.QueryRow("SELECT * FROM address WHERE id=?", id).Scan(&add.ID,
		&add.Postal, &add.Floor, &add.Unit, &add.Created, &add.Updated)
	if err != nil {
		log.Warning.Println(err)
		return add, err
	}
	return add, nil
}
