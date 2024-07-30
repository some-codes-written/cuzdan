package commonbi

import (
	"immortality/pkg/common"
	"time"
)

const (
	BANKS            = "banks"
	BRANCHES         = "branches"
	CREDIT_CARD      = "credit_cards"
	CREDIT_CARD_TYPE = "credit_card_types"
	CURRENCIES       = "currencies"
	EXCHANGE_RATES   = "exchange_rates"
)

type Bank struct {
	common.EntityModel

	Name        string `json:"name" example:"Bank of America"`
	Address     string `json:"address" example:"1234 Main St"`
	Phone       string `json:"phone" example:"555-555-5555"`
	Email       string `json:"email" example:"john.doe@gmail.com"`
	Website     string `json:"website" example:"https://www.google.com"`
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}

type Branch struct {
	common.EntityModel

	BankID      uint   `json:"bankId" example:"1"` // Banka bilgileri
	Bank        Bank   `json:"bank"`               // Banka bilgileri
	Name        string `json:"name" example:"Bank of America - Main Branch"`
	Address     string `json:"address" example:"1234 Main St"`
	Phone       string `json:"phone" example:"555-555-5555"`
	Email       string `json:"email" example:"john.doe@gmail.com"`
	Website     string `json:"website" example:"https://www.google.com"`
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}

type CreditCardType struct {
	common.EntityModel

	Name        string `json:"name" example:"1"` // MasterCard, Visa, American Express, Discover
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}

type Currency struct {
	common.EntityModel

	Code            string `json:"code" example:"1"`                                    // USD, EUR, TRY
	IsLocalCurrency bool   `json:"isLocalCurrency" example:"1"`                         // 1, 0
	Description     string `json:"description" example:"US Dollar, Euro, Turkish Lira"` // US Dollar, Euro, Turkish Lira
}

type ExchangeRate struct {
	common.EntityModel

	CurrencyID uint      `json:"currencyId" example:"1"` // Currency
	RateDate   time.Time `json:"rateDate" example:"1"`   // Rate date
	Rate       float64   `json:"rate" example:"1"`       // Rate
	Buying     float64   `json:"buying" example:"1"`     // Buying
	Selling    float64   `json:"selling" example:"1"`    // Selling
}
