package contacts

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mousedownco/htmx-contact-app/views"
	"net/http"
	"strconv"
)

func HandleIndex(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var contacts []Contact
		page := 1
		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			page, _ = strconv.Atoi(pageParam)
		}
		q := r.URL.Query().Get("q")
		if q != "" {
			contacts = svc.Search(q)
		} else {
			contacts = svc.All(page)
		}
		data := map[string]interface{}{
			"Page":     page,
			"Contacts": contacts,
			"Query":    q,
		}
		view.Render(writer, r, data)
	}
}

func HandleNew(view *views.View) http.HandlerFunc {
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
			views.Flash(w, r, "Created New Contact!")
			http.Redirect(w, r, "/contacts", http.StatusFound)
		}
	}
}

func HandleView(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		view.Render(writer, r, map[string]interface{}{
			"Contact": c,
			"Errors":  map[string]string{},
		})
	}
}

func HandleEdit(svc *Service, view *views.View) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(writer, "Contact Not Found", http.StatusNotFound)
			return
		}
		view.Render(writer, r, map[string]interface{}{
			"Contact": c,
			"Errors":  map[string]string{},
		})
	}
}

func HandleEditPost(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		c.First = r.FormValue("first_name")
		c.Last = r.FormValue("last_name")
		c.Phone = r.FormValue("phone")
		c.Email = r.FormValue("email")

		vErrors := svc.Validate(c)
		if len(vErrors) > 0 {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  vErrors,
			})
		} else {
			e = svc.Save(c)
			if e != nil {
				view.Render(w, r, map[string]interface{}{
					"Contact": c,
					"Errors":  map[string]string{"General": e.Error()},
				})
			}
			views.Flash(w, r, "Updated Contact!")
			http.Redirect(w, r, fmt.Sprintf("/contacts/%d", c.Id), http.StatusFound)
		}
	}
}

func HandleDelete(svc *Service, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		e = svc.Delete(id)
		if e != nil {
			view.Render(w, r, map[string]interface{}{
				"Contact": c,
				"Errors":  map[string]string{"General": e.Error()},
			})
		}
		views.Flash(w, r, "Deleted Contact!")
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	}
}

func HandleEmailGet(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Vars(r)
		id, e := strconv.Atoi(mux.Vars(r)["id"])
		if e != nil {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c := svc.Find(id)
		if (c == Contact{}) {
			http.Error(w, "Contact Not Found", http.StatusNotFound)
			return
		}
		c.Email = r.FormValue("email")
		_, _ = w.Write([]byte(svc.Validate(c)["Email"]))
	}
}
