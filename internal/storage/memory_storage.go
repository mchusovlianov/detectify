package storage

import (
	"fmt"
	"sync"
)

type MemoryStorage struct {
	rw      *sync.RWMutex
	storage map[string]interface{}
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		rw:      &sync.RWMutex{},
		storage: make(map[string]interface{}),
	}
}

func (m MemoryStorage) Set(uuid string, result interface{}) error {
	m.rw.Lock()
	m.storage[uuid] = result
	m.rw.Unlock()

	return nil
}

func (m *MemoryStorage) Get(uuid string) (interface{}, error) {
	m.rw.RLock()
	data, ok := m.storage[uuid]
	if !ok {
		m.rw.RUnlock()
		return nil, fmt.Errorf("Key %v not found", uuid)
	}

	m.rw.RUnlock()
	return data, nil
}
