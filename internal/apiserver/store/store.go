package store

import (
	"context"
	"sync"

	"github.com/MortalSC/FastGO/internal/commonpkg/where"
	"gorm.io/gorm"
)

var (
	once sync.Once
	// Store is the global store instance
	Store *datastore
)

type IStore interface {
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error) error

	User() UserStore
	Post() PostStore
}

// transactionKey is the key used to store the transaction in the context
type transactionKey struct{}

type datastore struct {
	gormDBCore *gorm.DB

	// Support adding other database instances
}

// Ensure that the IStore interface is implemented by the datastore struct
var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	// Ensure the Store is initialized only once
	once.Do(func() {
		Store = &datastore{db}
	})
	return Store
}

// DB filters database instances based on the incoming conditions (wheres).
// If no conditions are passed in, the database instance (transaction instance or core database instance) in the context is returned.
func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.gormDBCore

	// get the transaction from the context
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// loop through the wheres and apply them to the db instance
	for _, whr := range wheres {
		db = whr.Where(db)
	}

	return db
}

// TX returns a new transaction instance.
// nolint: fatcontext
func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return store.gormDBCore.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// Users returns an instance that implements the UserStore interface
func (store *datastore) User() UserStore {
	return newUserStore(store)
}

// Posts returns an instance that implements the PostStore interface
func (store *datastore) Post() PostStore {
	return newPostStore(store)
}
