package message

import (
	"context"
	"github.com/quasimodo7614/nirvanatest/pkg/store/mongodb"
	"github.com/quasimodo7614/nirvanatest/pkg/types"
)

//var db = memory.MemDb
var db = mongodb.MS

// ListMessages returns all messages.
func ListMessages(ctx context.Context, count int) (interface{}, error) {
	all, err := db.List(ctx)
	return all, err
}

// GetMessage returns a message by id.
func GetMessage(ctx context.Context, id int) (interface{}, error) {
	return db.Get(ctx, id)
}

// CreateMessage returns a message by id.
func CreateMessage(ctx context.Context, message types.Message) (interface{}, error) {
	return db.Add(ctx, message)
}

// UpdateMessage returns a message by id.
func UpdateMessage(ctx context.Context, message types.Message) (interface{}, error) {
	return db.Upd(ctx, message)
}

// DeleteMessage returns a message by id.
func DeleteMessage(ctx context.Context, id int) error {
	return db.Del(ctx, id)
}
