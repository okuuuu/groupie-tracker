package server

import (
	"html/template"
	"main/res"
	"net/http"
	"regexp"
	"strconv"
)

type M map[string]interface{}

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := res.GetArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error!\n" + err.Error()))
	}
	var artistNumber int
	var display string
	if r.Method == "POST" {
		artistNumber, _ = strconv.Atoi(r.FormValue("artistID"))
		display = "block"
	} else {
		artistNumber = 1
		display = "none"
	}
	relation, err := res.GetRelation(artistNumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error!\n" + err.Error()))
	}

	renderTemplate(w, "index", M{
		"artist":   (*artists)[artistNumber-1],
		"artists":  artists,
		"relation": relation,
		"display":  display,
	})
}

var validPath = regexp.MustCompile("^/($)")

// MakeHandler creates a handler function.
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Redirect(w, r, "https://http.cat/404", http.StatusSeeOther)
			return
		}
		fn(w, r)
	}
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

// renderTemplate renders the given template with the given data.
func renderTemplate(w http.ResponseWriter, tmpl string, p any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
