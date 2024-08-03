package company_operations

import (
	"immortality/pkg/domain/company/company_dtos"
	"immortality/pkg/domain/company/company_models"
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

	dto := company_models.Company{
		CompanyTypeID: company.CompanyTypeID,
		Name:          company.Name,
		Description:   company.Description,
		Email:         company.Email,
		Phone:         company.Phone,
		Website:       company.Website,
		Address:       company.Address,
		IsActive:      company.IsActive,
		AuthPersonId:  company.AuthPersonId,
	}

	err = ops.Store.CreateCompany(&dto)

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

func (ops *CompanyOperations) GetCompanyType(id int) (*company_dtos.CompanyTypeDto, error) {

	res, err := ops.Store.GetCompanyType(id)

	if err != nil {
		return &company_dtos.CompanyTypeDto{}, err
	}

	dto := company_dtos.CompanyTypeDto{
		Name:        res.Name,
		Description: res.Description,
	}

	return &dto, nil

}
