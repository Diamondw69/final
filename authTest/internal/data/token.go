package data

import (
	pb "authTest/pkg/proto"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication" // Include a new authentication scope.
)

func generateToken(userID int64, ttl time.Duration, scope string) (*pb.Token, error) {

	token := &pb.Token{
		Id:     userID,
		Expiry: timestamppb.New(time.Now()),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]
	return token, nil
}

type TokenModel struct {
	DB *sql.DB
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*pb.Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}
	err = m.Insert(token)
	if err != nil {
		fmt.Println(err)
	}
	return token, err
}

func (m TokenModel) Insert(token *pb.Token) error {
	query := `
INSERT INTO tokens (hash, user_id, expiry, scope)
VALUES ($1, $2, $3, $4)`

	t, err := ptypes.Timestamp(token.Expiry)
	args := []interface{}{token.Hash, token.Id, t, token.Scope}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, args...)
	return err
}
