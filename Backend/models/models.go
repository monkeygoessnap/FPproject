package models

type Track struct {
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"-"`
	UserType string `json:"type"` //enum admin, merchant, customer
	Track
}

type Address struct {
	ID     int    `json:"add_id"` //ref user.id
	Postal string `json:"postal"`
	Floor  string `json:"floor"`
	Unit   string `json:"unit"`
	Track
}

type UserHealth struct {
	ID     int     `json:"user_id"` //ref user.id
	Gender string  `json:"gender"`  //enum male, female
	Height float32 `json:"height"`  //metric cm
	Weight float32 `json:"weight"`  //metric kg
	DOB    int     `json:"dob"`     //ddmmyy format
	Active string  `json:"active"`  //enum low, moderate, high
	Target string  `json:"target"`  //enum gain, lose, maintain
	Track
}

type Food struct {
	ID          int     `json:"food_id"`
	MerchantID  int     `json:"merchant_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Status      string  `json:"status"` //enum avail, soldout
	Description string  `json:"description"`
	ImgLink     string  `json:"imglink"` //image url
	Track
}

type CartItem struct {
	Food
	UserID  int    `json:"user_id"` //ref user.id
	Qty     int    `json:"qty"`
	Remarks string `json:"remarks"`
	Track
}

type Order struct {
	ID     int    `json:"order_id"`
	UserID int    `json:"user_id"` //ref user.id
	Status string `json:"status"`  //enum completed, pending
	Track
}

type OrderItem struct {
	Food
	OrderID int    `json:"order_id"` //ref order_id
	Qty     int    `json:"qty"`
	Remarks string `json:"remarks"`
	Track
}

type OtherRes struct {
	Msg string `json:"msg"`
}
