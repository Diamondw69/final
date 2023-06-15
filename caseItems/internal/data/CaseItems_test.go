package data

import (
	pb "ahah/pkg/proto"
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
)

type CaseItemServer struct {
	pb.UnimplementedUserServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func MakeAuthServer() *CaseItemServer {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	db, er := sql.Open("postgres", "postgres://postgres:24052004@localhost:5432/auth?sslmode=disable")
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}
	AuthSrv := &CaseItemServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	return AuthSrv
}

func TestCaseItemModel_InsertItem(t *testing.T) {

	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Create a test CaseItem
	item := &pb.CaseItem{
		ItemName:        "Test Item",
		ItemDescription: "Test Description",
		Type:            "Test Type",
		Stars:           5,
		Image:           []byte("test.jpg"),
	}

	// Insert the test CaseItem
	err := caseItemModel.InsertItem(item)
	if err != nil {
		t.Errorf("error inserting CaseItem: %v", err)
	}

	// Verify that the CaseItem has been assigned an ID
	if item.Id == 0 {
		t.Error("expected CaseItem ID to be assigned, but it is 0")
	}
}

func TestCaseItemModel_GetCaseItem(t *testing.T) {
	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Get the CaseItem by ID
	item, err := caseItemModel.GetCaseItem(1)
	if err != nil {
		t.Errorf("error getting CaseItem: %v", err)
	}

	t.Log(item)
}

func TestCaseItemModel_DeleteCaseItems(t *testing.T) {
	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Delete the CaseItem by ID
	err := caseItemModel.DeleteCaseItems(1)
	if err != nil {
		t.Errorf("error deleting CaseItem: %v", err)
	}

	// Attempt to get the deleted CaseItem by ID
	_, err = caseItemModel.GetCaseItem(1)
	if !errors.Is(err, ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestCaseItemModel_UpdateItem(t *testing.T) {
	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Create a modified CaseItem
	modifiedItem := &pb.CaseItem{
		Id:              2,
		ItemName:        "Modified Item",
		ItemDescription: "Modified Description",
		Type:            "Modified Type",
		Stars:           4,
		Image:           []byte("modified.jpg"),
	}

	// Update the CaseItem
	err := caseItemModel.UpdateItem(modifiedItem)
	if err != nil {
		t.Errorf("error updating CaseItem: %v", err)
	}

	// Get the updated CaseItem by ID
	item, err := caseItemModel.GetCaseItem(2)
	if err != nil {
		t.Errorf("error getting CaseItem: %v", err)
	}

	// Verify the updated CaseItem properties
	if item.ItemName != "Modified Item" {
		t.Errorf("unexpected CaseItem name: got %s, expected %s", item.ItemName, "Modified Item")
	}
}

func TestCaseItemModel_GetCaseItemByName(t *testing.T) {
	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Get the CaseItem by name
	item, err := caseItemModel.GetCaseItemByName("Test Item")
	if err != nil {
		t.Errorf("error getting CaseItem: %v", err)
	}

	// Verify the CaseItem properties
	if item.ItemName != "Test Item" {
		t.Errorf("unexpected CaseItem name: got %s, expected %s", item.ItemName, "Test Item")
	}
}

func TestCaseItemModel_GetAllCaseItems(t *testing.T) {
	serv := MakeAuthServer()
	// Create a CaseItemModel with the test database
	caseItemModel := CaseItemModel{DB: serv.DB}

	// Get all CaseItems
	items, err := caseItemModel.GetAllCaseItems()
	if err != nil {
		t.Errorf("error getting all CaseItems: %v", err)
	}

	log.Println(items)

}
