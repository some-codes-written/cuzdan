package initializer

import (
	"immortality/internal/jobs"
	"immortality/pkg/domain/commonbi"
	"immortality/pkg/domain/company"
	"immortality/pkg/domain/persons"
	"immortality/pkg/domain/users"
)

func Initialize() {

	var models []interface{}

	models = append(models, users.GetInterfaces()...)
	models = append(models, persons.GetInterfaces()...)
	models = append(models, commonbi.GetInterfaces()...)
	models = append(models, company.GetInterfaces()...)

	store := NewInitStore()
	store.Initialize(models...)

	//jobs.Schedule()

	commonbi.SeedCurrency(store.Db)

	persons.SeedPerson(store.Db)

	users.SeedUser(store.Db)

	company.SeedCompany(store.Db)

	jobs.RunAllOneTime()
}
