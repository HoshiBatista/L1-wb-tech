package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu      sync.RWMutex
	safeMap map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{safeMap: make(map[string]interface{})}
}

func (m *SafeMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.safeMap[key] = value
}

func (m *SafeMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.safeMap[key]

	return value, exists
}

func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.safeMap, key)
}

func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.safeMap)
}

func (m *SafeMap) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]string, 0, len(m.safeMap))

	for k := range m.safeMap {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)

			safeMap.Set(key, value)

			fmt.Printf("Горутина %d записала: %s -> %s\n", id, key, value)
		}(i)
	}

	wg.Wait()

	fmt.Println("\nСодержимое map:")

	for _, key := range safeMap.Keys() {
		if value, exists := safeMap.Get(key); exists {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	fmt.Printf("Всего элементов: %d\n", safeMap.Len())
}
