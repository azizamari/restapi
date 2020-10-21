package main

import (
	// "encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Auther Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request){

}

// Get a specific Book
func getBook(w http.ResponseWriter, r *http.Request){

}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request){

}

// Update Books
func updateBook(w http.ResponseWriter, r *http.Request){

}

// Delete a specific Book
func deleteBooks(w http.ResponseWriter, r *http.Request){

}

func main(){
	// Init Router
	r:=mux.NewRouter()

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))
}