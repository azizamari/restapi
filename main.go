package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
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

// Init books var as a slice Book struct
var books []Book

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "apllication/json")	
	json.NewEncoder(w).Encode(books)
}

// Get a specific Book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "apllication/json")	
	params:=mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "apllication/json")	
	var book Book
	_=json.NewDecoder(r.Body).Decode(&book)
	// book.ID=
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

	//Mock data @todo - implement DB
	books = append(books, Book{ID: "1",Isbn: "4485",Title: "Harry Potter", Author: &Author{FirstName: "JK", LastName: "Rowling"}})
	books = append(books, Book{ID: "2",Isbn: "49851",Title: "Rich Dad Poor Dad", Author: &Author{FirstName: "Robert", LastName: "Kyosaki"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))
}