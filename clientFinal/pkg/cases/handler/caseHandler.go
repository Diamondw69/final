package handlercase

import (
	pb "clientFinal/pkg/cases/proto"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	protoB "google.golang.org/protobuf/proto"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func ConnectGrpc() *grpc.ClientConn {

	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}

func CreateCaseHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewCaseServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"caseItems", // name
		"fanout",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)

	// Declare a queue
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return
	}

	err = ch.QueueBind(
		q.Name,      // queue name
		"",          // routing key
		"caseItems", // exchange
		false,
		nil,
	)
	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer name (empty generates a unique name)
		true,   // Auto-acknowledgment
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		return
	}
	// Start consuming messages

	caseitems := make([]*pb.CaseItem, 0)
	for i := 1; i <= 5; i++ {
		itemName := r.FormValue(fmt.Sprintf("star%d", i))
		if itemName != "" {
			caseReq := pb.CaseItemRequest{
				Name: itemName,
			}
			item, err := client.GetCaseItem(ctx, &caseReq)
			if item.Ok {
				for msg := range msgs {
					caseItems := &pb.CaseItem{}
					err := protoB.Unmarshal(msg.Body, caseItems)
					if err != nil {
						return
					}
					caseitems = append(caseitems, caseItems)
					break
				}
			}
			if err != nil {
				fmt.Println("Problem v create case handler")
				return
			}

		}
	}
	// Parse the price from the form data
	price, err := strconv.ParseInt(r.Form.Get("price"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	// Create the case item
	item := &pb.Case{
		Name:      r.Form.Get("name"),
		Price:     price,
		CaseItems: caseitems,
	}

	confirm, err := client.CreateCase(ctx, item)
	if err != nil {
		fmt.Println("Problem with grpc func")
	}

	if confirm.Ok {
		http.Redirect(w, r, "/", 303)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func DeleteCaseHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewCaseServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := pb.CaseRequest{Id: num}

	case1, _ := client.DeleteCase(ctx, &req)
	fmt.Println(case1.Message)

	http.Redirect(w, r, "/cases", 303)
}
