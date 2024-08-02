package persons

import (
	"immortality/database"
	"immortality/pkg/domain/persons/person_models"
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
	SeedPerson(db)
}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &person_models.Person{}, &person_models.Address{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
