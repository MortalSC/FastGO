package core

import (
	"context"
	"net/http"

	"github.com/MortalSC/FastGO/internal/commonpkg/errorx"
	"github.com/gin-gonic/gin"
)

// Validator is a type of validation function used to validate bound data structures.
type Validator[T any] func(context.Context, *T) error

// Binder defines the type of the binding function, which is used to bind the request data to the corresponding structure.
type Binder func(any) error

// Handler is a type of processing function used to handle data that has been bound and verified.
type Handler[T any, R any] func(ctx context.Context, req *T) (R, error)

// ErrorResponse defines the structure of the error response
// Used to return a unified formatting error message when an error occurs in the API request.
type ErrorResponse struct {
	Reason   string            `json:"reason",omitempty`
	Message  string            `json:"message",omitempty`
	Metadata map[string]string `json:"metadata",omitempty`
}

// HandleJSONRequest is a shortcut function for handling JSON requests.
func HandleJSONRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindJSON, handler, validators...)
}

// HandleQueryRequest is a shortcut function for handling Query parameter requests.
func HandleQueryRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindQuery, handler, validators...)
}

// HandleUriRequest is a shortcut function for handling URI requests.
func HandleUriRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindUri, handler, validators...)
}

// HandleRequest is a common request processing function.
// Be responsible for binding request data, performing verification, and calling the actual business processing logic functions.
func HandleRequest[T any, R any](c *gin.Context, binder Binder, handler Handler[T, R], validators ...Validator[T]) {
	var request T

	// Bind and validate the request data
	if err := ReadRequest(c, &request, binder, validators...); err != nil {
		WriteResponse(c, nil, err)
		return
	}

	// Call the actual business processing logic function
	response, err := handler(c.Request.Context(), &request)
	WriteResponse(c, response, err)
}

func ShouldBindJSON[T any](c *gin.Context, req *T, validatros ...Validator[T]) error {
	return ReadRequest(c, req, c.ShouldBindJSON, validatros...)
}

func ShouldBindQuery[T any](c *gin.Context, req *T, validatros ...Validator[T]) error {
	return ReadRequest(c, req, c.ShouldBindQuery, validatros...)
}

func ShouldBindUri[T any](c *gin.Context, req *T, validatros ...Validator[T]) error {
	return ReadRequest(c, req, c.ShouldBindUri, validatros...)
}

// ReadRequest is a common utility function used for binding and verifying request data.
// - It is responsible for calling the binding function to bind the request data.
// - If the target type implements the Default interface, its Default method will be called to set the default value.
// - Finally, execute the incoming validator to verify the data.
func ReadRequest[T any](c *gin.Context, req *T, binder Binder, validators ...Validator[T]) error {
	// Call the binding function to bind the request data
	if err := binder(req); err != nil {
		return errorx.ErrBind.WithMessage(err.Error())
	}

	if defaulter, ok := any(req).(interface{ Default() }); ok {
		defaulter.Default()
	}

	for _, validate := range validators {
		if validate == nil {
			continue
		}
		if err := validate(c.Request.Context(), req); err != nil {
			return err
		}
	}
	return nil
}

func WriteResponse(c *gin.Context, data any, err error) {
	if err != nil {
		errx := errorx.FromError(err)
		c.JSON(errx.Code, ErrorResponse{
			Reason:   errx.Reason,
			Message:  errx.Message,
			Metadata: errx.Metadata,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
