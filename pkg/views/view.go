package views

import (
	"bytes"
	"fmt"
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
	Data map[string]interface{}
}

func (v *View) Render(w http.ResponseWriter, data map[string]interface{}) {
	vd := ViewData{Data: data}
	var rb bytes.Buffer
	e := v.Template.ExecuteTemplate(&rb, v.Layout, vd)
	if e != nil {
		http.Error(w,
			fmt.Sprintf("Error rendering template: %v", e),
			http.StatusInternalServerError)
	} else {
		w.Write(rb.Bytes())
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
