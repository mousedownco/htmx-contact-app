package contacts

import (
	"html/template"
	"net/http"
)

func IndexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		e := tmpl.ExecuteTemplate(writer, "contacts", nil)
		if e != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
