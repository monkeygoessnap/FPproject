package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"time"

	"github.com/google/uuid"
)

func (d *Database) InsertFood(f models.Food) (string, error) {
	id := uuid.New().String()
	res, err := d.db.Exec("INSERT INTO food(id, merchant_id, name, price, status, description, imglink, created, updated) VALUES(?,?,?,?,?,?,?,?,?)",
		id, f.MerchantID, f.Name, f.Price, f.Status, f.Description, f.ImgLink, time.Now(), time.Now())
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

func (d *Database) DelFood(id string) (string, error) {
	res, err := d.db.Exec("DELETE FROM food WHERE id=?", id)
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

func (d *Database) UpdateFood(f models.Food) (string, error) {
	res, err := d.db.Exec("UPDATE food SET name=?, price=?, status=?, description=?, imglink=?, updated=? WHERE id=?",
		f.Name, f.Price, f.Status, f.Description, f.ImgLink, time.Now(), f.ID)
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
	return f.ID, nil
}

func (d *Database) GetFood(id string) (models.Food, error) {
	var f models.Food
	err := d.db.QueryRow("SELECT * FROM food WHERE id=?", id).Scan(&f.ID,
		&f.MerchantID, &f.Name, &f.Price, &f.Status, &f.Description, &f.ImgLink, &f.Created, &f.Updated)
	if err != nil {
		log.Warning.Println(err)
		return f, err
	}
	return f, nil
}

func (d *Database) GetFoodByMerchant(id string) ([]models.Food, error) {
	var food models.Food
	var foods []models.Food
	r, err := d.db.Query("SELECT * FROM food WHERE merchant_id=?", id)
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}
	for r.Next() {
		if err := r.Scan(&food.ID, &food.MerchantID, &food.Name, &food.Price, &food.Status, &food.Description, &food.ImgLink, &food.Created, &food.Updated); err != nil {
			log.Warning.Println(err)
			return nil, err
		}
		foods = append(foods, food)
	}
	return foods, nil
}
