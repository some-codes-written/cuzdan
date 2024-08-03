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

func (ops *CompanyOperations) CreateCompany(model *company_dtos.CompanyDto) error {

	err := ops.Validator.ValidateCompany(model)
	if err != nil {
		return err
	}

	dto := company_models.Company{
		CompanyTypeID: model.CompanyTypeID,
		Name:          model.Name,
		Description:   model.Description,
		Email:         model.Email,
		Phone:         model.Phone,
		Website:       model.Website,
		Address:       model.Address,
		IsActive:      model.IsActive,
		AuthPersonId:  model.AuthPersonId,
	}

	err = ops.Store.CreateCompany(&dto)

	if err != nil {
		return err
	}

	return nil
}

func (ops *CompanyOperations) CreateCompanyType(model *company_dtos.CompanyTypeDto) (*company_dtos.CompanyTypeDto, error) {

	err := ops.Validator.ValidateCompanyType(model)
	if err != nil {
		return nil, err
	}

	temp := company_models.CompanyType{
		Name:        model.Name,
		Description: model.Description,
	}

	err = ops.Store.CreateCompanyType(&temp)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (ops *CompanyOperations) GetCompany(id int) (*company_dtos.CompanyDto, error) {

	res, err := ops.Store.GetCompany(id)
	if err != nil {
		return nil, err
	}

	temp := company_dtos.CompanyDto{
		CompanyTypeID: res.CompanyTypeID,
		Name:          res.Name,
		Description:   res.Description,
		Email:         res.Email,
		Phone:         res.Phone,
		Website:       res.Website,
		Address:       res.Address,
		IsActive:      res.IsActive,
		AuthPersonId:  res.AuthPersonId,
	}
	return &temp, nil
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
