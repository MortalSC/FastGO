package core

import (
	"net/http"

	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteResponse(c *gin.Context, data any, err error) {
	if err != nil {
		errx := errorx.FromError(err) // error -> *errorx.ErrorX
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Reason:  errx.Reason,
			Message: errx.Message,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
