package store

import (
	"context"

	"github.com/quasimodo7614/nirvanatest/pkg/types"
)

const DefaultCountNum  = 10

type MsgStore interface {
	Add(ctx context.Context, obj types.Message) error
	Upd(ctx context.Context, obj types.Message) error
	List(ctx context.Context,count int) ([]types.Message, error)
	Get(ctx context.Context, id int) (types.Message, error)
	Del(ctx context.Context, id int) error
}
