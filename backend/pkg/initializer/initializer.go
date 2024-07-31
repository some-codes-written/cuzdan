package initializer

import (
	"immortality/pkg/domain/accounting"
	"immortality/pkg/domain/commonbi"
	"immortality/pkg/domain/company"
	"immortality/pkg/domain/persons"
	"immortality/pkg/domain/users"
)

func Initialize() {

	users.SetupDatabase()

	persons.SetupDatabase()

	company.SetupDatabase()

	commonbi.SetupDatabase()

	accounting.SetupDatabase()

	//jobs.Schedule()
	// jobs.RunAllOneTime()
}
