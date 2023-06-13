package main

import (
	"ahah/configs/connections"
	"ahah/pkg/service"
	_ "github.com/lib/pq"
)

func main() {
	//Connect to DB, RabbitMQ and grpc
	rabbitMq := Connections.ConnectToRabbitMQ("RABBITMQ")
	Db := Connections.DbConnection()

	defer Db.Close()
	defer rabbitMq.Close()

	//Start a grpc server
	service.NewGrpcServer(rabbitMq, Db)

}
