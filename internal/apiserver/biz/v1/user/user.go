package user

import (
	"context"
	"log/slog"
	"sync"

	"github.com/MortalSC/FastGO/internal/apiserver/model"
	"github.com/MortalSC/FastGO/internal/apiserver/store"
	"github.com/MortalSC/FastGO/internal/commonpkg/where"
	"github.com/MortalSC/FastGO/internal/pkg/contextx"
	"github.com/MortalSC/FastGO/internal/pkg/conversion"
	"github.com/MortalSC/FastGO/internal/pkg/known"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/errgroup"
)

type UserBiz interface {
	Create(ctx context.Context, req *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error)
	Update(ctx context.Context, req *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error)
	Delete(ctx context.Context, req *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error)
	Get(ctx context.Context, req *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error)
	List(ctx context.Context, req *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error)

	UserExpansion
}

// UserExpansion is an interface that defines additional methods for the UserBiz
type UserExpansion interface{}

type userBiz struct {
	store store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(store store.IStore) *userBiz {
	return &userBiz{
		store: store,
	}
}

func (b *userBiz) Create(ctx context.Context, req *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error) {
	var userM model.User
	_ = copier.Copy(&userM, req)

	if err := b.store.User().Create(ctx, &userM); err != nil {
		return nil, err
	}

	return &apiv1.CreateUserResponse{
		UserID: userM.UserID,
	}, nil
}

func (b *userBiz) Update(ctx context.Context, req *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("user_id", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	if req.Username != nil {
		userM.Username = *req.Username
	}
	if req.Email != nil {
		userM.Email = *req.Email
	}
	if req.Nickname != nil {
		userM.Nickname = *req.Nickname
	}
	if req.Phone != nil {
		userM.Phone = *req.Phone
	}

	if err := b.store.User().Update(ctx, userM); err != nil {
		return nil, err
	}

	return &apiv1.UpdateUserResponse{}, nil
}

func (b *userBiz) Delete(ctx context.Context, req *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error) {
	if err := b.store.User().Delete(ctx, where.F("user_id", contextx.UserID(ctx))); err != nil {
		return nil, err
	}
	return &apiv1.DeleteUserResponse{}, nil
}

func (b *userBiz) Get(ctx context.Context, req *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("user_id", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	return &apiv1.GetUserResponse{
		User: conversion.UserModelToUserV1(userM),
	}, nil
}

func (b *userBiz) List(ctx context.Context, req *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error) {
	whr := where.P(int(req.Offset), int(req.Limit))
	count, userList, err := b.store.User().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	eg.SetLimit(known.MaxErrGroupConcurrency)

	for _, user := range userList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				count, _, err := b.store.User().List(ctx, where.F("user_id", user.UserID))
				if err != nil {
					return err
				}

				converted := conversion.UserModelToUserV1(user)
				converted.PostCount = count
				m.Store(user.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		slog.ErrorContext(ctx, "Failed to wait all function calls returned", "err", err)
		return nil, err
	}

	users := make([]*apiv1.User, 0, len(userList))
	for _, item := range userList {
		user, _ := m.Load(item.ID)
		user = append(users, user.(*apiv1.User))
	}

	slog.DebugContext(ctx, "Get users from backend storage", "count", len(users))

	return &apiv1.ListUserResponse{
		Total: count,
		Users: users,
	}, nil
}
