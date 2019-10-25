package views

import (
	"net/http"
	"text/template"
)

// HomeView generates the main page for the application
func HomeView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/layout.html", "assets/index.html"))
	tmpl.ExecuteTemplate(w, "layout", "")
}
