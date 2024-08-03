package company_store

import "immortality/pkg/common"

type CompanyStore struct {
	common.StoreBase
}

func NewCompanyStore() *CompanyStore {
	store := new(CompanyStore)
	store.Connect()
	return store
}
