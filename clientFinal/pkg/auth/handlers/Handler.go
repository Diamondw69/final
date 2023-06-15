package handlers

import (
	"clientFinal/internal/data"
	pb "clientFinal/pkg/auth/proto"
	"clientFinal/pkg/inventory/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func ConnectGrpc() *grpc.ClientConn {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Parsing form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	//Making user object
	user := pb.User{
		Name:    name,
		Email:   email,
		Role:    "user",
		Balance: 5000,
	}
	password1 := data.Set(password)
	user.Password = password1

	//Connecting to grpc
	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Calling register function from grpc
	_, err = client.Register(ctx, &user)
	if err != nil {
		log.Fatalf("Error caused by :%s", err)
	}

	http.Redirect(w, r, "/login", 303)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	password1 := r.Form.Get("password")
	email := r.Form.Get("email")

	//Making password object to fix nil bug
	password := pb.Password{
		PlainText: "",
		Hash:      []byte("a"),
	}

	//Creating user object to fill data
	user := pb.User{
		Email: email,
	}

	user.Password = &password
	user.Password.PlainText = password1

	conn := ConnectGrpc()
	defer conn.Close()
	conn2, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	client2 := proto.NewInventoryServiceClient(conn2)
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token, err := client.Login(ctx, &user)
	if err != nil {
		log.Fatalf("Error caused by :%s", err)
	}

	req := proto.InventoryRequest{
		TokenValue: token.PlainText,
		Id:         token.Id,
	}

	inventory, err := client2.GetInventory(ctx, &req)
	if inventory == nil {
		conf, _ := client2.NewInventory(ctx, &req)
		fmt.Println(conf.Message)
	}

	//Creating a cookie to use token in future
	cookie := http.Cookie{
		Name:   "token",
		Value:  token.PlainText,
		MaxAge: 3600,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", 303)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//Searching for a cookie, if one exists it will be deleted
	//Our project`s frontend works on this cookie, that is why deleting a cookie will make a user unauthorized
	logout, err := r.Cookie("token")
	if err == nil {
		logout.MaxAge = -1
		http.SetCookie(w, logout)
	}
	http.Redirect(w, r, "/", 303)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	token, err := r.Cookie("token")
	if err != nil {
		log.Fatalf("Token nety")
	}

	conn := ConnectGrpc()
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Making update request
	update := pb.Update{
		TokenValue: token.Value,
		Name:       r.Form.Get("name"),
	}

	//Getting confirm response
	confirm, _ := client.UpdateUser(ctx, &update)

	if confirm.Ok {
		http.Redirect(w, r, "/profile", 303)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}
