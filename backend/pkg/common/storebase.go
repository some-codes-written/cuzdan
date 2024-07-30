package common

import (
	"immortality/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreBaseGeneric[T any] struct {
	Db *gorm.DB

	Item  T
	Items []T
}

type StoreBase struct {
	StoreBaseGeneric[EntityModel]
}

func (s *StoreBase) GetDbInstance() *gorm.DB {
	if s.Db == nil {
		s.Connect()
	}
	return s.Db
}

func (s *StoreBase) Connect() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	s.Db = db
	return db, nil
}

func (s *StoreBase) Disconnect() {
	database.Disconnect(s.Db)
}

func (s *StoreBaseGeneric[T]) GetByHash(db *gorm.DB, hash uuid.UUID) (T, error) {
	var entity T
	err := db.Where("hash = ?", hash).First(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (s *StoreBase) SoftDeleteRange(entities ...interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		for _, entity := range entities {

			if err := tx.Save(entity).Error; err != nil {
				return err
			}

		}
		return nil
	})
	return res
}

func (s *StoreBase) SoftDelete(entity interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(entity).Error; err != nil {
			return err
		}
		return nil
	})
	return res
}

func (s *StoreBase) DeleteRange(entities ...interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		for _, entity := range entities {
			if err := tx.Delete(entity).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return res
}

func (s *StoreBase) Delete(entity interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(entity).Error; err != nil {
			return err
		}
		return nil
	})
	return res
}

func (s *StoreBase) UpdateRange(entities ...interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		for _, entity := range entities {

			if err := tx.Save(entity).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return res
}

func (s *StoreBase) Update(entity interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(entity).Error; err != nil {
			return err
		}
		return nil
	})
	return res
}

func (s *StoreBase) CreateRange(entities ...interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		for _, entity := range entities {

			if err := tx.Create(entity).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return res
}

func (s *StoreBase) Create(entity interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(entity).Error; err != nil {
			return err
		}
		return nil
	})
	return res
}

func (s *StoreBase) Migrate(entities ...interface{}) error {
	res := s.Db.Transaction(func(tx *gorm.DB) error {
		for _, entity := range entities {
			if err := tx.AutoMigrate(entity); err != nil {
				return err
			}
		}
		return nil
	})
	return res
}
