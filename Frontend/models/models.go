package models

type Track struct {
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type User struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType string `json:"type"` //enum admin, merchant, customer
	Track
}

type Address struct {
	ID     string `json:"add_id"` //ref user.id
	Postal string `json:"postal"`
	Floor  string `json:"floor"`
	Unit   string `json:"unit"`
	Track
}

type UserHealth struct {
	ID     string  `json:"user_id"` //ref user.id
	Gender string  `json:"gender"`  //enum male, female
	Height float32 `json:"height"`  //metric cm
	Weight float32 `json:"weight"`  //metric kg
	DOB    string  `json:"dob"`     //ddmmyy format
	Active string  `json:"active"`  //enum low, moderate, high
	Target string  `json:"target"`  //enum gain, lose, maintain
	Track
}

type Food struct { //need getall
	ID          string  `json:"food_id"`
	MerchantID  string  `json:"merchant_id"` //ref user.id
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Status      string  `json:"status"` //enum avail, soldout
	Description string  `json:"description"`
	ImgLink     string  `json:"imglink"` //image url
	Calories    int     `json:"calories"`
	Track
}

type CartItem struct { //need getall
	ID      string `json:"food_id"` //ref food.id
	UserID  string `json:"user_id"` //ref user.id
	Qty     int    `json:"qty"`
	Remarks string `json:"remarks"`
	Track
}

type AddData struct {
	Calories int
	Age      int
	BMI      float32
}

// type Order struct { //need getall
// 	ID     string `json:"order_id"`
// 	UserID string `json:"user_id"` //ref user.id
// 	Status string `json:"status"`  //enum completed, pending
// 	Track
// }

// type OrderItem struct { //need getall
// 	Food
// 	OrderID string `json:"order_id"` //ref order_id
// 	Qty     int `json:"qty"`
// 	Remarks string `json:"remarks"`
// 	Track
// }
