package main

import (
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"inventory/configs/Connection"
	"inventory/pkg/serviceI"
)

func main() {
	//Connect to DB, RabbitMQ and grpc
	rabbitMq := Connection.ConnectToRabbitMQ("RABBITMQ")
	Db := Connection.DbConnection()

	defer Db.Close()
	defer rabbitMq.Close()

	//Start a grpc server
	serviceI.NewGrpcServer(rabbitMq, Db)

}
