package users

import (
	"fmt"
	"immortality/pkg/domain/persons/person_models"
	"immortality/pkg/domain/users/user_models"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {

	res := db.Transaction(func(tx *gorm.DB) error {

		var person person_models.Person
		err := tx.Where("gsm = ?", "5372112339").First(&person)
		if err.Error != nil {
			fmt.Println("err: ", err.Error)
			return err.Error
		}
		user := user_models.User{
			Email:     "oguzhan.saricam@gmail.com",
			Gsm:       "",
			FirstName: "Oğuz",
			LastName:  "Sarıçam",
			PersonID:  person.ID,
		}
		user.SetDefaultsviaCreation()
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		cre := user_models.Credential{
			UserID:   1,
			Username: "ouzsrcm",
			Email:    "oguzhan.saricam@gmail.com",
			Password: "123456",
		}
		cre.SetDefaultsviaCreation()
		if err := tx.Create(&cre).Error; err != nil {
			return err
		}

		return nil

	})
	fmt.Println("res: ", res)
	return res
}
