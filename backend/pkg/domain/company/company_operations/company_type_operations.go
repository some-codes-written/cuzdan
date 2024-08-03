package company_operations

import (
	"immortality/pkg/domain/company/company_dtos"
	"immortality/pkg/domain/company/company_models"
)

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

func (ops *CompanyOperations) GetCompanyTypeList() (*[]company_dtos.CompanyTypeDto, error) {

	res, err := ops.Store.GetCompanyTypeList()
	if err != nil {
		return nil, err
	}

	temp := []company_dtos.CompanyTypeDto{}
	for _, v := range *res {
		dto := company_dtos.CompanyTypeDto{
			Name:        v.Name,
			Description: v.Description,
		}
		temp = append(temp, dto)
	}

	return &temp, nil
}

func (ops *CompanyOperations) RemoveCompanyType(id int) error {

	err := ops.Store.RemoveCompanyType(id)
	if err != nil {
		return err
	}

	return nil
}
