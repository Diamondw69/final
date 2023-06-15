package handlerI

import (
	protoI "clientFinal/pkg/inventory/proto"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func ConnectGrpc() *grpc.ClientConn {

	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}

func ToInventoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	token, _ := r.Cookie("token")

	conn := ConnectGrpc()

	client := protoI.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := protoI.InventoryRequest{
		TokenValue: token.Value,
		Id:         num,
	}

	conf, err := client.ToInventory(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf.Message)

	http.Redirect(w, r, "/", 303)
}
