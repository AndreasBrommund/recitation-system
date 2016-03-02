package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/DavidSkeppstedt/recitation/db"
	"github.com/gorilla/sessions"
)

var database db.Database
var store = sessions.NewCookieStore([]byte("DD1368"))

//Should do this recursivly for the whole views folder...
var templates = template.Must(
	template.ParseFiles(
		"views/admin.html",
		"views/course.html",
		"views/recitation.html",
		"views/enroll.html",
		"views/profile.html",
		"views/student.html"))

func StartServer(port string) {
	var err error
	database, err = db.NewDatabase("./config/db.json", "dev")
	if err != nil {
		log.Fatal(err)
	}
	router := NewRouter()
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
