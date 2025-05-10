package errorx

import (
	"errors"
	"fmt"
	"net/http"

	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

type ErrorX struct {
	Code     int               `json:"code",omitempty`
	Reason   string            `json:"reason",omitempty`
	Message  string            `json:"message",omitempty`
	Metadata map[string]string `json:"metadata",omitempty`
}

func New(code int, reason string, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: format,
	}
}

func (err *ErrorX) Error() string {
	return fmt.Sprintf("error: code: %d reason: %s, message: %s, metadata=%v", err.Code, err.Reason, err.Message, err.Metadata)
}

func (err *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func (err *ErrorX) WithMetadata(md map[string]string) *ErrorX {
	err.Metadata = md
	return err
}

func (err *ErrorX) KV(kvs ...string) *ErrorX {
	if err.Metadata == nil {
		err.Metadata = make(map[string]string)
	}

	for i := 0; i < len(kvs); i += 2 {
		if i+1 < len(kvs) {
			err.Metadata[kvs[i]] = kvs[i+1]
		}
	}
	return err
}

func (err *ErrorX) GRPCStatus() *status.Status {
	details := errdetails.ErrorInfo{
		Reason:   err.Reason,
		Metadata: err.Metadata,
	}
	s, _ := status.New(httpstatus.ToGRPCCode(err.Code), err.Message).WithDetails(&details)
	return s
}

func (err *ErrorX) WithRequestID(requestID string) *ErrorX {
	return err.KV("X-Request-ID", requestID)
}

func (err *ErrorX) Is(target error) bool {
	if errx := new(ErrorX); errors.As(target, &errx) {
		return errx.Code == err.Code && errx.Reason == err.Reason
	}
	return false
}

// Code 返回错误的 HTTP 代码.
func Code(err error) int {
	if err == nil {
		return http.StatusOK //nolint:mnd
	}
	return FromError(err).Code
}

// Reason 返回特定错误的原因.
func Reason(err error) string {
	if err == nil {
		return ErrInternal.Reason
	}
	return FromError(err).Reason
}

// FromError 尝试将一个通用的 error 转换为自定义的 *ErrorX 类型.
func FromError(err error) *ErrorX {
	// 如果传入的错误是 nil，则直接返回 nil，表示没有错误需要处理.
	if err == nil {
		return nil
	}

	// 检查传入的 error 是否已经是 ErrorX 类型的实例.
	// 如果错误可以通过 errors.As 转换为 *ErrorX 类型，则直接返回该实例.
	if errx := new(ErrorX); errors.As(err, &errx) {
		return errx
	}

	// gRPC 的 status.FromError 方法尝试将 error 转换为 gRPC 错误的 status 对象.
	// 如果 err 不能转换为 gRPC 错误（即不是 gRPC 的 status 错误），
	// 则返回一个带有默认值的 ErrorX，表示是一个未知类型的错误.
	gs, ok := status.FromError(err)
	if !ok {
		return New(ErrInternal.Code, ErrInternal.Reason, err.Error())
	}

	// 如果 err 是 gRPC 的错误类型，会成功返回一个 gRPC status 对象（gs）.
	// 使用 gRPC 状态中的错误代码和消息创建一个 ErrorX.
	ret := New(httpstatus.FromGRPCCode(gs.Code()), ErrInternal.Reason, gs.Message())

	// 遍历 gRPC 错误详情中的所有附加信息（Details）.
	for _, detail := range gs.Details() {
		if typed, ok := detail.(*errdetails.ErrorInfo); ok {
			ret.Reason = typed.Reason
			return ret.WithMetadata(typed.Metadata)
		}
	}

	return ret
}
