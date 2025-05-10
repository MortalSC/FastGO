package errorx

import "net/http"

var ErrPostNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
