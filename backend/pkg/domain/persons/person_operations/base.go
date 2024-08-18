package person_operations

import (
	"immortality/pkg/domain/persons/person_store"
	"immortality/pkg/domain/persons/person_validators"
)

type PersonOperations struct {
	Store     *person_store.PersonStore
	Validator *person_validators.PersonValidators
}

func NewPersonOperations() *PersonOperations {
	return &PersonOperations{
		Store:     person_store.NewPersonStore(),
		Validator: person_validators.NewPersonValidator(),
	}
}
