package handlers

import (
	"encoding/json"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-ybaspinar/domain/books"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-ybaspinar/pkg/newpsqldb"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ApiRoutes() {
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{param1}", getBook).Methods("GET")
	router.HandleFunc("/byid/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/deletedbook/{param1}", getDeletedBook).Methods("GET")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	db, _ := newpsqldb.NewPsqlDB()
	booksRepo := books.NewBooksepository(db)
	booksRepo.Delete(id)
	writer.WriteHeader(http.StatusOK)
}

func getDeletedBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	param1, _ := strconv.Atoi(vars["param1"])
	db, _ := newpsqldb.NewPsqlDB()
	booksRepo := books.NewBooksepository(db)
	payload := booksRepo.FindDeleted(string(param1))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(payload)
}

func getBookByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	db, _ := newpsqldb.NewPsqlDB()
	booksRepo := books.NewBooksepository(db)
	payload := booksRepo.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(payload)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	param1, _ := strconv.Atoi(vars["param1"])
	db, _ := newpsqldb.NewPsqlDB()
	booksRepo := books.NewBooksepository(db)
	payload := booksRepo.Search(string(param1))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(payload)
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	db, _ := newpsqldb.NewPsqlDB()
	booksRepo := books.NewBooksepository(db)
	payload := booksRepo.ListAll()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(payload)
	println("getBooks")
}
