package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
)

type DBci struct {
	models.CartItems
}

func GetAllCi(id int) []DBci {
	var allCis []DBci
	var ci DBci
	rows, err := db.Query("SELECT * FROM cart_items WHERE user_id=?", id)
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&ci.ID, &ci.UserID, &ci.ItemID, &ci.Qty, &ci.Request)
		allCis = append(allCis, ci)
	}
	return allCis
}

func GetCi(id int) DBci {
	var ci DBci
	if err := db.QueryRow("SELECT * FROM cart_items WHERE id=?", id).Scan(&ci.ID, &ci.UserID, &ci.ItemID, &ci.Qty, &ci.Request); err != nil {
		log.Warning.Println(err)
	}
	return ci
}

func DelCi(id int) {
	_, err := db.Exec("DELETE FROM cart_items WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func DelAllCis(id int) {
	_, err := db.Exec("DELETE FROM cart_items WHERE user_id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func AddCi(ci DBci) {
	_, err := db.Exec("INSERT INTO cart_items(user_id,item_id,quantity,request) VALUES (?,?,?,?)",
		ci.UserID, ci.ItemID, ci.Qty, ci.Request)
	if err != nil {
		log.Warning.Println(err)
	}
}

func UpdateCi(ci DBci) {
	_, err := db.Exec("UPDATE cart_items SET quantity=?,request=? WHERE id=?",
		ci.Qty, ci.Request, ci.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
