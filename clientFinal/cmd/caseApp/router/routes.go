package router

import (
	"clientFinal/pkg/auth/handlers"
	"clientFinal/pkg/caseItems/handlerCaseItem"
	handlercase "clientFinal/pkg/cases/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeRouter() *mux.Router {

	router := mux.NewRouter()

	//Static files
	router.HandleFunc("/images/{id}", handlerCaseItem.ImageHandler).Methods("GET", "OPTIONS")
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

	//crud User
	router.HandleFunc("/profile", handlers.ProfileHtmlHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/profile", handlers.UpdateUserHandler).Methods("POST", "OPTIONS")

	//Logout
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "OPTIONS")

	//crud CaseItems
	router.HandleFunc("/caseitemadd", handlerCaseItem.CreateCaseItemHTMLHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/caseitemadd", handlerCaseItem.CreateCaseItemHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/listcaseitems", handlerCaseItem.GetAllCaseItemsHTMLHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/deleteCaseItem/{id}", handlerCaseItem.DeleteCaseItemHandler).Methods("GET", "OPTIONS")

	//crud Cases
	router.HandleFunc("/caseadd", handlercase.CreateCaseHTMLHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/caseadd", handlercase.CreateCaseHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/case/{id}", handlercase.ViewCaseHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/cases", handlercase.ViewAllCaseHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/deletecase/{id}", handlercase.DeleteCaseHandler).Methods("GET", "OPTIONS")

	return router
}
