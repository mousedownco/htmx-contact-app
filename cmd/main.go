package main

import (
	"github.com/mousedownco/htmx-cognito/pkg/contacts"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var tmplFiles []string

	tmplFiles, e := templateFiles("templates")
	if e != nil {
		log.Fatal(e)
	}

	tmpls, e := template.ParseFiles(tmplFiles...)
	if e != nil {
		log.Fatal(e)
	}

	cs := contacts.Service{}

	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", http.RedirectHandler("/contacts", http.StatusTemporaryRedirect))
	http.HandleFunc("/contacts", contacts.IndexHandler(cs, tmpls))

	_ = http.ListenAndServe(":8080", nil)

}

func templateFiles(dir string) ([]string, error) {
	var templateFiles []string
	e := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".gohtml") {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return templateFiles, nil
}
