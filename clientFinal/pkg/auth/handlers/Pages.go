package handlers

import (
	"html/template"
	"net/http"
)

func RegisterHTMLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/register.html"))
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func HomeHtmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/home.html"))
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func LoginHtmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/login.html"))
	tmpl.ExecuteTemplate(w, "login.html", nil)
}
