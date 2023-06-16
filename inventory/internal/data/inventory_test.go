package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"inventory/pkg/proto"
	"log"
	"testing"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func MakeAuthServer() *InventoryServer {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	db, er := sql.Open("postgres", "postgres://almaz:24052004@localhost:5432/auth?sslmode=disable")
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}
	AuthSrv := &InventoryServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	return AuthSrv
}
func TestGenerateInventory(t *testing.T) {
	tests := []struct {
		name          string
		id            int64
		expectedError error
	}{
		{
			name:          "Invalid case ID",
			id:            4,
			expectedError: nil,
		},
		{
			name:          "Invalid case ID",
			id:            -1,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generateInventory(tt.id)

			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
func TestNewInventory(t *testing.T) {
	Srv := MakeAuthServer()
	tests := []struct {
		name          string
		id            int64
		expectedError error
	}{
		{
			name:          "Invalid case ID",
			id:            56,
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InventoryModel{DB: Srv.DB}.New(tt.id)

			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
func TestInventory_Insert(t *testing.T) {
	Srv := MakeAuthServer()
	tests := []struct {
		name          string
		inventory     *proto.Inventory
		expectedError error
	}{
		{
			name: "Insert new inventory",
			inventory: &proto.Inventory{
				Id:     3,
				Userid: 1,
				Items: []*proto.CaseItem{
					{
						Id:              1,
						ItemName:        "Item 1",
						ItemDescription: "Description 1",
						Type:            "Type 1",
						Stars:           5,
						Image:           []byte("image data"),
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InventoryModel{DB: Srv.DB}.Insert(tt.inventory)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
func TestCheckInventory(t *testing.T) {
	Srv := MakeAuthServer()
	tests := []struct {
		name string
		id   int64
	}{
		{
			name: "valid case ID",
			id:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ture := InventoryModel{DB: Srv.DB}.Check(tt.id)
			if ture {
				fmt.Println("has inventory")
			} else {
				fmt.Println("hasn't inventory")
			}
		})
	}
}
func TestGetByIdInventory(t *testing.T) {
	Srv := MakeAuthServer()
	tests := []struct {
		name string
		id   int64
	}{
		{
			name: "valid case ID",
			id:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cheto, err := InventoryModel{DB: Srv.DB}.GetByID(tt.id)
			if err != nil {
				t.Errorf("unexpected error: got %v", err)
			}
			if cheto == nil {
				t.Errorf("oh no")
			}
		})
	}
}
