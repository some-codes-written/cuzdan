package persons

import (
	"errors"
	"immortality/pkg/common"

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

func (s *PersonStore) GetPerson(id uint) (*Person, error) {
	var person *Person
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.First(&Person{}, id).Table(PERSONS)
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

func (s *PersonStore) GetPersons() ([]*Person, error) {
	var persons []*Person
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.Table(PERSONS).Find(&persons)
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
