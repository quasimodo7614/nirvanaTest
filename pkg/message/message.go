package message

import (
	"context"
	"errors"

	"github.com/quasimodo7614/nirvanatest/pkg/store/mongodb"
	"github.com/quasimodo7614/nirvanatest/pkg/types"
)

//var db = memory.MemDb
var db = mongodb.MS

// ListMessages returns all messages.
func ListMessages(ctx context.Context, count int) ([]types.Message, error) {
	return db.List(ctx, count)
}

// GetMessage returns a message by id.
func GetMessage(ctx context.Context, id int) (types.Message, error) {
	return db.Get(ctx, id)
}

// CreateMessage returns a message by id.
func CreateMessage(ctx context.Context, message types.Message) (types.Message, error) {
	// judge weather this id exist before
	_, err := db.Get(ctx, message.ID)
	if err == nil {
		return types.Message{}, errors.New("already exist")
	}

	return message, db.Add(ctx, message)
}

// UpdateMessage returns a message by id.
func UpdateMessage(ctx context.Context, id int, message types.Message) (types.Message, error) {
	// judge this id exist or not
	_, err := db.Get(ctx, id)
	if err != nil {
		return types.Message{}, err
	}
	return message, db.Upd(ctx, message)
}

// DeleteMessage returns a message by id.
func DeleteMessage(ctx context.Context, id int) error {
	return db.Del(ctx, id)
}
