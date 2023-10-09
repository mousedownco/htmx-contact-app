package main

import (
	"github.com/mousedownco/htmx-cognito/handlers"
	"log"
	"net/http"
)

func main() {
	uh, e := handlers.NewUsersHandler()
	if e != nil {
		log.Fatal(e)
	}

	http.HandleFunc("/", uh.Get)

	http.ListenAndServe(":8080", nil)

}
