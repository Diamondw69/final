package serviceI

import (
	"context"
	"database/sql"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	protoB "google.golang.org/protobuf/proto"
	"inventory/configs/Connection"
	"inventory/internal/data"
	"inventory/pkg/proto"
	_case "inventory/pkg/proto/case"
	"inventory/pkg/proto/user"
	"log"
	"net"
	"os"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func NewGrpcServer(rabbitMQConn *amqp.Connection, db *sql.DB) {

	//Getting data from env
	Connection.EnvLoader("./configs/env/.env")

	grpcServer := grpc.NewServer()

	//Making new auth server and starting server
	AuthSrv := &InventoryServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	proto.RegisterInventoryServiceServer(grpcServer, AuthSrv)

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

func (a *InventoryServer) ToInventory(ctx context.Context, req *proto.InventoryRequest) (*proto.Confirm, error) {
	log.Printf("Adding CaseItem to Inventory by id:%d", req.Id)

	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}
	conf := GetCaseItemById(ctx, req.Id)

	caseItem := &proto.CaseItem{}

	if conf.Ok {
		for msg := range msgs {
			err := protoB.Unmarshal(msg.Body, caseItem)
			if err != nil {
				return nil, err
			}
			break
		}
	}
	userid := GetUserFromToken(ctx, req.TokenValue)

	inventory, _ := data.InventoryModel.GetByID(data.InventoryModel{DB: a.DB}, userid)
	inventory.Items = append(inventory.Items, caseItem)
	err = data.InventoryModel.Update(data.InventoryModel{DB: a.DB}, inventory)

	return conf, nil
}

func (a *InventoryServer) NewInventory(ctx context.Context, req *proto.InventoryRequest) (*proto.Confirm, error) {
	log.Printf("Creating new inventory for a user by id:%d", req.Id)

	err := data.InventoryModel.New(data.InventoryModel{DB: a.DB}, req.Id)
	if err != nil {
		return nil, err
	}
	conf := proto.Confirm{
		Ok:      true,
		Message: "Inventory was created",
	}
	return &conf, nil
}

func (a *InventoryServer) GetInventory(ctx context.Context, req *proto.InventoryRequest) (*proto.Inventory, error) {

	log.Printf("Getting inventory by id:%d", req.Id)

	check := data.InventoryModel.Check(data.InventoryModel{DB: a.DB}, req.Id)

	if !check {
		return nil, nil
	}

	inventory, err := data.InventoryModel.GetByID(data.InventoryModel{DB: a.DB}, req.Id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func GetCaseItemById(ctx context.Context, id int64) *proto.Confirm {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := _case.NewCaseServiceClient(conn)

	req := _case.CaseItemRequest{
		Id:   id,
		Name: "",
	}

	item, _ := client.GetCaseItem(ctx, &req)

	conf := proto.Confirm{
		Ok:      item.Ok,
		Message: item.Message,
	}

	return &conf
}

func GetUserFromToken(ctx context.Context, tokenValue string) int64 {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := user.NewUserServiceClient(conn)

	prof := user.Profile{TokenValue: tokenValue}

	user1, _ := client.ProfileUser(ctx, &prof)

	return user1.Id
}
