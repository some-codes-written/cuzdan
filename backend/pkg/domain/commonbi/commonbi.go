package commonbi

import "immortality/database"

func Setup() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	// SeedingDatabase()
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
