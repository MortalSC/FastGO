package errorx

import "net/http"

var (
	OK          = &ErrorX{Code: http.StatusOK, Reason: "OK", Message: "ok"}
	ErrInternal = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalServerError", Message: "internal server error"}
	ErrNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound", Message: "not found"}

	ErrDBWrite = &ErrorX{Code: http.StatusInternalServerError, Reason: "DBWrite", Message: "database write error"}
	ErrDBRead  = &ErrorX{Code: http.StatusInternalServerError, Reason: "DBRead", Message: "database read error"}

	ErrBind            = &ErrorX{Code: http.StatusBadRequest, Reason: "BindError", Message: "Error occurred while binding the request body to the struct."}
	ErrInvalidArgument = &ErrorX{Code: http.StatusBadRequest, Reason: "InvalidArgument", Message: "Argument verification failed."}

	ErrTokenInvalid = &ErrorX{Code: http.StatusUnauthorized, Reason: "TokenInvalid", Message: "invalid token"}
	ErrTokenExpired = &ErrorX{Code: http.StatusUnauthorized, Reason: "TokenExpired", Message: "token expired"}
	ErrSignToken    = &ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthenticated.SignToken", Message: "Error occurred while signing the JSON web token."}

	ErrInvalidPassword = &ErrorX{Code: http.StatusUnauthorized, Reason: "InvalidPassword", Message: "invalid password"}
	// ErrBadRequest   = &ErrorX{Code: http.StatusBadRequest, Reason: "BadRequest", Message: "bad request"}
	// ErrUnauthorized = &ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthorized", Message: "unauthorized"}
	// ErrForbidden    = &ErrorX{Code: http.StatusForbidden, Reason: "Forbidden", Message: "forbidden"}
)
