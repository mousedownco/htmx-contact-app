package contacts

import (
	"html/template"
	"net/http"
)

type Contact struct {
	Id string
}

type Service struct {
	Contacts map[string]Contact
}

func (s *Service) All() []Contact {
	var contacts []Contact
	for _, c := range s.Contacts {
		contacts = append(contacts, c)
	}
	return contacts
}

func (s *Service) Search(_ string) []Contact {
	var contacts []Contact
	return contacts
}

func IndexHandler(svc Service, tmpl *template.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var contacts []Contact
		q := r.URL.Query().Get("q")
		if q != "" {
			contacts = svc.Search(q)
		} else {
			contacts = svc.All()
		}
		e := tmpl.ExecuteTemplate(writer, "contacts", contacts)
		if e != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
