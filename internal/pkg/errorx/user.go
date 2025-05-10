package errorx

import "net/http"

var (
	ErrUserNameInvalid = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "InvalidArgument.UserNameInvalid",
		Message: "Invalid username: Username must consist of letters, digits, and underscores only, and its length must be between 3 and 20 characters.",
	}

	ErrPasswordInvalid = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "InvalidArgument.PasswordInvalid",
		Message: "Password is incorrect.",
	}

	ErrUserAlreadyExists = &ErrorX{Code: http.StatusBadRequest, Reason: "AlreadyExist.UserAlreadyExists", Message: "User already exists."}
	ErrUserNotFound      = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound.UserNotFound", Message: "User not found."}
)
