package server

import (
	"FPproject/Frontend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func cart(w http.ResponseWriter, r *http.Request) {

	var tpldata []interface{}
	var cart []models.CartItem
	data, status := newRequest(r, http.MethodGet, "/allci", nil)
	if status != 200 {
		tpl.ExecuteTemplate(w, "err.html", nil)
		return
	}
	json.Unmarshal(data, &cart)
	var foods []models.Food
	for _, v := range cart {
		var food models.Food
		fooddata, _ := newRequest(r, http.MethodGet, "/food/"+v.ID, nil)
		json.Unmarshal(fooddata, &food)
		foods = append(foods, food)
	}
	var uh models.UserHealth
	data, _ = newRequest(r, http.MethodGet, "/uh", nil)
	json.Unmarshal(data, &uh)

	calData := tCal(cart, foods, uh)
	tpldata = append(tpldata, cart, foods, calData)

	if r.Method == http.MethodPost {
		if r.FormValue("submit") == "order" {
			_, status := newRequest(r, http.MethodDelete, "/ci", nil)
			if status != 200 {
				tpl.ExecuteTemplate(w, "err.html", nil)
				return
			}
		}
		if r.FormValue("delete") != "" {
			id := r.FormValue("delete")
			_, status := newRequest(r, http.MethodDelete, "/ci/"+id, nil)
			if status != 200 {
				tpl.ExecuteTemplate(w, "err.html", nil)
				return
			}
		}
		if r.FormValue("edit") != "" {
			id := r.FormValue("edit")
			qty, _ := strconv.Atoi(r.FormValue(id))
			new := models.CartItem{
				ID:  id,
				Qty: qty,
			}
			_, status := newRequest(r, http.MethodPut, "/ci", new)
			if status != 200 {
				tpl.ExecuteTemplate(w, "err.html", nil)
				return
			}
		}
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return

	}

	tpl.ExecuteTemplate(w, "cart.html", tpldata)
}

func res(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var tpldata []interface{}
	name := r.URL.Query().Get("name")
	var add models.Address
	data, _ := newRequest(r, http.MethodGet, "/mercadd/"+id, nil)
	json.Unmarshal(data, &add)
	var foods []models.Food
	data, _ = newRequest(r, http.MethodGet, "/allfood/"+id, nil)
	json.Unmarshal(data, &foods)
	var uadd models.Address
	udata, _ := newRequest(r, http.MethodGet, "/add", nil)
	json.Unmarshal(udata, &uadd)
	fmt.Println(uadd.Postal, add.Postal)
	dist, cal := distCal(uadd.Postal, add.Postal)
	dc := map[string]float32{
		"distance": dist,
		"cal":      cal,
	}
	tpldata = append(tpldata, name, add, foods, dc)

	if r.Method == http.MethodPost {
		id := r.FormValue("add")
		qty, _ := strconv.Atoi(r.FormValue(id))
		new := models.CartItem{
			ID:  id,
			Qty: qty,
		}
		_, status := newRequest(r, http.MethodPost, "/ci", new)
		if status != 200 {
			tpl.ExecuteTemplate(w, "err.html", nil)
			return
		}
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		//http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "res.html", tpldata)
}

func browse(w http.ResponseWriter, r *http.Request) {

	var mercs []models.User
	data, _ := newRequest(r, http.MethodGet, "/merc", nil)
	json.Unmarshal(data, &mercs)
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		http.Redirect(w, r, "/browse/res?id="+id+"&name="+name, http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "browse.html", mercs)
}
