package users

import (
	"html/template"
	"net/http"
)

func IndexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := tmpl.ExecuteTemplate(w, "users/index", nil)
		if e != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
