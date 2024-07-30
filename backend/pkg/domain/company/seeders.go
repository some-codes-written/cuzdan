package company

import (
	"fmt"
	"immortality/pkg/domain/persons"

	"gorm.io/gorm"
)

func SeedCompany(db *gorm.DB) error {

	res := db.Transaction(func(tx *gorm.DB) error {

		companyType := CompanyType{
			Name: "Organization",
		}
		companyType.SetDefaultsviaCreation()
		if err := tx.Create(&companyType).Error; err != nil {
			return err
		}

		var person persons.Person
		err := tx.Where("gsm = ?", "5372112339").First(&person)
		if err.Error != nil {
			fmt.Println("err: ", err.Error)
			return err.Error
		}

		company := Company{
			CompanyTypeID: companyType.ID,
			Name:          "General",
			Description:   "General Description",
			Email:         "general@mail.com",
			Phone:         "123456789",
			Website:       "https://www.general.com/",
			Address:       "General Address",
			IsActive:      true,
			AuthPersonId:  person.ID,
		}
		company.SetDefaultsviaCreation()
		if err := tx.Create(&company).Error; err != nil {
			return err
		}
		indivCompanyType := CompanyType{
			Name: "Individual",
		}
		indivCompanyType.SetDefaultsviaCreation()
		if err := tx.Create(&indivCompanyType).Error; err != nil {
			return err
		}

		companyPersonType := CompanyPersonType{
			Name: "Owner",
		}
		companyPersonType.SetDefaultsviaCreation()
		if err := tx.Create(&companyPersonType).Error; err != nil {
			return err
		}

		return nil
	})
	return res
}
