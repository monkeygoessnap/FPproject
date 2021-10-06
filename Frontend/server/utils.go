package server

import (
	"FPproject/Frontend/log"
	"FPproject/Frontend/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jasonwinn/geocoder"
	"golang.org/x/crypto/bcrypt"
)

func getToken(r *http.Request) string {
	token, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return token.Value
}

func newRequest(req *http.Request, method, url string, jsonD interface{}) ([]byte, int) {
	base := "http://localhost:8080/api/v1"
	jsonV, _ := json.Marshal(jsonD)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}
	r, _ := http.NewRequest(method, base+url, bytes.NewBuffer(jsonV))
	r.Header.Set("Content-type", "application/json")
	r.Header.Set("access_token", getToken(req))
	resp, err := client.Do(r)
	if err != nil {
		log.Warning.Println(err)
		return nil, http.StatusServiceUnavailable
	}
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, resp.StatusCode
}

func convFloat(input string) float32 {
	output, err := strconv.ParseFloat(input, 32)
	if err != nil {
		log.Warning.Println(err)
		return 0
	}
	return float32(output)
}

func hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}

func setCookie(w http.ResponseWriter, r *http.Request, tk string, expiry string) {

	unix, _ := strconv.Atoi(expiry)
	expiryUnix := time.Unix(int64(unix), 0)

	token := &http.Cookie{
		Name:     "token",
		Value:    tk,
		HttpOnly: true,
		Expires:  expiryUnix,
		Path:     "/",
		Secure:   true,
	}
	http.SetCookie(w, token)
}

func bmi(weight float32, height float32) float32 {
	/* According to HealthHub,
	BMI value of 23 and above indicates that yout weight is outside of the healthy weight range for
	your height
	*/
	return weight / height / height * 10000

}

func ageCal(dob string) int {

	layout := "020106" //ddmmyy
	t1 := time.Now()
	t2, err := time.Parse(layout, dob)
	if err != nil {
		log.Warning.Println(err)
		return 0
	}
	diff := int(t1.Sub(t2).Hours() / 24 / 365)
	return diff
}

func calories(gender, dob, active string, height, weight float32) int {
	//https://www.k-state.edu/paccats/Contents/PA/PDF/Physical%20Activity%20and%20Controlling%20Weight.pdf
	/*activity multiplier
	low = 1.2 x BMR (little to no exercise)
	moderate = 1.55 x BMR (moderate exercise 6 times a week)
	high = 1.9 x BMR (training 2 or more times a day)
	*/
	//Using Harris-Benedict Formula
	age := float32(ageCal(dob))
	var mul float32
	switch active {
	case "low":
		mul = 1.2
	case "moderate":
		mul = 1.55
	case "high":
		mul = 1.9
	default:
		mul = 1.0
	}
	switch gender {
	case "male":
		bmr := 66 + (13.7 * weight) + (5 * height) - (6.8 * age)
		return int(bmr * mul)
	case "female":
		bmr := 655 + (9.6 * weight) + (1.8 * height) - (4.7 * age)
		return int(bmr * mul)
	default:
		return 0
	}
	return 0
}

type Tcal struct {
	Cal    int
	UCal   int
	Target string
	Msg    string
	Color  string
}

func tCal(carts []models.CartItem, foods []models.Food, uh models.UserHealth) Tcal {
	var cl Tcal
	for i, v := range foods {
		cl.Cal = (v.Calories * carts[i].Qty) + cl.Cal
	}
	userCal := calories(uh.Gender, uh.DOB, uh.Active, uh.Height, uh.Weight)
	switch uh.Target {
	case "lose":
		if cl.Cal > userCal {
			cl.Msg = "Calories exceeded!"
			cl.Color = "red"
		} else {
			cl.Msg = "Within calories goal"
			cl.Color = "green"
		}
	case "gain":
		if cl.Cal > userCal {
			cl.Msg = "Calories goal achieved!"
			cl.Color = "green"
		} else {
			cl.Msg = "Calories goal not achieved yet"
			cl.Color = "yellow"
		}
	case "maintain":
		if cl.Cal > int((float32(userCal) * 1.05)) {
			cl.Msg = "Calories exceeded!"
			cl.Color = "red"
		} else if cl.Cal < int((float32(userCal) * 1.05)) {
			cl.Msg = "Calories goal not achieved yet"
			cl.Color = "yellow"
		} else {
			cl.Msg = "Calories goal achieved!"
			cl.Color = "green"
		}
	}
	cl.UCal = userCal
	cl.Target = uh.Target
	return cl
}

func distCal(from string, to string) (float32, float32) {

	//https://www.healthline.com/nutrition/can-you-lose-weight-by-walking-an-hour-a-day#calories-burned
	//average walking pace of 3mph/4.8kph
	//average calories burnt per hour == 193, calories burnt per km == 40.21

	directions := geocoder.NewDirections(from+" SG", []string{to + " SG"})
	distance, err := directions.Distance("k")
	if err != nil {
		return 0, 0
	}
	calBurnt := float32(distance * 40.21)
	return float32(distance), calBurnt
}
