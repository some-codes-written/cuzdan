package exchangerates

import (
	"encoding/xml"
	"time"
)

type Currency struct {
	XMLName  xml.Name `xml:"Tarih_Date"`
	Currency []CurrencyItem
}

type CurrencyItem struct {
	XMLName         xml.Name `xml:"Currency"`
	CurrencyCode    string   `xml:"CurrencyCode,attr"`
	Unit            int      `xml:"Unit"`
	Name            string   `xml:"Isim"`
	Currencyname    string   `xml:"CurrencyName"`
	ForexBuying     float32  `xml:"ForexBuying"`
	ForexSelling    float32  `xml:"ForexSelling"`
	BanknoteBuying  float32  `xml:"BanknoteBuying"`
	BanknoteSelling float32  `xml:"BanknoteSelling"`
	CrossRateUSD    float32  `xml:"CrossRateUSD"`
	CrossRateOther  float32  `xml:"CrossRateOther"`
}

type CurrencyKKTC struct {
	XMLName         xml.Name            `xml:"KKTCMB_Doviz_Kurlari"`
	CurrencyDate    customTime          `xml:"Kur_Tarihi"`
	AnnouncementNo  customTime          `xml:"Duyuru_No"`
	ValidDate       customTime          `xml:"Gecerli_Tarih_Araligi"`
	ValidDateStr    customTime          `xml:"Gecerli_Tarih_Araligi_Str"`
	Currencies      []CurrencyItemKKTC  `xml:"Resmi_Kurlar"`
	CrossCurrencies []CrossCurrencyItem `xml:"Capraz_Kurlar"`
}

type CurrencyItemKKTC struct {
	XMLName          xml.Name `xml:"Resmi_Kurlar"`
	Unit             int      `xml:"Birim"`
	Symbol           int      `xml:"Sembol"`
	Name             int      `xml:"Isim"`
	Buying           int      `xml:"Doviz_Alis"`
	Selling          int      `xml:"Doviz_Satis"`
	EffectiveBuying  int      `xml:"Efektif_Alis"`
	EffectiveSelling int      `xml:"Efektif_Satis"`
}

type CrossCurrencyItem struct {
	XMLName            xml.Name `xml:"Capraz_Kurlar"`
	SymbolFrom         string   `xml:"Sembol_From"`
	CurrencyFrom       float32  `xml:"Doviz_From"`
	CrossCurrencyValue float32  `xml:"Capraz_Kur_Deger"`
	CurrencyNameTo     string   `xml:"Doviz_To"`
	SymbolTo           string   `xml:"Sembol_To"`
}

type customTime struct {
	time.Time
}
