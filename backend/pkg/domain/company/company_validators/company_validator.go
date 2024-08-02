package company_validators

import (
	"errors"
	"immortality/pkg/domain/company/company_dtos"
)

func ValidateCompany(company company_dtos.CompanyDto) error {
	if company.CompanyTypeID == 0 {
		return errors.New("company type id is required")
	}

	if company.Name == "" {
		return errors.New("company name is required")
	}

	return nil
}
