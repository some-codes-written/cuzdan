package database

import (
	"immortality/internal/postgres_database"

	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return postgres_database.Connect()
	//return postgres_database.Connect()
}

func Disconnect(db *gorm.DB) {
	postgres_database.Disconnect(db)
}

func Migrate(dst ...interface{}) error {
	return postgres_database.AutoMigrate(dst...)
}
