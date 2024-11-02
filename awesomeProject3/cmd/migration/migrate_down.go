package main

import (
	"awesomeProject3/cmd/db"
	"awesomeProject3/cmd/env"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	config := env.GetMysqlConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pass, config.Host, config.Port, config.Db)

	var err error
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Prosta metoda usuwania tabel
	err = db.DB.Migrator().DropTable(&env.Book{}, &env.User{})
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	fmt.Println("Tables dropped successfully")
}
