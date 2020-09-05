package structures

import (
	"container/list"
	"sync"
)

// QueueWithSync is the queue interface
type QueueWithSync struct {
	queue *list.List
	mutex sync.Mutex
}

// Enqueue add a new element in to the queue
func (f *QueueWithSync) Enqueue(element interface{}) {
	f.mutex.Lock()
	f.queue.PushBack(element)
	f.mutex.Unlock()
}

// Dequeue remove an element from the queue
func (f *QueueWithSync) Dequeue() interface{} {
	f.mutex.Lock()
	last := f.queue.Back()
	defer f.mutex.Unlock()
	if last == nil {
		return nil
	}
	return f.queue.Remove(last)
}

// NewQueue creates a new instance of QueueWithSync
func NewQueue() *QueueWithSync {
	return &QueueWithSync{
		queue: list.New(),
	}
}
