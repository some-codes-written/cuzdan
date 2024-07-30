package jobs

import "immortality/pkg/domain/commonbi"

func TCExchangeRates() {

	commonbi.GatherAllExchangeRates()

}
