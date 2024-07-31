package accounting

import "immortality/pkg/domain/accounting/migration"

func SetupDatabase() {
	migration.Setup()
}
