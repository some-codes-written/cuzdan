package commonbi

import (
	"immortality/database"
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

	SeedCurrency(db)

	//TODO: Add more seed functions here
}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &Bank{}, &Branch{}, &CreditCardType{}, &Currency{}, &ExchangeRate{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}

func GetActiveUserId() *uint {
	return nil
}
