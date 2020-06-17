package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/quasimodo7614/nirvanatest/pkg/store"
	"github.com/quasimodo7614/nirvanatest/pkg/types"
)

type MemStore struct {
	Valm map[int]types.Message
	sync.RWMutex
}

var MemDb *MemStore

func init() {
	MemDb = NewMemStore()
}

func NewMemStore() *MemStore {
	return &MemStore{
		Valm: make(map[int]types.Message),
	}
}

func (s *MemStore) Add(ctx context.Context, obj types.Message) (error) {
	s.Lock()
	s.Valm[obj.ID] = obj
	s.Unlock()
	return nil
}

func (s *MemStore) Del(ctx context.Context, id int) error {
	s.Lock()
	delete(s.Valm, id)
	s.Unlock()
	return nil
}

func (s *MemStore) Upd(ctx context.Context, obj types.Message) (error) {
	s.Lock()
	s.Valm[obj.ID] = obj
	s.Unlock()
	return nil
}

func (s *MemStore) Get(ctx context.Context, id int) (types.Message, error) {
	s.Lock()
	m, ok := s.Valm[id]
	s.Unlock()
	if !ok {
		return types.Message{}, errors.New("not found")
	}
	return m, nil
}

func (s *MemStore) List(ctx context.Context, count int) ([]types.Message, error) {
	if count == 0 {
		count = store.DefaultCountNum
	}
	var ret []types.Message
	s.Lock()
	cur := 1
	for _, v := range s.Valm {
		ret = append(ret, v)
		if cur > count {
			break
		}
		cur ++
	}
	s.Unlock()

	return ret, nil
}
