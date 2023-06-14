package handlerCaseItem

import (
	pb "clientFinal/pkg/caseItems/proto"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func ConnectGrpc() *grpc.ClientConn {

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}

func CreateCaseItemHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data1, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itemname := r.Form.Get("itemname")
	itemdesc := r.Form.Get("itemdescription")
	typee := r.Form.Get("type")
	stars := r.Form.Get("stars")
	num, err := strconv.ParseInt(stars, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	item := &pb.CaseItem{
		ItemName:        itemname,
		ItemDescription: itemdesc,
		Type:            typee,
		Stars:           num,
		Image:           data1,
	}

	//Connecting to grpc
	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	confirm, err := client.CreateCaseItem(ctx, item)
	if err != nil {
		fmt.Println("Cannot create durak")
	}

	if confirm.Ok {
		fmt.Println("vse harasho")
	} else {
		fmt.Println("vse ploho")
	}
}

func DeleteCaseItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//Connecting to grpc
	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	caseItem := pb.CaseItem{Id: num}

	confirm, err := client.DeleteCaseItem(ctx, &caseItem)
	if err != nil {
		fmt.Println("Problem in delete handler")
	}

	if confirm.Ok {
		fmt.Println("Vse ok")
	} else {
		fmt.Println("Vse plohaa")
	}

	http.Redirect(w, r, "/listcaseitems", 303)
}
