package commonbi

import (
	"immortality/pkg/common"

	"gorm.io/gorm"
)

type CommonBIStore struct {
	common.StoreBase
}

func NewCommonBIStore() *CommonBIStore {
	store := new(CommonBIStore)
	store.Connect()
	return store
}

func (s *CommonBIStore) GetCurrencyByCode(code string) (Currency, error) {
	var currency Currency
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("code = ?", code).First(&currency)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return currency, tempres
	}
	return currency, nil
}

func (s *CommonBIStore) CreateCurrency(model *Currency) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {

		model.SetDefaultsviaCreation()
		res := tx.Create(model)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *CommonBIStore) GetCurrency(id uint) (*Currency, error) {
	var currency Currency
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&currency, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &currency, nil
}

func (s *CommonBIStore) GetCurrencies() (*[]Currency, error) {
	var currencies []Currency
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&currencies)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &currencies, nil
}

func (s *CommonBIStore) CreateExchangeRate(model *ExchangeRate) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {

		model.SetDefaultsviaCreation()
		res := tx.Create(model)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *CommonBIStore) GetExchangeRate(id uint) (*ExchangeRate, error) {
	var exchangeRate ExchangeRate
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&exchangeRate, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &exchangeRate, nil
}

func (s *CommonBIStore) GetLastExchangeRateByCurrency(currency_id uint) (*ExchangeRate, error) {
	var exchangeRate ExchangeRate
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(CURRENCIES).Where("currency_id = ?", currency_id).Last(&exchangeRate)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &exchangeRate, nil
}
