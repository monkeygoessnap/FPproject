package models

type Users struct {
	ID         int    `json:"user_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Password   string `json:"-"`
	UserType   string `json:"-"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Address struct {
	ID     int    `json:"add_id"`
	Postal string `json:"postal"`
	Floor  string `json:"floor"`
	Unit   string `json:"unit"`
}

type MerchantDetails struct {
	ID        int    `json:"merchant_id"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type Health struct {
	ID        int     `json:"health_id"`
	Height    float32 `json:"height"`
	Weight    float32 `json:"weight"`
	Age       int     `json:"age"`
	BMI       float32 `json:"bmi"`
	Active    string  `json:"active"`
	TargetW   float32 `json:"target_weight"`
	TargetBMI float32 `json:"target_bmi"`
	TargetCal float32 `json:"target_cal"`
	Reset     string  `json:"reset_time"`
}

type Items struct {
	ID         int     `json:"item_id"`
	MerchantID int     `json:"merchant_id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Status     string  `json:"status"`
	Calories   float32 `json:"calories"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}

type CartItems struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	ItemID  int    `json:"item_id"`
	Qty     int    `json:"qty"`
	Request string `json:"request"`
}

type Order struct {
	ID         int    `json:"order_id"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type OrderItems struct {
	OrderID int    `json:"order_id"`
	ItemID  int    `json:"item_id"`
	Qty     int    `json:"qty"`
	Request string `json:"request"`
}

type OtherRes struct {
	Msg string `json:"msg"`
}
