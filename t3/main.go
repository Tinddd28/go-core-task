package t3

import (
	"sync"
)

type StringIntMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	copyMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[key]
	return val, ok
}
