package accounting_migration

import (
	"immortality/database"
	"immortality/pkg/domain/accounting/accounting_models"
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
		&accounting_models.AccountingType{},
		&accounting_models.ExpenseType{},
		&accounting_models.Accounting{},
		&accounting_models.CashflowType{},
		&accounting_models.Cashflow{},
		&accounting_models.CashflowAccountingDetail{},
	)
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	return database.Migrate(dst...)
}
