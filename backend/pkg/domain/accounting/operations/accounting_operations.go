package operations

import "immortality/pkg/domain/accounting/store"

type AccountingOperations struct {
	store *store.AccountingStore
}

func NewAccountingOperations() *AccountingOperations {
	return &AccountingOperations{
		store: store.NewAccountingStore(),
	}
}
