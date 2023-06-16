package data

import (
	"case/pkg/proto"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
)

type CaseServer struct {
	proto.UnimplementedCaseServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func MakeAuthServer() *CaseServer {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	db, er := sql.Open("postgres", "postgres://almaz:24052004@localhost:5432/auth?sslmode=disable")
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}
	AuthSrv := &CaseServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	return AuthSrv
}

var (
	ErrCase = errors.New("case error")
)

func TestCaseModel_InsertItem(t *testing.T) {
	Srv := MakeAuthServer()
	tests := []struct {
		name          string
		item          *proto.Case
		expectedError error
	}{
		{
			name: "Insert new item",
			item: &proto.Case{
				Name:  "Test Item",
				Price: 10,
				CaseItems: []*proto.CaseItem{
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
			err := CaseModel{DB: Srv.DB}.InsertItem(tt.item)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
func TestCaseModel_GetCaseID(t *testing.T) {
	Srv := MakeAuthServer()
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
			expectedError: ErrRecordNotFound,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caseItem, err := CaseModel{DB: Srv.DB}.GetCaseID(tt.id)

			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
			fmt.Println(caseItem)
		})
	}
}
func TestGetAllCase(t *testing.T) {
	Srv := MakeAuthServer()
	caseItems, err := CaseModel{DB: Srv.DB}.GetAllCase()
	if err != nil {
		t.Errorf("unexpected error: got %v, expected %v", err, err)
	}
	if caseItems == nil {
		t.Errorf("cases are empty")
	}
}
func TestDeleteCase(t *testing.T) {
	Srv := MakeAuthServer()
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
			expectedError: ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CaseModel{DB: Srv.DB}.DeleteCase(tt.id)

			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
