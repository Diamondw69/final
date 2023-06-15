package handlers

import (
	pb "clientFinal/pkg/auth/proto"
	inventory "clientFinal/pkg/inventory/proto"
	"context"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
)

type Profiles struct {
	ID      int64                 `json:"id"`
	Name    string                `json:"name"`
	Email   string                `json:"email"`
	Role    string                `json:"role"`
	Balance int64                 `json:"balance"`
	Items   []*inventory.CaseItem `json:"items"`
}

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

	conn2, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn2.Close()

	client2 := inventory.NewInventoryServiceClient(conn2)
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	profile := pb.Profile{TokenValue: token.Value}
	user, _ := client.ProfileUser(ctx, &profile)

	req := inventory.InventoryRequest{
		TokenValue: token.Value,
		Id:         user.Id,
	}
	inventory1, _ := client2.GetInventory(ctx, &req)

	profile1 := Profiles{
		ID:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Role:    user.Role,
		Balance: user.Balance,
		Items:   inventory1.Items,
	}
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/profile.html"))
	tmpl.ExecuteTemplate(w, "profile.html", profile1)
}
