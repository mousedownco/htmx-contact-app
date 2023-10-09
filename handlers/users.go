package handlers

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed pages/*.gohtml
var pages embed.FS

type UsersHandler struct {
	pages *template.Template
}

func NewUsersHandler() (UsersHandler, error) {
	p, err := template.ParseFS(pages, "pages/*.gohtml")
	if err != nil {
		return UsersHandler{}, err
	}
	return UsersHandler{pages: p}, nil
}

func (h UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	err := h.pages.ExecuteTemplate(w, "users.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
