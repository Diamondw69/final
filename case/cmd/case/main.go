package main

import (
	"case/configs/Conn"
	"case/pkg/serviceC"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	//Connect to DB, RabbitMQ and grpc
	rabbitMq := Conn.ConnectToRabbitMQ("RABBITMQ")
	Db := Conn.DbConnection()

	defer Db.Close()
	defer rabbitMq.Close()

	//Start a grpc server
	serviceC.NewGrpcServer(rabbitMq, Db)

}
