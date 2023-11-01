package main

import (
	"github.com/gorilla/mux"
	contacts2 "github.com/mousedownco/htmx-cognito/contacts"
	"github.com/mousedownco/htmx-cognito/views"
	"log"
	"net/http"
)

var staticDir = "static"
var port = ":8080"

func main() {
	cs := contacts2.NewService("contacts.json")

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(staticDir))))
	r.Handle("/",
		http.RedirectHandler("/contacts", http.StatusTemporaryRedirect))
	r.Handle("/contacts",
		contacts2.HandleIndex(cs, views.NewView("layout", "contacts/index.gohtml")))
	r.Handle("/contacts/new",
		contacts2.HandleNew(views.NewView("layout", "contacts/new.gohtml"))).Methods("GET")
	r.Handle("/contacts/new",
		contacts2.HandleNewPost(cs, views.NewView("layout", "contacts/new.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}",
		contacts2.HandleView(cs, views.NewView("layout", "contacts/show.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts2.HandleEdit(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts2.HandleEditPost(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}/delete",
		contacts2.HandleDeletePost(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("POST")
	log.Printf("Starting server on port %s", port)
	http.Handle("/", r)
	_ = http.ListenAndServe(port, nil)

}
