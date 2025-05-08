package errorx

import "net/http"

var (
	OK          = &ErrorX{Code: http.StatusOK, Reason: "OK", Message: "ok"}
	ErrInternal = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalServerError", Message: "internal server error"}
	ErrNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound", Message: "not found"}
	// ErrBadRequest   = &ErrorX{Code: http.StatusBadRequest, Reason: "BadRequest", Message: "bad request"}
	// ErrUnauthorized = &ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthorized", Message: "unauthorized"}
	// ErrForbidden    = &ErrorX{Code: http.StatusForbidden, Reason: "Forbidden", Message: "forbidden"}
)
