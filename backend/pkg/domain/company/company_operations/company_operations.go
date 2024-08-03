package company_operations

import (
	"immortality/pkg/domain/company/company_dtos"
	"immortality/pkg/domain/company/company_models"
)

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

func (ops *CompanyOperations) GetCompanyList() (*[]company_dtos.CompanyDto, error) {

	res, err := ops.Store.GetCompanyList()
	if err != nil {
		return nil, err
	}

	temp := []company_dtos.CompanyDto{}
	for _, item := range *res {
		temp = append(temp, company_dtos.CompanyDto{
			CompanyTypeID: item.CompanyTypeID,
			Name:          item.Name,
			Description:   item.Description,
			Email:         item.Email,
			Phone:         item.Phone,
			Website:       item.Website,
			Address:       item.Address,
			IsActive:      item.IsActive,
			AuthPersonId:  item.AuthPersonId,
		})
	}

	return &temp, nil
}

func (ops *CompanyOperations) RemoveCompany(id int) error {
	return ops.Store.RemoveCompany(id)
}

func (ops *CompanyOperations) RemoveCompanyPerson(id int) error {
	return ops.Store.RemoveCompanyPerson(id)
}
