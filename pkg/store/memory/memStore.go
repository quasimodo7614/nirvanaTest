package memory

import (
	"context"
	"errors"
	"nirvanaTest/pkg/types"
	"sync"
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

func (s *MemStore) Add(ctx context.Context, obj interface{}) (interface{}, error) {
	m, ok := obj.(types.Message)
	if !ok {
		return nil, errors.New("type is not message")
	}
	if _, ok2 := s.Valm[m.ID]; ok2 {
		return nil, errors.New("already exist")
	}
	s.Lock()
	s.Valm[m.ID] = m
	s.Unlock()
	return m, nil
}

func (s *MemStore) Del(ctx context.Context, id interface{}) error {
	s.Lock()
	delete(s.Valm, id.(int))
	s.Unlock()
	return nil
}

func (s *MemStore) Upd(ctx context.Context, obj interface{}) (interface{}, error) {
	m, ok := obj.(types.Message)
	if !ok {
		return nil, errors.New("type is not message")
	}
	if _, ok2 := s.Valm[m.ID]; !ok2 {
		return nil, errors.New("obj not exist")
	}
	s.Lock()
	s.Valm[m.ID] = m
	s.Unlock()
	return m, nil
}

func (s *MemStore) Get(ctx context.Context, id interface{}) (interface{}, error) {
	s.Lock()
	m, ok := s.Valm[id.(int)]
	s.Unlock()
	if !ok {
		return nil, errors.New("not found")
	}
	return m, nil
}
func (s *MemStore) List(ctx context.Context) (interface{}, error) {
	var ret []types.Message
	s.Lock()
	for _, v := range s.Valm {
		ret = append(ret, v)
	}
	s.Unlock()
	if len(ret) == 0 {
		return nil, errors.New("no result found")
	}
	return ret, nil
}
