package handlerCaseItem

import (
	pb "clientFinal/pkg/caseItems/proto"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CreateCaseItemHTMLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/createCaseItem.html"))
	tmpl.ExecuteTemplate(w, "createCaseItem.html", nil)
}

func GetAllCaseItemsHTMLHandler(w http.ResponseWriter, r *http.Request) {
	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := pb.Confirm{
		Ok:      false,
		Message: "",
	}

	items, err := client.GetAllCaseItems(ctx, &c)
	if err != nil {
		fmt.Println("Problem v html handler")
	}

	tmpl = template.Must(template.ParseFiles("cmd/caseApp/static/templates/listitems.html"))
	tmpl.ExecuteTemplate(w, "listitems.html", items.CaseItems)
}
func ImageHandler(w http.ResponseWriter, r *http.Request) {

	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Get the ID of the image from the request URL
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	req := pb.CaseItemRequest{
		Id:   num,
		Name: "",
	}
	var data1 []byte
	x, _ := client.ShowCaseItem(ctx, &req)
	data1 = x.Image
	w.Header().Set("Content-Type", "image/png") // Or "image/png" if the image is in PNG format

	// Write the image data to the HTTP response body
	_, err = w.Write(data1)
	if err != nil {
		log.Fatal(err)
	}
}
