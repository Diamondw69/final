package serviceC

import (
	"case/configs/Conn"
	"case/internal/data"
	"case/pkg/proto"
	protoItem "case/pkg/proto1"
	"context"
	"database/sql"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	rb "google.golang.org/protobuf/proto"
	"log"
	"net"
	"os"
	"time"
)

type CaseServer struct {
	proto.UnimplementedCaseServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func NewGrpcServer(rabbitMQConn *amqp.Connection, db *sql.DB) {

	//Getting data from env
	Conn.EnvLoader("./configs/env/.env")

	grpcServer := grpc.NewServer()

	//Making new auth server and starting server
	AuthSrv := &CaseServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	proto.RegisterCaseServiceServer(grpcServer, AuthSrv)

	TcpPort := os.Getenv("TCP_PORT")

	listener, err := net.Listen("tcp", TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server started on port " + TcpPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *CaseServer) CreateCase(ctx context.Context, case1 *proto.Case) (*proto.Confirm, error) {
	log.Printf("Creating a case:%s", case1.Name)
	err := data.CaseModel.InsertItem(data.CaseModel{DB: a.DB}, case1)
	fmt.Println(err)
	if err != nil {
		return &proto.Confirm{
			Ok:      false,
			Message: "CaseItem was not created",
		}, err
	}

	return &proto.Confirm{
		Ok:      true,
		Message: "CaseItem was created",
	}, nil
}

func (a *CaseServer) ViewCase(ctx context.Context, req *proto.CaseRequest) (*proto.Case, error) {
	log.Printf("Showing a case by Id:%d", req.Id)

	case1, err := data.CaseModel.GetCaseID(data.CaseModel{DB: a.DB}, req.Id)
	if err != nil {
		return nil, err
	}
	return case1, nil
}

func (a *CaseServer) DeleteCase(ctx context.Context, req *proto.CaseRequest) (*proto.Confirm, error) {
	log.Printf("Deleting a case by id:%d", req.Id)

	confirm := proto.Confirm{
		Ok:      false,
		Message: "Case was not deleted",
	}

	err := data.CaseModel.DeleteCase(data.CaseModel{DB: a.DB}, req.Id)
	if err != nil {
		return &confirm, err
	}

	confirm.Ok = true
	confirm.Message = "Case was successfully deleted"

	return &confirm, nil
}

func (a *CaseServer) ShowAllCases(ctx context.Context, conf *proto.Confirm) (*proto.Cases, error) {
	log.Printf("Showing all cases")

	cases, err := data.CaseModel.GetAllCase(data.CaseModel{DB: a.DB})
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func (a *CaseServer) GetCaseItem(ctx context.Context, req *proto.CaseItemRequest) (*proto.Confirm, error) {
	log.Printf("Getting CaseItem by name:%s", req.Name)

	ch, err := a.rabbitMQConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := protoItem.NewUserServiceClient(conn)

	request := protoItem.CaseItemRequest{
		Id:   req.Id,
		Name: req.Name,
	}
	caseItem, _ := client.ShowCaseItem(ctx, &request)

	// Publish a message to the exchange
	message, _ := rb.Marshal(caseItem)
	err = ch.PublishWithContext(ctx,
		"caseItems", // exchange
		"",          // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	conf := proto.Confirm{Ok: true}
	return &conf, nil
}
