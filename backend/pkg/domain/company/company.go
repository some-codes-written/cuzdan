package company

import (
	"immortality/database"
	"immortality/pkg/domain/company/company_models"
	"immortality/pkg/domain/company/company_seeders"
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
	company_seeders.SeedCompany(db)
}

func GetInterfaces() []interface{} {
	temp := make([]interface{}, 0)
	return append(temp,
		&company_models.Company{},
		&company_models.CompanyType{},
		&company_models.CompanyPerson{},
		&company_models.CompanyPersonType{},
	)
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
