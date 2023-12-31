package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/handlers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Transactions REST API!")
	fmt.Println("Transactions REST API")
}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/transactions", handlers.GetAllTransactions).Methods(http.MethodGet)
	myRouter.HandleFunc("/transaction/{id}", handlers.GetTransaction).Methods(http.MethodGet)
	myRouter.HandleFunc("/transaction/{id}", handlers.UpdateTransaction).Methods(http.MethodPut)
	myRouter.HandleFunc("/transaction", handlers.AddTransaction).Methods(http.MethodPost)
	myRouter.HandleFunc("/transaction/{id}", handlers.DeleteTransaction).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()

}

// https://github.com/Fuerback/transactions-go/tree/main
