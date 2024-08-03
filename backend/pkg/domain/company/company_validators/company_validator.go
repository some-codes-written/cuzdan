package company_validators

import (
	"errors"
	"immortality/pkg/domain/company/company_dtos"
)

type CompanyValidator struct {
}

func NewCompanyValidator() *CompanyValidator {
	return &CompanyValidator{}
}

func (s *CompanyValidator) ValidateCompanyType(companyType *company_dtos.CompanyTypeDto) error {

	if companyType.Name == "" || len(companyType.Name) < 2 {
		return errors.New("company type name is required and must be at least 2 characters")
	}

	return nil
}

func (s *CompanyValidator) ValidateCompany(company *company_dtos.CompanyDto) error {

	if company.Name == "" || len(company.Name) < 2 {
		return errors.New("company name is required and must be at least 2 characters")
	}

	if company.CompanyTypeID == 0 {
		return errors.New("company type is required")
	}

	return nil
}
