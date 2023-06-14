package handlercase

import (
	pb "clientFinal/pkg/caseItems/proto"
	"clientFinal/pkg/cases/proto"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateCaseHTMLHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf := pb.Confirm{Ok: true}
	caseItems, _ := client.GetAllCaseItems(ctx, &conf)
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/caseInput.html"))
	tmpl.ExecuteTemplate(w, "caseInput.html", caseItems.CaseItems)
}

func ViewCaseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	conn := ConnectGrpc()
	defer conn.Close()

	client := proto.NewCaseServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := proto.CaseRequest{Id: num}

	case1, err := client.ViewCase(ctx, &req)

	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/case.html"))
	tmpl.ExecuteTemplate(w, "case.html", case1)
}

func ViewAllCaseHandler(w http.ResponseWriter, r *http.Request) {
	conn := ConnectGrpc()
	defer conn.Close()

	client := proto.NewCaseServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := proto.Confirm{Ok: true}

	case1, _ := client.ShowAllCases(ctx, &req)

	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/allcases.html"))
	tmpl.ExecuteTemplate(w, "allcases.html", case1.Cases)
}
