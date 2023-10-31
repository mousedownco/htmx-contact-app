package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
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
	return &Service{DbPath: dbPath, Contacts: db}
}

type Service struct {
	DbPath   string
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
	vErrors := make(map[string]string)
	if c.Email == "" {
		vErrors["Email"] = "Email Required"
		return vErrors
	}
	for _, dbc := range s.Contacts {
		if c.Email == dbc.Email && c.Id != dbc.Id {
			vErrors["Email"] = "Email Must Be Unique"
			return vErrors
		}
	}
	return vErrors
}

func (s *Service) Save(c Contact) error {
	v := s.Validate(c)
	if len(v) > 0 {
		return errors.New(fmt.Sprintf("unresolved errors: %v", v))
	}
	if c.Id == 0 {
		var maxId int
		for _, dbc := range s.Contacts {
			if dbc.Id > maxId {
				maxId = dbc.Id
			}
		}
		c.Id = maxId + 1
	}
	s.Contacts[c.Id] = c
	return s.SaveDb()
}

func (s *Service) SaveDb() error {
	var contacts []Contact
	for _, c := range s.Contacts {
		contacts = append(contacts, c)
	}
	dbb, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.DbPath, dbb, os.ModePerm)
}

func HandleIndex(svc *Service, view *views.View) http.HandlerFunc {
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
		view.Render(writer, r, data)
	}
}

func HandleNewGet(view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		view.Render(writer, r, map[string]interface{}{
			"Contact": Contact{},
			"Errors":  map[string]string{},
		})
	}
}

func HandleNewPost(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Contact{
			First: r.FormValue("first_name"),
			Last:  r.FormValue("last_name"),
			Phone: r.FormValue("phone"),
			Email: r.FormValue("email"),
		}
		vErrors := svc.Validate(c)
		if len(vErrors) > 0 {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  vErrors,
			})
		} else {
			e := svc.Save(c)
			if e != nil {
				view.Render(w, r, map[string]interface{}{
					"Contact": c,
					"Errors":  map[string]string{"General": e.Error()},
				})
			}
			views.Flash(w, "Created New Contact!")
			http.Redirect(w, r, "/contacts", http.StatusTemporaryRedirect)
		}
	}
}
