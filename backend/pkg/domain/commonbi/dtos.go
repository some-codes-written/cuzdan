package commonbi

import "time"

type CurrencyDto struct {
	ID              uint   `json:"id" example:"1"`                                      // 1, 2, 3
	Code            string `json:"code" example:"1"`                                    // USD, EUR, TRY
	IsLocalCurrency bool   `json:"isLocalCurrency" example:"1"`                         // 1, 0
	Description     string `json:"description" example:"US Dollar, Euro, Turkish Lira"` // US Dollar, Euro, Turkish Lira
}

type ExchangeRateDto struct {
	ID         uint      `json:"id" example:"1"`         // 1, 2, 3
	CurrencyId uint      `json:"currencyId" example:"1"` // Currency
	RateDate   time.Time `json:"rateDate" example:"1"`   // Rate date
	Rate       float64   `json:"rate" example:"1"`       // Rate
}
