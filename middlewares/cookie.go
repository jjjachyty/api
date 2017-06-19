package middlewares

import (
	"net/http"
	"sync"
	"time"
)

type MemoryCookie struct {
	cook   map[string]*http.Cookie
	MaxAge time.Duration
	lock   sync.RWMutex
}

func NewMemoryCooke(maxAge time.Duration) *MemoryCookie {
	return &MemoryCookie{
		MaxAge: maxAge,
		cook:   make(map[string]*http.Cookie),
	}
}

func (store *MemoryCookie) Set(key, value string) error {
	store.lock.Lock()
	cook, ok := store.cook[key]
	if !ok {
		cook = new(http.Cookie)
		cook.Expires = time.Now().Add(store.MaxAge)

		cook.Value = value
	}
	store.cook[key] = cook

	store.lock.Unlock()
	return nil
}

func (store *MemoryCookie) Get(key string) (string, bool) {
	store.lock.Lock()
	defer store.lock.Unlock()
	cook, ok := store.cook[key]
	if !ok {
		return "", false
	}
	v := cook.Value
	return v, true
}

func (store *MemoryCookie) Del(key string) bool {
	store.lock.Lock()
	defer store.lock.Unlock()
	delete(store.cook, key)
	return true
}
