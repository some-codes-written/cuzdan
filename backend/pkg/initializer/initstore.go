package initializer

import "immortality/pkg/common"

type InitStore struct {
	common.StoreBase
}

func NewInitStore() *InitStore {
	store := new(InitStore)
	store.Connect()
	return store
}

func (s *InitStore) Initialize(entities ...interface{}) {
	s.Migrate(entities...)
}
