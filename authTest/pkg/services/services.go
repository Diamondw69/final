package services

import (
	"authTest/configs/Connections"
	"authTest/internal/data"
	"authTest/internal/validator"
	"authTest/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"os"
	"time"
)

type AuthServer struct {
	proto.UnimplementedUserServiceServer
	rabbitMQConn *amqp.Connection
	*sql.DB
}

func NewGrpcServer(rabbitMQConn *amqp.Connection, db *sql.DB) {
	Connections.EnvLoader("./configs/env/.env")

	grpcServer := grpc.NewServer()

	AuthSrv := &AuthServer{
		rabbitMQConn: rabbitMQConn,
		DB:           db,
	}
	proto.RegisterUserServiceServer(grpcServer, AuthSrv)

	TcpPort := os.Getenv("TCP_PORT")

	listener, err := net.Listen("tcp", TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server started on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *AuthServer) Register(ctx context.Context, user *proto.User) (*proto.Confirm, error) {

	log.Printf("Registering a user : %s", user.Name)

	//Using validator
	v := validator.New()

	confirm := proto.Confirm{
		Ok:      false,
		Message: "Register failed",
	}

	if data.ValidateUser(v, user); !v.Valid() {
		if v.Errors["password"] != "" {
			err := errors.New(v.Errors["password"])
			return &confirm, err
		} else if v.Errors["name"] != "" {
			err := errors.New(v.Errors["name"])
			return &confirm, err
		}
	}

	err := data.UserModel.Insert(data.UserModel{DB: a.DB}, user)
	if err != nil {
		return nil, err
	}

	confirm.Ok = true
	confirm.Message = "Register was successfully"

	return &confirm, nil

}

func (a *AuthServer) Login(ctx context.Context, user *proto.User) (*proto.Token, error) {

	log.Printf("User trying to log in : %s", user.Email)

	//Using validator
	v := validator.New()
	tokenK := proto.Token{
		PlainText: "Broken Token",
		Hash:      []byte("a"),
		Id:        0,
		Expiry:    timestamppb.New(time.Now()),
		Scope:     "",
	}
	if data.ValidateUser(v, user); !v.Valid() {
		if v.Errors["password"] != "" {
			err := errors.New(v.Errors["password"])
			return &tokenK, err
		}
	}

	userPassPlain := user.Password.PlainText

	user, err := data.UserModel.GetByEmail(data.UserModel{DB: a.DB}, user.Email)

	user.Password.PlainText = userPassPlain
	if err != nil {
		return nil, err
	}

	match, err := data.Matches(user.Password.PlainText, user.Password.Hash)

	if err != nil {
		return nil, err
	}
	if !match {
		return nil, err
	}

	token, err := data.TokenModel.New(data.TokenModel{DB: a.DB}, user.Id, 3*24*time.Hour, data.ScopeAuthentication)

	return token, nil

}

func (a *AuthServer) UpdateUser(ctx context.Context, update *proto.Update) (*proto.Confirm, error) {

	log.Printf("Updating a user : %s", update.Name)

	user, _ := data.UserModel.GetForToken(data.UserModel{DB: a.DB}, data.ScopeAuthentication, update.TokenValue)
	user.Name = update.Name
	err := data.UserModel.Update(data.UserModel{DB: a.DB}, user)

	if err != nil {
		return &proto.Confirm{
			Ok:      false,
			Message: "Update was failed",
		}, err
	}

	return &proto.Confirm{
		Ok:      true,
		Message: "Update was successfully",
	}, nil
}

func (a *AuthServer) ProfileUser(ctx context.Context, update *proto.Profile) (*proto.User, error) {

	log.Printf("Getting Profile")
	fmt.Println(update.TokenValue)
	user, err := data.UserModel.GetForToken(data.UserModel{DB: a.DB}, data.ScopeAuthentication, update.TokenValue)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
