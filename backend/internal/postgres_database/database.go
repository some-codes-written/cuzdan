package postgres_database

import (
	"fmt"
	"immortality/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	config := config.GetConfig()

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DatabaseConfig.Host, config.DatabaseConfig.Username, config.DatabaseConfig.Password, config.DatabaseConfig.Name, config.DatabaseConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

func Disconnect(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func AutoMigrate(dst ...interface{}) error {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(dst...)
}
