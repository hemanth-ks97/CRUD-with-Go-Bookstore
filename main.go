package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `josn:"lastname"`
}

var Books = make(map[string]Book)

func slicebuilder() []Book {
	res := make([]Book, len(Books))

	i := 0
	for _,v := range Books{
		res[i] = v
		i += 1
	}

	return res
}

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	res := slicebuilder()

	json.NewEncoder(w).Encode(res)
}

func getBookbyID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)

	res := Books[params["id"]]

	json.NewEncoder(w).Encode(res)
}

func deleteBookbyID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	delete(Books, params["id"])

	json.NewEncoder(w).Encode(slicebuilder())
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	Books[book.ID] = book

	json.NewEncoder(w).Encode(Books[book.ID])
}

func updateBookbyID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	Books[book.ID] = book

	json.NewEncoder(w).Encode(Books[book.ID])
}

func main(){
	
	Books["0"] = Book{ID:"0", ISBN:"1341234asdfasdf", Title:"FirstBook", Author: &Author{Firstname:"Jared", Lastname:"Letto"}}
	Books["1"] = Book{ID:"1", ISBN:"1341234asdfasdg", Title:"SecondBook", Author: &Author{Firstname:"Michael", Lastname:"Faraday"}}
	Books["2"] = Book{ID:"2", ISBN:"1341234asdfasdh", Title:"ThirdBook", Author: &Author{Firstname:"Scott", Lastname:"Tomi"}}

	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBookbyID).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", deleteBookbyID).Methods("DELETE")
	r.HandleFunc("/books/{id}", updateBookbyID).Methods("PUT")

	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", r)
}