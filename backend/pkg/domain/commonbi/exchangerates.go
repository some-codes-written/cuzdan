package commonbi

import (
	"immortality/pkg/exchangerates"
	"time"
)

func GatherAllExchangeRates() {

	store := NewCommonBIStore()

	res, _ := exchangerates.Tc()

	for _, v := range res.Currency {
		currency, _ := store.GetCurrencyByCode(v.CurrencyCode)
		exchangeRate := ExchangeRate{
			CurrencyID: currency.ID,
			Buying:     float64(v.ForexBuying),
			Selling:    float64(v.ForexSelling),
			RateDate:   time.Now(),
		}
		_ = store.CreateExchangeRate(&exchangeRate)
	}

}
