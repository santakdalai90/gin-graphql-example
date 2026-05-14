package storage

import "sync"

type LocalStorage struct {
	store sync.Map
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		store: sync.Map{},
	}
}

func (ls *LocalStorage) PutObject(key any, obj any) {
	ls.store.Store(key, obj)
}

func (ls *LocalStorage) GetObject(key any) (any, bool) {
	return ls.store.Load(key)
}

func (ls *LocalStorage) GetAll() []any {
	objects := make([]any, 0)
	ls.store.Range(func(k, v any) bool {
		_ = k
		objects = append(objects, v)
		return true
	})
	return objects
}
