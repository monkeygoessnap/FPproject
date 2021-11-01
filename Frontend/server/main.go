package server

import (
	"FPproject/Frontend/log"
	"html/template"
	"net/http"
)

var tpl *template.Template

func Run() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)

	http.HandleFunc("/profile", profile)
	http.HandleFunc("/editprofile", editprofile)
	http.HandleFunc("/browse", browse)
	http.HandleFunc("/browse/res", res)
	http.HandleFunc("/cart", cart)

	log.Info.Println("Frontend running at :8181")
	log.Error.Println(http.ListenAndServeTLS(":8181", "certs/cert.pem", "certs/key.pem", nil))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	_, status := newRequest(r, http.MethodGet, "/healthcheck", nil)

	log.Info.Println("Healthcheck code: %v", status)
	tpl.ExecuteTemplate(w, "healthcheck.html", status)
}
