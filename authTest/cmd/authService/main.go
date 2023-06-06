package main

import (
	"authTest/configs/Connections"
	"authTest/pkg/services"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	//Connect to DB, RabbitMQ and grpc
	rabbitMq := Connections.ConnectToRabbitMQ("RABBITMQ")
	Db := Connections.DbConnection()

	defer Db.Close()
	defer rabbitMq.Close()

	//Start a grpc server
	services.NewGrpcServer(rabbitMq, Db)

}
