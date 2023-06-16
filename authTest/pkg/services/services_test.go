package services

import (
	"authTest/internal/data"
	"authTest/pkg/proto"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
	"time"
)

func MakeAuthServer() *AuthServer {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	db, er := sql.Open("postgres", "postgres://diamondw:7PTGiP7WLhz5yHF7PRbkk7inqCLrjjP5@dpg-ci5tp218g3n4q9uac9u0-a.oregon-postgres.render.com/auth_rs5f")
	if er != nil {
		log.Fatalf("postgres doesnt work : %s", er)
	}
	AuthSrv := &AuthServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	return AuthSrv
}

func TestAuthServer_Register(t *testing.T) {
	Srv := MakeAuthServer()

	authServer := AuthServer{DB: Srv.DB}

	pass := proto.Password{PlainText: "password123",
		Hash: []byte("aaaaa")}
	// Create a test user
	user := &proto.User{
		Name:     "Test User",
		Email:    "taaest@example.com",
		Password: &pass,
	}

	// Register the test user
	confirm, err := authServer.Register(context.Background(), user)
	if err != nil {
		t.Errorf("error registering user: %v", err)
	}

	// Verify the registration confirmation
	if confirm == nil || !confirm.Ok {
		t.Error("user registration failed")
	}

	// Retrieve the registered user from the database
	registeredUser, err := data.UserModel{DB: Srv.DB}.GetByEmail(user.Email)
	if err != nil {
		t.Errorf("error retrieving registered user: %v", err)
	}

	// Verify the retrieved user matches the registered user
	if registeredUser == nil || registeredUser.Name != user.Name || registeredUser.Email != user.Email {
		t.Error("retrieved user does not match the registered user")
	}

	err = data.UserModel{DB: Srv.DB}.DeleteUser("Test User")
	if err != nil {
		t.Errorf("error deleting test user: %v", err)
	}
}

func TestAuthServer_Login(t *testing.T) {
	Srv := MakeAuthServer()

	authServer := AuthServer{DB: Srv.DB}

	pass := proto.Password{
		PlainText: "Almazalmaz1",
		Hash:      []byte("a"),
	}

	// Create a test user
	user := &proto.User{
		Name:     "Almazaa",
		Email:    "almazsydykov761@gmail.com",
		Password: &pass,
	}

	// Attempt to log in with the test user
	token, err := authServer.Login(context.Background(), user)
	if err != nil {
		t.Errorf("error logging in: %v", err)
	}
	fmt.Println(token)

	err = data.UserModel{DB: Srv.DB}.DeleteUser("Almaz")
	if err != nil {
		t.Errorf("error deleting test user: %v", err)
	}

}

func TestAuthServer_UpdateUser(t *testing.T) {
	Srv := MakeAuthServer()

	authServer := AuthServer{DB: Srv.DB}
	pass := proto.Password{
		PlainText: "Almazalmaz1",
		Hash:      []byte("a"),
	}
	// Create a test user
	user := &proto.User{
		Name:     "Tedstd User",
		Email:    "tddssgest@example.com",
		Password: &pass,
	}

	// Register the test user
	_, err := authServer.Register(context.Background(), user)
	if err != nil {
		t.Errorf("error registering user: %v", err)
	}

	// Create an update request
	update := &proto.Update{
		TokenValue: "your_token_value_here",
		Name:       "Updated Name",
	}

	// Update the user
	confirm, err := authServer.UpdateUser(context.Background(), update)
	if err != nil {
		t.Errorf("error updating user: %v", err)
	}

	// Verify the confirmation message
	if !confirm.Ok || confirm.Message != "Update was successfully" {
		t.Error("unexpected confirmation message")
	}
	err = data.UserModel{DB: Srv.DB}.DeleteUser("Tedstd User")
	if err != nil {
		t.Errorf("error deleting test user: %v", err)
	}
}

func TestAuthServer_ProfileUser(t *testing.T) {
	Srv := MakeAuthServer()

	authServer := AuthServer{DB: Srv.DB}

	pass := proto.Password{
		PlainText: "Almazalmaz1",
		Hash:      []byte("a"),
	}
	// Create a test user
	user := &proto.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: &pass,
	}

	// Register the test user
	_, err := authServer.Register(context.Background(), user)
	if err != nil {
		t.Errorf("error registering user: %v", err)
	}

	// Generate a token for the test user
	token, err := data.TokenModel.New(data.TokenModel{DB: Srv.DB}, 1, time.Hour, data.ScopeAuthentication)
	if err != nil {
		t.Errorf("error generating token: %v", err)
	}

	// Create a profile request with the generated token
	profile := &proto.Profile{
		TokenValue: token.PlainText,
	}

	// Get the user profile
	result, err := authServer.ProfileUser(context.Background(), profile)
	if err != nil {
		t.Errorf("error getting user profile: %v", err)
	}

	// Verify the user profile
	if result == nil || result.Name != "Test User" || result.Email != "test@example.com" {
		t.Error("unexpected user profile")
	}
}
