package persons

import (
	"errors"
	"immortality/pkg/common"
	"immortality/pkg/domain/persons/person_models"

	"gorm.io/gorm"
)

type PersonStore struct {
	common.StoreBase
}

func NewPersonStore() *PersonStore {
	store := new(PersonStore)
	store.Connect()
	return store
}

func (s *PersonStore) GetPerson(id uint) (*person_models.Person, error) {
	var person *person_models.Person
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.First(&person_models.Person{}, id).Table(person_models.PERSONS)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("Person not found")
		}
		res.First(&person)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return person, nil
}

func (s *PersonStore) GetPersons() ([]*person_models.Person, error) {
	var persons []*person_models.Person
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.Table(person_models.PERSONS).Find(&persons)
		if res == nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return persons, nil
}
