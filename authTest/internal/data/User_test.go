package data

import (
	"authTest/pkg/proto"
	pb "authTest/pkg/proto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

type AuthServer struct {
	proto.UnimplementedUserServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

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

func TestMatches(t *testing.T) {
	pass1 := "almazalmaz1"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass1), 12)

	tests := []struct {
		plaintextPassword string
		hash              []byte
		expected          bool
	}{
		{
			plaintextPassword: pass1,
			hash:              hash,
			expected:          true,
		},
		{
			plaintextPassword: "somePassword",
			hash:              []byte("I dont like making tests"),
			expected:          false,
		},
	}

	for _, tt := range tests {
		t.Run("Maches test", func(t *testing.T) {
			result, err := Matches(tt.plaintextPassword, tt.hash)
			if err != nil {
				if result != tt.expected {
					t.Errorf("Unexpected error. Expected %v but got %v", tt.expected, err)
				}
				return
			}
			if result != tt.expected {
				t.Errorf("Result doesnt satisfy expected. Expected %v but got %v", tt.expected, result)
			}
		})
	}
}

func TestUserModel_GetByEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		expectedError error
	}{
		{
			name:          "Get existing user by email",
			email:         "almazsydykov768@gmail.com",
			expectedError: nil,
		},
		{
			name:          "Get non-existing user by email",
			email:         "none@example.com",
			expectedError: ErrRecordNotFound,
		},
	}
	Srv := MakeAuthServer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := UserModel{DB: Srv.DB}.GetByEmail(tt.email)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
			if tt.expectedError == nil {
				if user == nil {
					t.Error("user is nil but no error was returned")
				} else {
					fmt.Println(user.Id)
				}
			}
		})
	}
}

func TestUserModel_Insert(t *testing.T) {
	Srv := MakeAuthServer()

	tests := []struct {
		name          string
		user          *pb.User
		expectedError error
	}{
		{
			name: "Insert new user",
			user: &pb.User{
				Name:  "Test User",
				Email: "test@example.com",
				Password: &pb.Password{
					PlainText: "password123",
					Hash:      []byte("hashvalue"),
				},
				Role:    "user",
				Balance: 0,
			},
			expectedError: nil,
		},
		{
			name: "Insert duplicate user",
			user: &pb.User{
				Name:  "Duplicate User",
				Email: "test@example.com", // Same email as the first test case
				Password: &pb.Password{
					PlainText: "password456",
					Hash:      []byte("hashvalue"),
				},
				Role:    "user",
				Balance: 0,
			},
			expectedError: ErrDuplicateEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UserModel{DB: Srv.DB}.Insert(tt.user)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}

		})

	}
	err := UserModel{DB: Srv.DB}.DeleteUser("Test User")
	if err != nil {
		t.Errorf("error deleting test user: %v", err)
	}
}

func TestUserModel_DeleteUser(t *testing.T) {
	Srv := MakeAuthServer()

	tests := []struct {
		name          string
		userName      string
		expectedError error
	}{
		{
			name:          "Delete existing user",
			userName:      "Admin",
			expectedError: nil,
		},
		{
			name:          "Delete non-existing user",
			userName:      "Non-existing User",
			expectedError: ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UserModel{DB: Srv.DB}.DeleteUser(tt.userName)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
	user := &pb.User{
		Name:  "Admin",
		Email: "xatest@example.com",
		Password: &pb.Password{
			PlainText: "password123",
			Hash:      []byte("hashvalue"),
		},
		Role:    "user",
		Balance: 0,
	}
	err := UserModel{DB: Srv.DB}.Insert(user)
	if err != nil {
		t.Errorf("error inserting test user: %v", err)
	}
}

func TestUserModel_GetForToken(t *testing.T) {
	Srv := MakeAuthServer()

	tests := []struct {
		name           string
		tokenScope     string
		tokenPlaintext string
		expectedError  error
	}{
		{
			name:           "Get user for valid token",
			tokenScope:     "scope",
			tokenPlaintext: "password123",
			expectedError:  ErrRecordNotFound,
		},
		{
			name:           "Get user for invalid token",
			tokenScope:     "scope",
			tokenPlaintext: "invalidpassword",
			expectedError:  ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := UserModel{DB: Srv.DB}.GetForToken(tt.tokenScope, tt.tokenPlaintext)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
			if tt.expectedError == nil && user == nil {
				t.Error("user is nil but no error was returned")
			} else {
				fmt.Println(user.Id)
			}
		})
	}
}

func TestUserModel_Update(t *testing.T) {
	Srv := MakeAuthServer()

	tests := []struct {
		name          string
		user          *pb.User
		expectedError error
	}{
		{
			name: "Update existing user",
			user: &pb.User{
				Id:       1,
				Name:     "Updated User",
				Email:    "updated@example.com",
				Password: &pb.Password{Hash: []byte("updatedhash")},
				Role:     "updated role",
				Balance:  100.0,
			},
			expectedError: nil,
		},
		{
			name: "Update non-existing user",
			user: &pb.User{
				Id:       10,
				Name:     "Non-existing User",
				Email:    "nonexisting@example.com",
				Password: &pb.Password{Hash: []byte("nonexistinghash")},
				Role:     "non-existing role",
				Balance:  200.0,
			},
			expectedError: ErrEditConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UserModel{DB: Srv.DB}.Update(tt.user)
			if err != tt.expectedError {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedError)
			}
		})
	}
}
