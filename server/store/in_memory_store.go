package store

import (
	"errors"
	"fmt"
)

func NewInMemoryStore(storeID string) *MemoryStore {
	return &MemoryStore{
		storeID: storeID,
		keyValues: make(map[string]string),
	}
}

type MemoryStore struct {
	storeID string
	keyValues map[string]string
}

func (ms *MemoryStore) Set(key, value string) error {
	ms.keyValues[key] = value
	return nil
}

func (ms *MemoryStore) Get(key string) (string, error) {
	ret, ok := ms.keyValues[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("No value for the key: %s", key))
	}
	return ret, nil
}

func (ms *MemoryStore) Delete(key string) error {
	delete(ms.keyValues, ms.keyValues[key])
	return nil
}