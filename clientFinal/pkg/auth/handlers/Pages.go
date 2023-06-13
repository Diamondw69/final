package handlers

import (
	pb "clientFinal/pkg/auth/proto"
	"context"
	"html/template"
	"log"
	"net/http"
)

//We need Pages.go to group handler that only used to show html pages

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

// ProfileHtmlHandler method uses token to get data about user from server
func ProfileHtmlHandler(w http.ResponseWriter, r *http.Request) {

	token, err := r.Cookie("token")
	if err != nil {
		log.Fatalf("Token nety")
	}

	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	profile := pb.Profile{TokenValue: token.Value}
	user, _ := client.ProfileUser(ctx, &profile)

	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/profile.html"))
	tmpl.ExecuteTemplate(w, "profile.html", user)
}
