package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
)

type DBadd struct {
	models.Address
}

func GetAllAdd() []DBadd {
	var allAdd []DBadd
	var add DBadd
	rows, err := db.Query("SELECT * FROM address")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&add.ID, &add.Postal, &add.Floor, &add.Unit)
		allAdd = append(allAdd, add)
	}
	return allAdd
}

func GetAdd(id int) DBadd {
	var add DBadd
	if err := db.QueryRow("SELECT * FROM address WHERE id=?", id).Scan(&add.ID,
		&add.Postal, &add.Floor, &add.Unit); err != nil {
		log.Warning.Println(err)
	}
	return add
}

func DelAdd(id int) {
	_, err := db.Exec("DELETE FROM address WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func AddAdd(add DBadd) {
	_, err := db.Exec("INSERT INTO address(id, postal, floor, unit) VALUES (?,?,?,?)",
		add.ID, add.Postal, add.Floor, add.Unit)
	if err != nil {
		log.Warning.Println(err)
	}
}

func UpdateAdd(add DBadd) {
	_, err := db.Exec("UPDATE address SET postal=?, floor=?, unit=? WHERE id=?",
		add.Postal, add.Floor, add.Unit, add.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
