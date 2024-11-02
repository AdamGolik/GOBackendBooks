package main

import (
	"awesomeProject3/cmd/api"
	database "awesomeProject3/cmd/db"
	"awesomeProject3/cmd/env"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config := env.GetMysqlConfig()
	database.ConnectDatabase(config)

	route := mux.NewRouter()
	apiV1 := route.PathPrefix("/api/v1").Subrouter()

	// Endpoints dla książek
	apiV1.HandleFunc("/book", api.GetBook).Methods("GET")
	apiV1.HandleFunc("/book/{id}", api.GetBookById).Methods("GET")
	apiV1.HandleFunc("/book", api.CreateBook).Methods("POST")
	apiV1.HandleFunc("/book/{id}", api.UpdateBookById).Methods("PUT")
	apiV1.HandleFunc("/book/{id}", api.DeleteBookById).Methods("DELETE")
	apiV1.HandleFunc("/book/search/{value}", api.SearchBook).Methods("GET")

	// Endpoints dla użytkowników
	apiV1.HandleFunc("/user", api.GetUser).Methods("GET")
	apiV1.HandleFunc("/user/{id}", api.GetUserById).Methods("GET")
	apiV1.HandleFunc("/user", api.CreateUser).Methods("POST")
	apiV1.HandleFunc("/user/{id}", api.UpdateUserById).Methods("PUT")
	apiV1.HandleFunc("/user/{id}", api.DeleteUserById).Methods("DELETE")
	apiV1.HandleFunc("/user/search/{value}", api.SearchUser).Methods("GET")

	http.Handle("/", route)
	http.ListenAndServe(":8080", nil)
}
