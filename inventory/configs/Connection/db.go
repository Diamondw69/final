package Connection

import (
	"database/sql"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var Path string = "./configs/env/.env"

func DbConnection() *sql.DB {

	EnvLoader(Path)

	psqlInfo := os.Getenv("POSTGRESQL")

	db, er := sql.Open("postgres", psqlInfo)
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}

	return db
}

func EnvLoader(pathToEnv string) {
	err := godotenv.Load(pathToEnv)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func ConnectToRabbitMQ(rabbitMq string) *amqp.Connection {

	EnvLoader(Path)

	RabbitMq := os.Getenv(rabbitMq)

	rabbitMQConn, err := amqp.Dial(RabbitMq)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	return rabbitMQConn
}
