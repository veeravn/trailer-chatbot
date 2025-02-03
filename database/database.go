// database/database.go
package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDatabase connects to the database and runs migrations
func InitializeDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("unloading.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&Trailer{})
}

// Trailer represents unloading requests
type Trailer struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Number     string `json:"number"`
	DockingBay string `json:"docking_bay"`
	Status     string `json:"status"`
}

