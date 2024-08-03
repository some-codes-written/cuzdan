package company_operations

import (
	"immortality/pkg/domain/company/company_dtos"
	"immortality/pkg/domain/company/company_store"
	"immortality/pkg/domain/company/company_validators"
)

type CompanyOperations struct {
	Store     *company_store.CompanyStore
	Validator *company_validators.CompanyValidator
}

func NewCompanyOperations() *CompanyOperations {
	return &CompanyOperations{
		Store:     company_store.NewCompanyStore(),
		Validator: company_validators.NewCompanyValidator(),
	}
}

func (ops *CompanyOperations) CreateCompany(company company_dtos.CompanyDto) error {

	err := ops.Validator.ValidateCompany(company)
	if err != nil {
		return err
	}

	err = ops.Store.CreateCompany(company)

	if err != nil {
		return err
	}

	return nil
}

func (ops *CompanyOperations) GetCompany(id int) error {

	err := ops.Store.GetCompany(id)

	if err != nil {
		return err
	}

	return nil
}
