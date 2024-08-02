package store

import (
	"immortality/pkg/common"
	"immortality/pkg/domain/company/dtos"
	"immortality/pkg/domain/company/models"

	"gorm.io/gorm"
)

type CompanyStore struct {
	common.StoreBase
}

func NewCompanyStore() *CompanyStore {
	store := new(CompanyStore)
	store.Connect()
	return store
}

func (store *CompanyStore) GetCompany(id int) error {
	company := models.Company{}
	res := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).First(&company).Error
		if err != nil {
			return err
		}
		return nil
	})
	if res != nil {
		return res
	}
	return nil
}

func (store *CompanyStore) CreateCompany(model dtos.CompanyDto) error {

	res := store.Db.Transaction(func(tx *gorm.DB) error {

		//TODO: Add db validation
		//TODO: add automapper
		company := models.Company{
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

		xres := tx.Create(&company)
		if xres.Error != nil {
			return xres.Error
		}
		return nil

	})

	if res != nil {
		return res
	}

	return nil
}
