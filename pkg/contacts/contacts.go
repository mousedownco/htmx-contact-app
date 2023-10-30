package contacts

import (
	"encoding/json"
	"github.com/mousedownco/htmx-cognito/pkg/views"
	"net/http"
	"os"
	"strings"
)

type Contact struct {
	Id    int    `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func NewService(dbPath string) *Service {
	dbb, err := os.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}
	var contacts []Contact
	err = json.Unmarshal(dbb, &contacts)
	if err != nil {
		panic(err)
	}
	db := make(map[int]Contact)
	for _, c := range contacts {
		db[c.Id] = c
	}
	return &Service{Contacts: db}
}

type Service struct {
	Contacts map[int]Contact
}

func (s *Service) All() []Contact {
	var contacts []Contact
	for _, c := range s.Contacts {
		contacts = append(contacts, c)
	}
	return contacts
}

func (s *Service) Search(q string) []Contact {
	var results []Contact
	for _, c := range s.Contacts {
		first := strings.Contains(strings.ToLower(c.First), strings.ToLower(q))
		last := strings.Contains(strings.ToLower(c.Last), strings.ToLower(q))
		phone := strings.Contains(strings.ToLower(c.Phone), strings.ToLower(q))
		email := strings.Contains(strings.ToLower(c.Email), strings.ToLower(q))
		if first || last || phone || email {
			results = append(results, c)
		}
	}
	return results
}

func (s *Service) Validate(c Contact) map[string]string {
	errors := make(map[string]string)
	if c.First == "" {
		errors["First"] = "First name is required"
	}
	if c.Last == "" {
		errors["Last"] = "Last name is required"
	}
	if c.Phone == "" {
		errors["Phone"] = "Phone number is required"
	}
	if c.Email == "" {
		errors["Email"] = "Email address is required"
	}
	return errors
}

func IndexHandler(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var contacts []Contact
		q := r.URL.Query().Get("q")
		if q != "" {
			contacts = svc.Search(q)
		} else {
			contacts = svc.All()
		}
		data := map[string]interface{}{
			"Contacts": contacts,
			"Query":    q,
		}
		view.Render(writer, data)
	}
}

func AddHandler(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			data := map[string]interface{}{
				"Contact": Contact{},
				"Errors":  map[string]string{},
			}
			view.Render(writer, data)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
