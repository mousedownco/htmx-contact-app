package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const PageSize = 10

type Contact struct {
	Id    int    `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func NewService(dbPath string) *Service {
	dbb, e := os.ReadFile(dbPath)
	if e != nil {
		panic(e)
	}
	var contacts []Contact
	e = json.Unmarshal(dbb, &contacts)
	if e != nil {
		panic(e)
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

func (s *Service) All(page ...int) []Contact {
	var contacts []Contact
	for _, c := range s.Contacts {
		contacts = append(contacts, c)
	}
	sort.Slice(contacts, func(l, r int) bool {
		return contacts[l].Id < contacts[r].Id
	})
	if len(page) > 0 {
		start := PageSize * (page[0] - 1)
		end := start + PageSize
		if start > len(contacts) {
			return []Contact{}
		}
		if end > len(contacts) {
			end = len(contacts)
		}
		contacts = contacts[start:end]
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

func (s *Service) Find(id int) Contact {
	return s.Contacts[id]
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

func (s *Service) Save(c Contact) (Contact, error) {
	v := s.Validate(c)
	if len(v) > 0 {
		return c, errors.New(fmt.Sprintf("unresolved errors: %v", v))
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
	return c, s.SaveDb()
}

func (s *Service) Delete(id int) error {
	delete(s.Contacts, id)
	return s.SaveDb()
}

func (s *Service) SaveDb() error {
	dbb, e := json.MarshalIndent(s.All(), "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(s.DbPath, dbb, os.ModePerm)
}
