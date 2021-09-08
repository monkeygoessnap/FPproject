package database

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
)

type DBHealth struct {
	models.Health
}

func GetAllHealth() []DBHealth {
	var allHealth []DBHealth
	var health DBHealth
	rows, err := db.Query("SELECT * FROM health")
	if err != nil {
		log.Warning.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&health.ID, &health.Height, &health.Weight, &health.Age,
			&health.BMI, &health.Active, &health.TargetW, &health.TargetBMI, &health.TargetCal, &health.Reset)
		allHealth = append(allHealth, health)
	}
	return allHealth
}

func GetHealth(id int) DBHealth {
	var health DBHealth
	if err := db.QueryRow("SELECT * FROM health WHERE id=?", id).Scan(&health.ID, &health.Height, &health.Weight, &health.Age,
		&health.BMI, &health.Active, &health.TargetW, &health.TargetBMI, &health.TargetCal, &health.Reset); err != nil {
		log.Warning.Println(err)
	}
	return health
}

func DelHealth(id int) {
	_, err := db.Exec("DELETE FROM health WHERE id=?", id)
	if err != nil {
		log.Warning.Println(err)
	}
}

func AddHealth(health DBHealth) {
	_, err := db.Exec("INSERT INTO health(id, height, weight, age, bmi, active, target_weight, target_bmi, target_cal, reset) VALUES (?,?,?,?,?,?,?,?,?,?)",
		health.ID, health.Height, health.Weight, health.Age,
		health.BMI, health.Active, health.TargetW, health.TargetBMI, health.TargetCal, &health.Reset)
	if err != nil {
		log.Warning.Println(err)
	}
}

func UpdateHealth(health DBHealth) {
	_, err := db.Exec("UPDATE healths SET height=?, weight=?, age=?, bmi=?, active=?, target_weight=?, target_bmi=?, target_cal=?, reset=? WHERE id=?",
		health.Height, health.Weight, health.Age, health.BMI, health.Active, health.TargetW, health.TargetBMI, health.TargetCal, health.Reset, health.ID)
	if err != nil {
		log.Warning.Println(err)
	}
}
