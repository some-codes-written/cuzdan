package mysql_database

import (
	"fmt"
	"immortality/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// obsoloete
func Connect() (*gorm.DB, error) {
	//TODO: Move this to a config file
	config := config.GetConfig()
	var dsn string = config.DatabaseConfig.Username + ":" + config.DatabaseConfig.Password + "@tcp(" + config.DatabaseConfig.Host + ":3306)/" + config.DatabaseConfig.Name + "?charset=utf8mb4&parseTime=True"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
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
