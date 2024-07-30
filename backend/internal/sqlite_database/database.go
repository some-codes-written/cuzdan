package sqlite_database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	// Enable CGO for sqlite
	os.Setenv("CGO_ENABLED", "1")

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	return db, err
}

func AutoMigrate(dst ...interface{}) error {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(dst...)
}
