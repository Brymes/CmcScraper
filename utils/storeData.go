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
	err := os.Remove("Coins.db")
	if err != nil {
		fmt.Println("Error whilst deleting database")
		//panic(err)
	}

	log.Println("Creating Coins.db...")
	file, err := os.Create("Coins.db") // Create SQLite file
	if err != nil {
		panic(err)
	}

	file.Close()

	log.Println("Coins.db created")

	db, err := gorm.Open(sqlite.Open("Coins.db"), &gorm.Config{})
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

func InsertData(resp *Config) {
	db, err := gorm.Open(sqlite.Open("Coins.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Create(&resp)

	//FIXME check if this works oo
	if result.Error != nil {
		log.Println(&resp)
		log.Println(&result.Error)
		panic(result.Error)
	}
	return
}
