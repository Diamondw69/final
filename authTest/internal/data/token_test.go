package data

import (
	pb "authTest/pkg/proto"
	"bytes"
	"crypto/sha256"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	userID := int64(1)
	ttl := time.Hour
	scope := "testScope"

	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		t.Errorf("error generating token: %v", err)
	}

	// Verify the generated token properties
	if token.Id != userID {
		t.Errorf("unexpected token ID: got %d, expected %d", token.Id, userID)
	}

	if token.Scope != scope {
		t.Errorf("unexpected token scope: got %s, expected %s", token.Scope, scope)
	}

	// Verify the token hash is generated correctly
	expectedHash := sha256.Sum256([]byte(token.PlainText))
	if !bytes.Equal(token.Hash, expectedHash[:]) {
		t.Errorf("unexpected token hash: got %v, expected %v", token.Hash, expectedHash[:])
	}
}

func TestTokenModel_New(t *testing.T) {
	Srv := MakeAuthServer()

	tokenModel := TokenModel{DB: Srv.DB}
	userID := int64(1)
	ttl := time.Hour
	scope := "testScope"

	token, err := tokenModel.New(userID, ttl, scope)
	if err != nil {
		t.Errorf("error creating new token: %v", err)
	}

	// Verify the token is inserted successfully
	if token == nil {
		t.Error("token is nil, expected a valid token")
	}

	// Verify the inserted token properties
	if token.Id == 0 {
		t.Error("unexpected token ID: got 0, expected a non-zero value")
	}

	if token.Scope != scope {
		t.Errorf("unexpected token scope: got %s, expected %s", token.Scope, scope)
	}

}

func TestTokenModel_Insert(t *testing.T) {
	Srv := MakeAuthServer()

	tokenModel := TokenModel{DB: Srv.DB}

	// Create a test token
	token := &pb.Token{
		Id:     1,
		Expiry: timestamppb.New(time.Now()),
		Scope:  "testScope",
	}

	// Insert the token into the database
	err := tokenModel.Insert(token)
	if err != nil {
		t.Errorf("error inserting token: %v", err)
	}
}
