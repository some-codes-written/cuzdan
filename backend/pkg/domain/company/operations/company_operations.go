package operations

import (
	"immortality/pkg/domain/company/dtos"
	"immortality/pkg/domain/company/store"
	"immortality/pkg/domain/company/validators"
)

type CompanyOperations struct {
	store *store.CompanyStore
}

func NewCompanyOperations() *CompanyOperations {
	return &CompanyOperations{
		store: store.NewCompanyStore(),
	}
}

func (store *CompanyOperations) CreateCompany(company dtos.CompanyDto) error {

	err := validators.ValidateCompany(company)
	if err != nil {
		return err
	}

	err = store.CreateCompany(company)

	if err != nil {
		return err
	}

	return nil
}
