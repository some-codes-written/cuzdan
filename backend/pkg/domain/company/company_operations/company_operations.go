package company_operations

import (
	"immortality/pkg/domain/company/company_dtos"
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

func (store *CompanyOperations) CreateCompany(company company_dtos.CompanyDto) error {

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

func (store *CompanyOperations) GetCompany(id int) error {

	err := store.GetCompany(id)

	if err != nil {
		return err
	}

	return nil
}
