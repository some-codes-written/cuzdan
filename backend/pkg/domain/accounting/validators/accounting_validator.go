package validators

import (
	"errors"
	"immortality/pkg/domain/accounting/models"
)

func ValidateAccounting(s models.Accounting) (error, bool) {

	if (s == models.Accounting{}) {
		return errors.New("accounting is empty"), false
	}

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
