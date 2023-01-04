package server

import (
	"html/template"
	"main/res"
	"net/http"
	"regexp"
)

// Ascii represents the ASCII art entity.
type Ascii struct {
	Word   string
	Art    string
	Banner string
}

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := res.GetArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error!\n" + err.Error()))
	}

	renderTemplate(w, "index", artists)
}

var validPath = regexp.MustCompile("^/(concerts|$)")

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
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
