package person_store

import "immortality/pkg/common"

type PersonStore struct {
	common.StoreBase
}

func NewPersonStore() *PersonStore {
	store := new(PersonStore)
	store.Connect()
	return store
}
