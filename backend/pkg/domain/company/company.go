package company

import (
	"immortality/database"
	"immortality/pkg/domain/company/models"
)

func SetupDatabase() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	SeedingDatabase()
}

func SeedingDatabase() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	SeedCompany(db)
}

func GetInterfaces() []interface{} {
	temp := make([]interface{}, 0)
	return append(temp,
		&models.Company{},
		&models.CompanyType{},
		&models.CompanyPerson{},
		&models.CompanyPersonType{},
	)
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
