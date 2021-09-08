package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"
)

type DBitem struct {
	models.Items
}

func GetAllItem() []DBitem {
	var allItems []DBitem
	var item DBitem
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&item.ID, &item.MerchantID, &item.Name, &item.Price, &item.Status,
			&item.Calories, &item.Created_at, &item.Updated_at)
		allItems = append(allItems, item)
	}
	return allItems
}

func GetItem(id int) DBitem {
	var item DBitem
	if err := db.QueryRow("SELECT * FROM items WHERE id=?", id).Scan(&item.ID, &item.MerchantID, &item.Name, &item.Price, &item.Status,
		&item.Calories, &item.Created_at, &item.Updated_at); err != nil {
		log.Warning.Println(err)
	}
	return item
}

func DelItem(id int) {
	_, err := db.Exec("DELETE FROM items WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func DelAllItems(id int) {
	_, err := db.Exec("DELETE FROM items WHERE merchant_id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func AddItem(item DBitem) {
	_, err := db.Exec("INSERT INTO items(merchant_id,name,price,status,calories,created_at) VALUES (?,?,?,?,?,?)",
		item.MerchantID, item.Name, item.Price, item.Status, item.Calories, time.Now())
	if err != nil {
		log.Warning.Println(err)
	}
}

func UpdateItem(item DBitem) {
	_, err := db.Exec("UPDATE items SET name=?, price=?, status=?, calories=?, updated_at=? WHERE id=?",
		item.Name, item.Price, item.Status, item.Calories, time.Now(), item.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
