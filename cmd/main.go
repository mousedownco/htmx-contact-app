package main

import (
	"github.com/gorilla/mux"
	"github.com/mousedownco/htmx-cognito/pkg/contacts"
	"github.com/mousedownco/htmx-cognito/pkg/views"
	"net/http"
)

var staticDir = "static"

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
		contacts.HandleNewGet(views.NewView("layout", "contacts/new.gohtml"))).Methods("GET")
	r.Handle("/contacts/new",
		contacts.HandleNewPost(cs, views.NewView("layout", "contacts/new.gohtml"))).Methods("POST")
	http.Handle("/", r)
	_ = http.ListenAndServe(":8080", nil)

}
