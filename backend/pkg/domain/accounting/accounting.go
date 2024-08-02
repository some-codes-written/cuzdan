package accounting

import "immortality/pkg/domain/accounting/accounting_migration"

func SetupDatabase() {
	accounting_migration.Setup()
}
