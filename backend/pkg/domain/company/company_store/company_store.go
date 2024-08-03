package company_store

import (
	"immortality/pkg/common"
	"immortality/pkg/domain/company/company_models"

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

func (store *CompanyStore) GetCompany(id int) (*company_models.Company, error) {

	res := company_models.Company{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).First(&res).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (store *CompanyStore) CreateCompany(model *company_models.Company) error {

	res := store.Db.Transaction(func(tx *gorm.DB) error {

		//TODO: Add db validation

		xres := tx.Create(&model)
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

func (store *CompanyStore) GetCompanyType(id int) (*company_models.CompanyType, error) {

	res := company_models.CompanyType{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).First(&res).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &res, nil
}
