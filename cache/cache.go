package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	// TODO: ваш код
	mu *sync.RWMutex
	cache map[string]string
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	// TODO: ваш код
	return &Cache{
		&sync.RWMutex{},
		make(map[string]string),
	}
	// panic("implement me")
}

func (c Cache) Set(k, v string) {
	// TODO implement me
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[k] = v
	// panic("implement me")
}

func (c Cache) Get(k string) (v string, ok bool) {
	// TODO implement me
	c.mu.RLock()
	defer c.mu.RUnlock()
	if val, ok := c.cache[k]; ok {
		return val, ok
	}
	return "", false
	// panic("implement me")
}
