package data

import (
	"errors"
	"sync"
)

// KeyValueStore menyimpan data key-value secara in-memory.
type KeyValueStore struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewKeyValueStore membuat instance baru dari KeyValueStore.
func NewKeyValueStore(initialData map[string]string) *KeyValueStore {
	return &KeyValueStore{
		store: initialData,
	}
}

// Get mengambil nilai dari key yang diberikan.
func (kv *KeyValueStore) Get(key string) (string, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	val, exists := kv.store[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return val, nil
}

// Set menetapkan nilai untuk key yang diberikan.
func (kv *KeyValueStore) Set(key string, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = value
}

// Delete menghapus key yang diberikan.
func (kv *KeyValueStore) Delete(key string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, exists := kv.store[key]; !exists {
		return errors.New("key not found")
	}
	delete(kv.store, key)
	return nil
}

// GetAll mengembalikan semua key-value yang tersimpan.
func (kv *KeyValueStore) GetAll() map[string]string {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	copy := make(map[string]string)
	for k, v := range kv.store {
		copy[k] = v
	}
	return copy
}
