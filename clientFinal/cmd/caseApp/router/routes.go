package router

import (
	"clientFinal/pkg/auth/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeRouter() *mux.Router {

	router := mux.NewRouter()

	//Static files
	fs := http.FileServer(http.Dir("./cmd"))
	router.PathPrefix("/cmd/").Handler(http.StripPrefix("/cmd/", fs))

	//Home pages
	router.HandleFunc("/", handlers.HomeHtmlHandler).Methods("GET", "OPTIONS")

	//Register pages
	router.HandleFunc("/register", handlers.RegisterHTMLHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST", "OPTIONS")

	//Login pages
	router.HandleFunc("/login", handlers.LoginHtmlHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST", "OPTIONS")

	//Logout
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "OPTIONS")

	return router
}
