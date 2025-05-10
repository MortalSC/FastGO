package handler

import (
	"github.com/MortalSC/FastGO/internal/apiserver/biz"
	"github.com/MortalSC/FastGO/internal/pkg/validation"
)

// Handler is a struct that defines the handler for the API server
type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

// NewHandler creates a new Handler instance
func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz: biz,
		val: val,
	}
}
