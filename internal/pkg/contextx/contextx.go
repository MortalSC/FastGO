package contextx

import "context"

type (
	// ContextKey is a type for context keys
	requestIDKey struct{}
)

// WithRequestID sets the request ID in the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID gets the request ID from the context
func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}
