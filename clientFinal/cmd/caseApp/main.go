package main

import (
	customRouter "clientFinal/cmd/caseApp/router"
	"fmt"
	"net/http"
)

func main() {
	router := customRouter.MakeRouter()
	fmt.Println("Starting a server on port :8080")
	http.ListenAndServe(":8080", router)

}
