package validators

import (
	"errors"
	"immortality/pkg/domain/company/dtos"
)

func ValidateCompany(company dtos.CompanyDto) error {
	if company.CompanyTypeID == 0 {
		return errors.New("company type id is required")
	}

	if company.Name == "" {
		return errors.New("company name is required")
	}

	return nil
}
