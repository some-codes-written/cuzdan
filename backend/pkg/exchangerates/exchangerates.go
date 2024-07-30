package exchangerates

import (
	"encoding/xml"
	"io"
	"net/http"
)

const (
	TC   string = "http://www.tcmb.gov.tr/kurlar/today.xml"
	KKTC string = "http://www.mb.gov.ct.tr/kur/gunluk.xml"
)

func Tc() (Currency, error) {
	var res = downloadString(TC)
	var tc Currency
	if err := xml.Unmarshal([]byte(res), &tc); err != nil {
		return tc, err
	}
	return tc, nil
}

func Kktc() (CurrencyKKTC, error) {
	var res = downloadString(KKTC)
	var tc CurrencyKKTC
	if err := xml.Unmarshal([]byte(res), &tc); err != nil {
		return tc, err
	}
	return tc, nil
}

func downloadString(param string) string {
	res, err := http.Get(param)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
