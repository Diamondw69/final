package service

import (
	pb "ahah/pkg/proto"
	"context"
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
)

func MakeAuthServer() *CaseItemServer {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	db, er := sql.Open("postgres", "postgres://diamondw:7PTGiP7WLhz5yHF7PRbkk7inqCLrjjP5@dpg-ci5tp218g3n4q9uac9u0-a.oregon-postgres.render.com/auth_rs5f")
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}
	AuthSrv := &CaseItemServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	return AuthSrv
}

func TestCaseItemServer_CreateCaseItem(t *testing.T) {
	serv := MakeAuthServer()

	// Create a CaseItemServer with the test database
	caseItemServer := CaseItemServer{DB: serv.DB}

	// Create a test CaseItem
	caseItem := &pb.CaseItem{
		ItemName:        "Test Item",
		ItemDescription: "Test Description",
		Type:            "Test Type",
		Stars:           5,
		Image:           []byte("test.jpg"),
	}

	// Call CreateCaseItem
	confirm, err := caseItemServer.CreateCaseItem(context.Background(), caseItem)
	if err != nil {
		t.Errorf("error creating CaseItem: %v", err)
	}

	// Verify the confirmation
	if !confirm.Ok {
		t.Errorf("unexpected confirmation status: got false, expected true")
	}

	// Verify the confirmation message
	expectedMessage := "CaseItem was created"
	if confirm.Message != expectedMessage {
		t.Errorf("unexpected confirmation message: got %s, expected %s", confirm.Message, expectedMessage)
	}

}

func TestCaseItemServer_DeleteCaseItem(t *testing.T) {
	serv := MakeAuthServer()

	// Create a CaseItemServer with the test database
	caseItemServer := CaseItemServer{DB: serv.DB}

	// Create a test CaseItem to delete
	caseItem := &pb.CaseItem{
		Id: 2,
	}

	// Call DeleteCaseItem
	confirm, err := caseItemServer.DeleteCaseItem(context.Background(), caseItem)
	if err != nil {
		t.Errorf("error deleting CaseItem: %v", err)
	}

	// Verify the confirmation
	if !confirm.Ok {
		t.Errorf("unexpected confirmation status: got false, expected true")
	}

	// Verify the confirmation message
	expectedMessage := "CaseItem was successfully deleted"
	if confirm.Message != expectedMessage {
		t.Errorf("unexpected confirmation message: got %s, expected %s", confirm.Message, expectedMessage)
	}
}

func TestCaseItemServer_ShowCaseItem(t *testing.T) {
	serv := MakeAuthServer()

	// Create a CaseItemServer with the test database
	caseItemServer := CaseItemServer{DB: serv.DB}

	// Create a test CaseItem request by ID
	reqByID := &pb.CaseItemRequest{
		Id: 3,
	}

	// Call ShowCaseItem by ID
	caseItemByID, err := caseItemServer.ShowCaseItem(context.Background(), reqByID)
	if err != nil {
		t.Errorf("error getting CaseItem by ID: %v", err)
	}

	// Verify the CaseItem by ID
	expectedItemName := "Test Item"
	if caseItemByID.ItemName != expectedItemName {
		t.Errorf("unexpected CaseItem name by ID: got %s, expected %s", caseItemByID.ItemName, expectedItemName)
	}

	// Create a test CaseItem request by Name
	reqByName := &pb.CaseItemRequest{
		Name: "Test Item",
	}

	// Call ShowCaseItem by Name
	caseItemByName, err := caseItemServer.ShowCaseItem(context.Background(), reqByName)
	if err != nil {
		t.Errorf("error getting CaseItem by Name: %v", err)
	}

	// Verify the CaseItem by Name
	if caseItemByName.ItemName != expectedItemName {
		t.Errorf("unexpected CaseItem name by Name: got %s, expected %s", caseItemByName.ItemName, expectedItemName)
	}
}

func TestCaseItemServer_GetAllCaseItems(t *testing.T) {
	serv := MakeAuthServer()

	// Create a CaseItemServer with the test database
	caseItemServer := CaseItemServer{DB: serv.DB}

	// Call GetAllCaseItems
	caseItems, err := caseItemServer.GetAllCaseItems(context.Background(), &pb.Confirm{})
	if err != nil {
		t.Errorf("error getting all CaseItems: %v", err)
	}
	t.Log(caseItems)

}
