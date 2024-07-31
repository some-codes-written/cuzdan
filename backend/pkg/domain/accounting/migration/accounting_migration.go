package migration

import (
	"immortality/database"
	"immortality/pkg/domain/accounting/models"
)

func Setup() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	// SeedingDatabase()

}

func SeedingDatabase() {
	//TODO: Seeding
}

func GetInterfaces() []interface{} {
	temp_models := make([]interface{}, 0)
	return append(temp_models,
		&models.AccountingType{},
		&models.ExpenseType{},
		&models.Accounting{},
		&models.CashflowType{},
		&models.Cashflow{},
		&models.CashflowAccountingDetail{},
	)
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	return database.Migrate(dst...)
}
