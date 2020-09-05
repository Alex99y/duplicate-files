package structures

import (
	"sync"
)

// MapWithSync have the map with mutex
type MapWithSync struct {
	mutex sync.RWMutex
	smap  map[string][]string
}

// AddElement add an element to the map
func (c *MapWithSync) AddElement(key string, value string) {
	c.mutex.Lock()
	c.smap[key] = append(c.smap[key], value)
	c.mutex.Unlock()
}

// GetMap returns the map
func (c *MapWithSync) GetMap() map[string][]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.smap
}

// NewMap creates a new MapWithSync instance
func NewMap() *MapWithSync {
	return &MapWithSync{
		smap: make(map[string][]string),
	}
}
