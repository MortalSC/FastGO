package handler

import (
	"log/slog"

	"github.com/MortalSC/FastGO/internal/pkg/core"
	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	v1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	slog.Info("Create post function called")

	var req v1.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateCreatePostRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Create(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	slog.Info("Update post function called")

	var req v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdatePostRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Update(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) DeletePost(c *gin.Context) {
	slog.Info("Delete post function called")

	var req v1.DeletePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateDeletePostRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Delete(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) GetPost(c *gin.Context) {
	slog.Info("Get post function called")

	var req v1.GetPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateGetPostRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Get(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) ListPosts(c *gin.Context) {
	slog.Info("List posts function called")

	var req v1.ListPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrBind)
		return
	}

	if err := h.val.ValidateListPostRequest(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, nil, errorx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().List(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
