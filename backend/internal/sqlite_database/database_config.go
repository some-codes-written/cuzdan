package sqlite_database

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	SettingName  string `json:"settingName"`
	SettingValue string `json:"settingValue"`
}

func PrepareConfig() error {
	return AutoMigrate(&Settings{})
}
