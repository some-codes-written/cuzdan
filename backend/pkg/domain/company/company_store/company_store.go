package company_store

import (
	"immortality/pkg/domain/company/company_dtos"
	"immortality/pkg/domain/company/company_models"

	"gorm.io/gorm"
)

const (
	COMPANIES       = "companies"
	COMPANY_PERSONS = "company_persons"
)

// GetCompanyPerson gets a company person array by company id
func (store *CompanyStore) GetCompanyPerson(company_id int) (*[]company_dtos.CompanyPersonDto, error) {

	temp := []company_dtos.CompanyPersonDto{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(COMPANY_PERSONS).Where("company_id = ?", company_id).Find(&temp).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &temp, nil

}

// GetCompanyList gets a list of all companies
func (store *CompanyStore) GetCompanyList() (*[]company_models.Company, error) {

	res := []company_models.Company{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(COMPANIES).Find(&res).Error
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

// GetCompany gets a company by id
func (store *CompanyStore) GetCompany(id int) (*company_models.Company, error) {

	res := company_models.Company{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(COMPANIES).Where("id = ?", id).First(&res).Error
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

// CreateCompany creates a new company
func (store *CompanyStore) CreateCompany(model *company_models.Company) error {

	res := store.Db.Transaction(func(tx *gorm.DB) error {

		//TODO: Add db validation

		xres := tx.Table(COMPANIES).Create(&model)
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

// CreateCompany creates a new company person
func (store *CompanyStore) CreateCompanyPerson(model *company_models.CompanyPerson) error {

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(COMPANY_PERSONS).Create(&model).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

// RemoveCompanyPerson removes a company person by id
func (store *CompanyStore) RemoveCompanyPerson(company_person_id int) error {

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(COMPANY_PERSONS).Where("id = ?", company_person_id).Delete(&company_models.CompanyPerson{})
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// RemoveCompany removes a company by id and removes all company persons
func (store *CompanyStore) RemoveCompany(company_id int) error {

	err := store.Db.Transaction(func(tx *gorm.DB) error {

		err := tx.Table(COMPANY_PERSONS).Where("company_id = ?", company_id).Delete(&company_models.CompanyPerson{}).Error
		if err != nil {
			return err
		}

		err = tx.Table(COMPANIES).Where("id = ?", company_id).Delete(&company_models.Company{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil

}

// FindCompany finds a company by model
func (store *CompanyStore) FindCompany(model *company_models.Company) (*company_models.Company, error) {

	res := company_models.Company{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(COMPANIES).Where(&model).Find(&res).Error
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
