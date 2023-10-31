package views

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
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
	fmt.Printf("rendering with flash: %q\n", flash)
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

func Flash(w http.ResponseWriter, value string) {
	fmt.Printf("setting flash: %q\n", value)
	c := &http.Cookie{
		Name:  FlashName,
		Value: base64.URLEncoding.EncodeToString([]byte(value))}
	http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request) string {
	c, e := r.Cookie(FlashName)
	if e != nil {
		if !errors.Is(e, http.ErrNoCookie) {
			fmt.Printf("Error getting flash cookie: %v", e)
		}
		return ""
	}
	value, e := base64.URLEncoding.DecodeString(c.Value)
	if e != nil {
		fmt.Printf("Error decoding flash cookie: %v", e)
		return ""
	}
	dc := &http.Cookie{
		Name:    FlashName,
		MaxAge:  -1,
		Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return string(value)
}
