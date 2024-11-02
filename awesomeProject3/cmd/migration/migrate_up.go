package main

import (
	"database/sql"
	"fmt"
	"log"

	"awesomeProject3/cmd/env"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := env.GetMysqlConfig()

	// Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.User, config.Pass, config.Host, config.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Nie udało się połączyć z bazą danych:", err)
	}
	defer db.Close()

	// Tworzenie bazy danych, jeśli nie istnieje
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Db))
	if err != nil {
		log.Fatal("Nie udało się utworzyć bazy danych:", err)
	}

	// Używanie bazy danych
	_, err = db.Exec(fmt.Sprintf("USE %s", config.Db))
	if err != nil {
		log.Fatal("Nie udało się ustawić bazy danych:", err)
	}

	// Tworzenie tabeli User, jeśli nie istnieje (uwzględnienie indeksów unikalnych)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		password VARCHAR(64) NOT NULL,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		phone VARCHAR(20),
		address VARCHAR(255)
	)`)
	if err != nil {
		log.Fatal("Nie udało się utworzyć tabeli users:", err)
	}

	// Tworzenie tabeli Book, jeśli nie istnieje
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		quantity INT NOT NULL,
		price FLOAT NOT NULL,
		author VARCHAR(100) NOT NULL
	)`)
	if err != nil {
		log.Fatal("Nie udało się utworzyć tabeli books:", err)
	}

	log.Println("Migracja w górę zakończona sukcesem.")
}
