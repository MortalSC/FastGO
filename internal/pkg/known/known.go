package known

const (
	// XRequestID is the header key for request ID
	XRequestID = "x-request-id"

	// XUserID is the header key for user ID
	XUserID = "x-user-id"
)

const (
	// MaxErrGroupConcurrency is the maximum concurrency for errgroup
	// It is used to limit the number of goroutines running simultaneously in the errgroup, thereby preventing resource exhaustion and enhancing the stability of the program.
	// The size of this value can be adjusted according to the scene requirements.
	MaxErrGroupConcurrency = 1000
)
