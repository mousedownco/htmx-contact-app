package views

import (
	"bytes"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"path/filepath"
)

var TemplatesDir = "templates"
var LayoutDir = "layout"
var FlashName = "_flash"

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	tmplFiles, e := layoutFiles(filepath.Join(TemplatesDir, LayoutDir))
	if e != nil {
		panic(e)
	}
	tmplFiles = append(tmplFiles, viewFiles(files)...)
	tmpl, e := template.ParseFiles(tmplFiles...)
	if e != nil {
		panic(e)
	}
	return &View{Template: tmpl, Layout: layout}
}

type ViewData struct {
	Data  map[string]interface{}
	Flash string
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	flash := GetFlash(w, r)
	vd := ViewData{Data: data, Flash: flash}
	var rb bytes.Buffer
	e := v.Template.ExecuteTemplate(&rb, v.Layout, vd)
	if e != nil {
		http.Error(w,
			fmt.Sprintf("Error rendering template: %v", e),
			http.StatusInternalServerError)
	} else {
		_, _ = w.Write(rb.Bytes())
	}
}

func layoutFiles(dir string) ([]string, error) {
	return filepath.Glob(dir + "/*.gohtml")
}

func viewFiles(files []string) []string {
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(TemplatesDir, file))
	}
	return paths
}

var store = sessions.NewCookieStore([]byte("a-secret-string"))

func Flash(w http.ResponseWriter, r *http.Request, value string) {
	session, e := store.Get(r, FlashName)
	if e != nil {
		fmt.Printf("Error getting session: %v\n", e)
		return
	}
	session.AddFlash(value, "message")
	e = session.Save(r, w)
	if e != nil {
		fmt.Printf("Error saving session: %v\n", e)
	}
}

func GetFlash(w http.ResponseWriter, r *http.Request) string {
	session, e := store.Get(r, FlashName)
	if e != nil {
		fmt.Printf("Error loading session: %v\n", e)
		return ""
	}

	fm := session.Flashes("message")
	if fm == nil {
		return ""
	}

	_ = session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}
