package middleware

import (
	"github.com/MortalSC/FastGO/internal/commonpkg/core"
	"github.com/MortalSC/FastGO/internal/pkg/contextx"
	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	"github.com/MortalSC/FastGO/pkg/token"
	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, nil, errorx.ErrTokenInvalid.WithMessage(err.Error()))
			c.Abort()
			return
		}

		ctx := contextx.WithUserID(c.Request.Context(), userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
