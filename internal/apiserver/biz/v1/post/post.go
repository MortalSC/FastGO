package post

import (
	"context"

	"github.com/MortalSC/FastGO/internal/apiserver/model"
	"github.com/MortalSC/FastGO/internal/apiserver/store"
	"github.com/MortalSC/FastGO/internal/commonpkg/where"
	"github.com/MortalSC/FastGO/internal/pkg/contextx"
	"github.com/MortalSC/FastGO/internal/pkg/conversion"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/jinzhu/copier"
)

type PostBiz interface {
	Create(ctx context.Context, req *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error)
	Update(ctx context.Context, req *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error)
	Delete(ctx context.Context, req *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error)
	Get(ctx context.Context, req *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error)
	List(ctx context.Context, req *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error)

	PostExpansion
}

// PostExpansion is an interface that defines additional methods for the PostBiz
type PostExpansion interface{}

type postBiz struct {
	store store.IStore
}

var _ PostBiz = (*postBiz)(nil)

func New(store store.IStore) *postBiz {
	return &postBiz{
		store: store,
	}
}

func (b *postBiz) Create(ctx context.Context, req *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error) {
	var postM model.Post
	_ = copier.Copy(&postM, req)
	postM.UserID = contextx.UserID(ctx)

	if err := b.store.Post().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &apiv1.CreatePostResponse{
		PostID: postM.PostID,
	}, nil
}

func (b *postBiz) Update(ctx context.Context, req *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error) {
	whr := where.F("user_id", contextx.UserID(ctx), "post_id", req.PostID)
	postM, err := b.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		postM.Title = *req.Title
	}

	if req.Content != nil {
		postM.Content = *req.Content
	}

	if err := b.store.Post().Update(ctx, postM); err != nil {
		return nil, err
	}

	return &apiv1.UpdatePostResponse{}, nil
}

func (b *postBiz) Delete(ctx context.Context, req *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error) {
	whr := where.F("user_id", contextx.UserID(ctx), "post_id", req.PostID)
	if err := b.store.Post().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &apiv1.DeletePostResponse{}, nil
}

func (b *postBiz) Get(ctx context.Context, req *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error) {
	whr := where.F("user_id", contextx.UserID(ctx), "post_id", req.PostID)
	postM, err := b.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	return &apiv1.GetPostResponse{
		Post: conversion.PostModelToPostV1(postM),
	}, nil
}

func (b *postBiz) List(ctx context.Context, req *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error) {
	whr := where.F("user_id", contextx.UserID(ctx)).P(int(req.Offset), int(req.Limit))
	if req.Title != nil {
		whr = whr.Q("title like ?", "%"+*req.Title+"%")
	}

	count, postList, err := b.store.Post().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	posts := make([]*apiv1.Post, 0, len(postList))
	for _, post := range postList {
		posts = append(posts, conversion.PostModelToPostV1(post))
	}

	return &apiv1.ListPostResponse{
		Total: count,
		Posts: posts,
	}, nil
}
