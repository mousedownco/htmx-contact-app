package main

import (
	"github.com/mousedownco/htmx-cognito/pkg/contacts"
	"github.com/mousedownco/htmx-cognito/pkg/views"
	"net/http"
)

func main() {
	cs := contacts.Service{}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("../static/"))))

	http.Handle("/",
		http.RedirectHandler("/contacts", http.StatusTemporaryRedirect))
	http.HandleFunc("/contacts", contacts.IndexHandler(cs,
		views.NewView("layout", "contacts/index.gohtml")))

	_ = http.ListenAndServe(":8080", nil)

}
