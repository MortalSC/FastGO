package errorx

import (
	"errors"
	"fmt"
)

type ErrorX struct {
	Code    int    `json:"code,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func New(code int, reason string, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

func (e *ErrorX) Error() string {
	return fmt.Sprintf("code: %d, reason: %s, message: %s", e.Code, e.Reason, e.Message)
}

func (e *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func FromError(err error) *ErrorX {
	if err == nil {
		return nil
	}
	if errx := new(ErrorX); errors.As(err, &errx) {
		return errx
	}
	return New(ErrInternal.Code, ErrInternal.Reason, err.Error())
}
