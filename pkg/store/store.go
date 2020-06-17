package store

import "context"

type Store interface {
	Add(ctx context.Context, obj interface{}) (interface{}, error)
	Upd(ctx context.Context, obj interface{}) (interface{}, error)
	List(ctx context.Context) (interface{}, error)
	Get(ctx context.Context, id interface{}) (interface{}, error)
	Del(ctx context.Context, id interface{}) error
}
