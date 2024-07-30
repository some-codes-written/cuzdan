package jobs

import (
	"github.com/jasonlvhit/gocron"
)

/*
	TODO: #6 burası db üzerinden ayarlanabilir olmalı şimdilik 00:00'da çalışacak şekilde ayarlandı
*/

func Schedule() {

	// gocron.Every(1).Day().At("00:00").Do(exchangerates.Tc)
	// gocron.Every(1).Day().At("00:00").Do(exchangerates.Kktc)

	gocron.Every(1).Day().At("00:00").Do(TCExchangeRates)
	<-gocron.Start()
}

// test için kullanılıyor
func RunAllOneTime() {
	TCExchangeRates()
}
