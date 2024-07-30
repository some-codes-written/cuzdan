package commonbi

import (
	"gorm.io/gorm"
)

var currencies = map[string]Currency{
	"TRY": Currency{
		Code:            "TRY",
		IsLocalCurrency: true,
		Description:     "Turkish Lira",
	},
	"USD": Currency{
		Code:            "USD",
		IsLocalCurrency: false,
		Description:     "US Dollar",
	},
	"EUR": Currency{
		Code:            "EUR",
		IsLocalCurrency: false,
		Description:     "Euro",
	},
	"GBP": Currency{
		Code:            "GBP",
		IsLocalCurrency: false,
		Description:     "Pound Sterling",
	},
	"CHF": Currency{
		Code:            "CHF",
		IsLocalCurrency: false,
		Description:     "Swiss Franc",
	},
	"SEK": Currency{
		Code:            "SEK",
		IsLocalCurrency: false,
		Description:     "Swedish Krona",
	},
	"DKK": Currency{
		Code:            "DKK",
		IsLocalCurrency: false,
		Description:     "Danish Krone",
	},
	"NOK": Currency{
		Code:            "NOK",
		IsLocalCurrency: false,
		Description:     "Norwegian Krone",
	},
	"JPY": Currency{
		Code:            "JPY",
		IsLocalCurrency: false,
		Description:     "Japanese Yen",
	},
	"CNY": Currency{
		Code:            "CNY",
		IsLocalCurrency: false,
		Description:     "Chinese Yuan",
	},
	"KWD": Currency{
		Code:            "KWD",
		IsLocalCurrency: false,
		Description:     "Kuwaiti Dinar",
	},
	"AED": Currency{
		Code:            "AED",
		IsLocalCurrency: false,
		Description:     "United Arab Emirates Dirham",
	},
	"AUD": Currency{
		Code:            "AUD",
		IsLocalCurrency: false,
		Description:     "AUSTRALIAN DOLLAR",
	},
	"CAD": Currency{
		Code:            "CAD",
		IsLocalCurrency: false,
		Description:     "CANADIAN DOLLAR",
	},
	"SAR": Currency{
		Code:            "SAR",
		IsLocalCurrency: false,
		Description:     "SAUDI RIYAL",
	},
	"BGN": Currency{
		Code:            "BGN",
		IsLocalCurrency: false,
		Description:     "BULGARIAN LEV",
	},
	"RON": Currency{
		Code:            "RON",
		IsLocalCurrency: false,
		Description:     "NEW LEU",
	},
	"RUB": Currency{
		Code:            "RUB",
		IsLocalCurrency: false,
		Description:     "RUSSIAN ROUBLE",
	},
	"IRR": Currency{
		Code:            "IRR",
		IsLocalCurrency: false,
		Description:     "IRANIAN RIAL",
	},
	"PKR": Currency{
		Code:            "PKR",
		IsLocalCurrency: false,
		Description:     "PAKISTANI RUPEE",
	},
	"QAR": Currency{
		Code:            "QAR",
		IsLocalCurrency: false,
		Description:     "QATARI RIAL",
	},
	"KRW": Currency{
		Code:            "KRW",
		IsLocalCurrency: false,
		Description:     "SOUTH KOREAN WON",
	},
	"AZN": Currency{
		Code:            "AZN",
		IsLocalCurrency: false,
		Description:     "AZERBAIJANI NEW MANAT",
	},
	"XDR": Currency{
		Code:            "XDR",
		IsLocalCurrency: false,
		Description:     "SPECIAL DRAWING RIGHT (SDR)",
	},
}

func SeedCurrency(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		for _, currency := range currencies {
			currency.SetDefaultsviaCreation()
			tx.Create(&currency)
		}
		return nil
	})
}
