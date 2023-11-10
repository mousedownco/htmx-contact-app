package main

import (
	"github.com/gorilla/mux"
	"github.com/mousedownco/htmx-contact-app/contacts"
	"github.com/mousedownco/htmx-contact-app/views"
	"log"
	"net/http"
)

var staticDir = "static"
var port = ":8080"

func main() {
	cs := contacts.NewService("contacts.json")

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(staticDir))))
	r.Handle("/",
		http.RedirectHandler("/contacts", http.StatusTemporaryRedirect))
	r.Handle("/contacts",
		contacts.HandleIndex(cs, views.NewView("layout", "contacts/index.gohtml")))
	r.Handle("/contacts/new",
		contacts.HandleNew(views.NewView("layout", "contacts/new.gohtml"))).Methods("GET")
	r.Handle("/contacts/new",
		contacts.HandleNewPost(cs, views.NewView("layout", "contacts/new.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}",
		contacts.HandleView(cs, views.NewView("layout", "contacts/show.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts.HandleEdit(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("GET")
	r.Handle("/contacts/{id:[0-9]+}/edit",
		contacts.HandleEditPost(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("POST")
	r.Handle("/contacts/{id:[0-9]+}/delete",
		contacts.HandleDeletePost(cs, views.NewView("layout", "contacts/edit.gohtml"))).Methods("POST")
	log.Printf("Starting server on port %s", port)
	http.Handle("/", r)
	_ = http.ListenAndServe(port, nil)

}
