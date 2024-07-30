package models

import "immortality/pkg/domain/commonbi"

// / restapi.Response
type ApiResponse struct {
	Status       string      `json:"status"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

type CurrencyListResponse struct {
	ApiResponse

	Currencies []commonbi.CurrencyDto `json:"currencies"`
}

type ExchangeRateListResponse struct {
	ApiResponse

	ExchangeRates []commonbi.ExchangeRateDto `json:"exchangeRates"`
}
