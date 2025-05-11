package handler

import (
	"fmt"
	"log/slog"

	"github.com/MortalSC/FastGO/internal/pkg/core"
	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	v1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	slog.Info("Login function called")

	var req v1.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateLoginRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Login(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	slog.Info("Refresh token function called")

	var req v1.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateRefreshTokenRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().RefreshToken(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	slog.Info("Change password function called")

	var req v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateChangePasswordRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().ChangePassword(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called")

	var req v1.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateCreateUserRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Create(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called")

	var req v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdateUserRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Update(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called")

	var req v1.DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateDeleteUserRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called")
	fmt.Println("Get user function called")

	userID := c.Param("user_id")
	if userID == "" {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage("user_id is required"))
		return
	}
	fmt.Printf("UserID: %s, Type: %T\n", userID, userID)

	var req v1.GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	fmt.Println("Get user function called 2")

	if err := h.val.ValidateGetUserRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Get(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) ListUsers(c *gin.Context) {
	slog.Info("List users function called")

	var req v1.ListUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateListUserRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().List(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
