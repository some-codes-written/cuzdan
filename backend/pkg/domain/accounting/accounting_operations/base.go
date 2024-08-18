package accounting_operations

import (
	"immortality/pkg/domain/accounting/accounting_store"
)

type AccountingOperations struct {
	store *accounting_store.AccountingStore
}

func NewAccountingOperations() *AccountingOperations {
	return &AccountingOperations{
		store: accounting_store.NewAccountingStore(),
	}
}
