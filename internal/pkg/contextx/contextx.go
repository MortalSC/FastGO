package contextx

import "context"

type (
	// requestIDKey defines the context key of the requestID.
	requestIDKey struct{}

	// userIDKey defines the context key of the userID.
	userIDkey struct{}
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

// WithUserID sets the user ID in the context
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDkey{}, userID)
}

// UserID gets the user ID from the context
func UserID(ctx context.Context) string {
	userID, _ := ctx.Value(userIDkey{}).(string)
	return userID
}
