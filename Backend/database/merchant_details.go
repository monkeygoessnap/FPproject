package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
)

type DBmd struct {
	models.MerchantDetails
}

func GetAllmd() []DBmd {
	var allmd []DBmd
	var md DBmd
	rows, err := db.Query("SELECT * FROM merchant_details")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&md.ID, &md.OpenTime, &md.CloseTime)
		allmd = append(allmd, md)
	}
	return allmd
}

func Getmd(id int) DBmd {
	var md DBmd
	if err := db.QueryRow("SELECT * FROM merchant_details WHERE id=?", id).Scan(&md.ID,
		&md.OpenTime, &md.CloseTime); err != nil {
		log.Warning.Println(err)
	}
	return md
}

func Delmd(id int) {
	_, err := db.Exec("DELETE FROM merchant_details WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func Addmd(md DBmd) {
	_, err := db.Exec("INSERT INTO merchant_details(id, open_time, close_time) VALUES (?,?,?)",
		md.ID, md.OpenTime, md.CloseTime)
	if err != nil {
		log.Warning.Println(err)
	}
}

func Updatemd(md DBmd) {
	_, err := db.Exec("UPDATE merchant_details SET open_time=?, close_time=? WHERE id=?",
		md.OpenTime, md.CloseTime, md.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
