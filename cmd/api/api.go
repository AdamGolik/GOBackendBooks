package api

import (
	"awesomeProject3/cmd/db"
	"awesomeProject3/cmd/env"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Struktura dla obsługi API
type API struct {
	addr string
	port string
}

// Nowa funkcja tworząca API
func NewAPI(addr, port string) *API {
	return &API{
		addr: addr,
		port: port,
	}
}

// Funkcje dla obsługi książek (Book)

func GetBook(writer http.ResponseWriter, request *http.Request) {
	var books []env.Book
	db.DB.Find(&books)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var book env.Book
	if err := db.DB.First(&book, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book env.Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&book).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var book env.Book
	if err := db.DB.First(&book, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	db.DB.Save(&book)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	if err := db.DB.Delete(&env.Book{}, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func SearchBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var books []env.Book
	searchTerm := "%" + params["value"] + "%"
	db.DB.Where("name LIKE ?", searchTerm).
		Or("author LIKE ?", searchTerm).
		Find(&books)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

// Funkcje dla obsługi użytkowników (User)

func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var user env.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func LoginUser(writer http.ResponseWriter, request *http.Request) {
	var loginDetails env.LoginDetails
	if err := json.NewDecoder(request.Body).Decode(&loginDetails); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var user env.User
	if err := db.DB.Where("username = ? OR email = ?", loginDetails.Username, loginDetails.Email).First(&user).Error; err != nil {
		http.Error(writer, "Invalid username or email", http.StatusUnauthorized)
		return
	}
	if !env.CheckPasswordHash(loginDetails.Password, user.Password) {
		http.Error(writer, "Invalid password", http.StatusUnauthorized)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	var users []env.User
	db.DB.Find(&users)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var user env.User
	if err := db.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func UpdateUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var user env.User
	if err := db.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	db.DB.Save(&user)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func DeleteUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	if err := db.DB.Delete(&env.User{}, params["id"]).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func SearchUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var users []env.User
	searchTerm := "%" + params["value"] + "%"
	db.DB.Where("username LIKE ?", searchTerm).
		Or("email LIKE ?", searchTerm).
		Or("phone LIKE ?", searchTerm).
		Or("address LIKE ?", searchTerm).
		Find(&users)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}
