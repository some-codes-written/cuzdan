package company

import "immortality/database"

func Setup() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	// SeedingDatabase()
}

func SeedingDatabase() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	SeedCompany(db)
}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &Company{}, &CompanyType{}, &CompanyPerson{}, &CompanyPersonType{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
