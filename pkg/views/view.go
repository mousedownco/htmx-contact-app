package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var TemplatesDir = "templates"
var LayoutDir = "layout"

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
	Data interface{}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{Data: data}
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
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
