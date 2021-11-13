package utils

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
	"os"
)

func CreateDb() {
	// Deleting the sqlite file because they are intended to be temporary
	//TODO
	err := os.Remove("sqlite-database.db")
	if err != nil {
		fmt.Println("Error whilst deleting database")
		//panic(err)
	}

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		panic(err)
	}

	file.Close()

	log.Println("sqlite-database.db created")

	db, err := gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	migrateError := db.AutoMigrate(Config{})
	if migrateError != nil {
		fmt.Println("Error while migrating Database")
		log.Fatal(migrateError)
		return
	}
	return
}
