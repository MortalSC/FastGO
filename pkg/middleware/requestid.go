package middleware

import (
	"github.com/MortalSC/FastGO/pkg/contextx"
	"github.com/MortalSC/FastGO/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get 'x-request-id' from header， if it does exist, a new uuid will be generated
		requestID := c.Request.Header.Get(known.XRequestID)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Set 'x-request-id' to context.Context
		ctx := contextx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// Set 'x-request-id' to response header
		c.Writer.Header().Set(known.XRequestID, requestID)

		c.Next()
	}
}
