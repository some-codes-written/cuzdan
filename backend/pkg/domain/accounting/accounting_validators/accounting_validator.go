package accounting_validators

import (
	"errors"
	"immortality/pkg/domain/accounting/accounting_models"
)

func ValidateAccounting(s accounting_models.Accounting) (error, bool) {

	if s.AccountingTypeID == 0 {
		return errors.New("accountingTypeID is empty"), false
	}

	if s.ExpenseTypeID == 0 {
		return errors.New("expenseTypeID is empty"), false
	}

	if s.PersonId == 0 {
		return errors.New("personId is empty"), false
	}

	if s.Amount == 0 {
		return errors.New("amount cannot be empty"), false
	}

	return nil, true

}
