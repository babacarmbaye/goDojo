package main

import (
	"fmt"
	"log"
	"net/http"

	"api/db"
	"api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Connect to database
	collection := db.ConnectDB()

	// Create an instance of Handler with the collection
	handler := &handlers.Handler{Collection: collection}

	// Setup routes
	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/books", handler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")

	// Start server
	fmt.Println("Server listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
