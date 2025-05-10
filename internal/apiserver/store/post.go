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

type PostStore interface {
	Create(ctx context.Context, obj *model.Post) error
	Update(ctx context.Context, obj *model.Post) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.Post, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.Post, error)

	PostExpansion
}

// PostExpansion is an interface that defines additional methods for the PostStore
type PostExpansion interface {
	// Add any additional methods for the PostStore here
	// For example:
	// GetByEmail(ctx context.Context, email string) (*model.Post, error)
}

// postStore is a struct that implements the PostStore interface
type postStore struct {
	// db instance
	store *datastore
}

var _ PostStore = (*postStore)(nil)

func newPostStore(store *datastore) *postStore {
	return &postStore{
		store: store,
	}
}

func (s *postStore) Create(ctx context.Context, obj *model.Post) error {
	if err := s.store.DB(ctx).Create(obj).Error; err != nil {
		slog.Error("Failed to insert post into database", "err", err, "post", obj)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *postStore) Update(ctx context.Context, obj *model.Post) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update post in database", "err", err, "post", obj)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *postStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx, opts).Delete(&model.Post{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete post from database", "err", err, "opts", opts)
		return errorx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *postStore) Get(ctx context.Context, opts *where.Options) (*model.Post, error) {
	var post model.Post
	err := s.store.DB(ctx, opts).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrPostNotFound.WithMessage(err.Error())
		}
		slog.Error("Failed to get post from database", "err", err, "opts", opts)
		return nil, errorx.ErrDBRead.WithMessage(err.Error())
	}
	return &post, nil
}

func (s *postStore) List(ctx context.Context, opts *where.Options) (int64, []*model.Post, error) {

	var (
		total int64
		posts []*model.Post
	)

	baseDB := s.store.DB(ctx, opts).Model(&model.Post{})

	if err := baseDB.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		slog.Error("Failed to count users", "err", err, "conditions", opts)
		return 0, nil, errorx.ErrDBRead.WithMessage(err.Error())
	}

	if err := baseDB.Order("id desc").Find(&posts).Error; err != nil {
		slog.Error("Failed to list users", "err", err, "conditions", opts)
		return 0, nil, errorx.ErrDBRead.WithMessage(err.Error())
	}

	return total, posts, nil
}
