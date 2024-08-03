package company_operations

import (
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
