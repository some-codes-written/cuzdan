package accounting_store

import "immortality/pkg/common"

type AccountingStore struct {
	common.StoreBase
}

func NewAccountingStore() *AccountingStore {
	store := new(AccountingStore)
	store.Connect()
	return store
}
