package validation

import "github.com/MortalSC/FastGO/internal/apiserver/store"

type Validator struct {
	store store.IStore
}

func NewValidation(store store.IStore) *Validator {
	return &Validator{
		store: store,
	}
}
