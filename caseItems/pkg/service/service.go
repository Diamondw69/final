package service

import (
	"ahah/configs/connections"
	"ahah/internal/data"
	pb "ahah/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type CaseItemServer struct {
	pb.UnimplementedUserServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func NewGrpcServer(rabbitMQConn *amqp.Connection, db *sql.DB) {

	//Getting data from env
	Connections.EnvLoader("./configs/env/.env")

	grpcServer := grpc.NewServer()

	//Making new auth server and starting server
	AuthSrv := &CaseItemServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	pb.RegisterUserServiceServer(grpcServer, AuthSrv)

	TcpPort := os.Getenv("TCP_PORT")

	listener, err := net.Listen("tcp", TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server started on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *CaseItemServer) CreateCaseItem(ctx context.Context, caseItem *pb.CaseItem) (*pb.Confirm, error) {
	log.Printf("Creating a caseItem :%s", caseItem.ItemName)
	err := data.CaseItemModel.InsertItem(data.CaseItemModel{DB: a.DB}, caseItem)
	fmt.Println(err)
	if err != nil {
		return &pb.Confirm{
			Ok:      false,
			Message: "CaseItem was not created",
		}, err
	}

	return &pb.Confirm{
		Ok:      true,
		Message: "CaseItem was created",
	}, nil
}

func (a *CaseItemServer) DeleteCaseItem(ctx context.Context, caseItem *pb.CaseItem) (*pb.Confirm, error) {
	log.Printf("Deleting Case Item :%d", caseItem.Id)

	confirm := pb.Confirm{
		Ok:      false,
		Message: "CaseItem was not deleted",
	}
	err := data.CaseItemModel.DeleteCaseItems(data.CaseItemModel{DB: a.DB}, caseItem.Id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return &confirm, err
		default:
			return &confirm, err
		}
	}

	confirm.Ok = true
	confirm.Message = "CaseItem was successfully deleted"

	return &confirm, nil
}

func (a *CaseItemServer) ShowCaseItem(ctx context.Context, req *pb.CaseItemRequest) (*pb.CaseItem, error) {
	log.Printf("Getting CaseItem by Id :%d or name:%s", req.Id, req.Name)

	if req.Id != 0 {
		item, err := data.CaseItemModel.GetCaseItem(data.CaseItemModel{DB: a.DB}, req.Id)
		if err != nil {
			return nil, err
		}
		return item, nil
	} else if req.Name != "" {
		item, err := data.CaseItemModel.GetCaseItemByName(data.CaseItemModel{DB: a.DB}, req.Name)
		if err != nil {
			return nil, err
		}
		return item, nil
	}
	return nil, nil
}

func (a *CaseItemServer) GetAllCaseItems(ctx context.Context, confirm *pb.Confirm) (*pb.CaseItems, error) {
	log.Printf("Getting all caseItems")

	caseItems, err := data.CaseItemModel.GetAllCaseItems(data.CaseItemModel{DB: a.DB})
	if err != nil {
		return nil, err
	}
	return caseItems, nil
}
