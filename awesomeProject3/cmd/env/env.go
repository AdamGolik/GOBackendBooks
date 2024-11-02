package env

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

// Struktura Book, reprezentująca książkę w bazie danych
type Book struct {
	Id       int     `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Author   string  `json:"author"`
}

// Struktura User, reprezentująca użytkownika w bazie danych
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Password string `json:"password"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

// Struktura do logowania
type LoginDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// BeforeCreate hook to hash password and validate email before creating a new user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = hashPassword(u.Password)
	if !isValidEmail(u.Email) {
		return fmt.Errorf("Invalid email address")
	}
	return nil
}

// Funkcja szyfrująca hasło za pomocą SHA-256
func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// Funkcja walidująca email
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Sprawdzanie hasła
func CheckPasswordHash(password, hash string) bool {
	return hashPassword(password) == hash
}

// Struktura MysqlConfig przechowująca konfigurację dla połączenia z bazą danych MySQL
type MysqlConfig struct {
	Host string
	Port int
	User string
	Pass string
	Db   string
}

// Funkcja GetMysqlConfig zwracająca konfigurację dla połączenia z bazą danych MySQL
func GetMysqlConfig() MysqlConfig {
	return MysqlConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",  // Zaktualizuj według swoich potrzeb
		Pass: "xxx",   // Zaktualizuj według swoich potrzeb
		Db:   "Books", // Zaktualizuj według swoich potrzeb
	}
}
