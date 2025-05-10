package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/MortalSC/FastGO/internal/apiserver/model"
	"github.com/MortalSC/FastGO/internal/commonpkg/where"
	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, obj *model.User) error
	Update(ctx context.Context, obj *model.User) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.User, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.User, error)

	UserExpansion
}

// UserExpansion is an interface that defines additional methods for the UserStore
type UserExpansion interface {
	// Add any additional methods for the UserStore here
	// For example:
	// GetByEmail(ctx context.Context, email string) (*model.User, error)
}

// userStore is a struct that implements the UserStore interface
type userStore struct {
	// db instance
	store *datastore
}

var _ UserStore = (*userStore)(nil)

func newUserStore(store *datastore) *userStore {
	return &userStore{
		store: store,
	}
}

func (s *userStore) Create(ctx context.Context, obj *model.User) error {
	if err := s.store.DB(ctx).Create(obj).Error; err != nil {
		slog.Error("Failed to insert user into database", "err", err, "user", obj)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Update(ctx context.Context, obj *model.User) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update user in database", "err", err, "user", obj)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx, opts).Delete(new(model.User)).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete user from database", "err", err, "opts", opts)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Get(ctx context.Context, opts *where.Options) (*model.User, error) {
	var obj model.User
	if err := s.store.DB(ctx, opts).First(&obj).Error; err != nil {
		slog.Error("Failed to get user from database", "err", err, "opts", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		return nil, errorx.ErrDBRead.WithMessage(err.Error())
	}
	return &obj, nil
}

func (s *userStore) List(ctx context.Context, opts *where.Options) (int64, []*model.User, error) {

	var (
		total int64
		users []*model.User
	)

	baseDB := s.store.DB(ctx, opts).Model(&model.User{})

	if err := baseDB.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		slog.Error("Failed to count users", "err", err, "conditions", opts)
		return 0, nil, errorx.ErrDBRead.WithMessage(err.Error())
	}

	if err := baseDB.Order("id desc").Find(&users).Error; err != nil {
		slog.Error("Failed to list users", "err", err, "conditions", opts)
		return 0, nil, errorx.ErrDBRead.WithMessage(err.Error())
	}

	return total, users, nil
}
