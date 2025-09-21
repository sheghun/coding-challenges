package main

import (
	"errors"
	"sync"
)

type LRUCache[T any] struct {
	data     map[string]*Node[T]
	capacity int
	size     int
	head     *Node[T]
	tail     *Node[T]
	mutex    sync.RWMutex
}

type Node[T any] struct {
	Key   string
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

func main() {
}

func NewLRUCache[T any](capacity int) *LRUCache[T] {
	return &LRUCache[T]{
		data:     make(map[string]*Node[T]),
		capacity: capacity,
		size:     0,
	}
}

func (l *LRUCache[T]) Get(key string) (T, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var t T
	// get the first one
	curr, ok := l.data[key]
	if !ok {
		return t, errors.New("key not found")
	}

	// move to head and return value
	l.moveToHead(curr)
	return curr.Value, nil
}

// Contains checks if a key exists without moving it to head (read-only operation)
func (l *LRUCache[T]) Contains(key string) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	_, exists := l.data[key]
	return exists
}

func (l *LRUCache[T]) moveToHead(node *Node[T]) {
	// if already at head, nothing to do
	if node == l.head {
		return
	}

	// remove from current position
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// update tail if this was the tail
	if node == l.tail {
		l.tail = node.Prev
	}

	// move to head
	node.Prev = nil
	node.Next = l.head
	l.head.Prev = node
	l.head = node
}

func (l *LRUCache[T]) Add(key string, value T) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// check if key already exists
	if existingNode, exists := l.data[key]; exists {
		// update value and move to head
		existingNode.Value = value
		l.moveToHead(existingNode)
		return nil
	}

	// create new node for new key
	node := &Node[T]{Value: value, Key: key}
	l.data[key] = node

	if l.size == 0 {
		l.head = node
		l.tail = node
		l.size++
		return nil
	}

	// add to the head
	l.head.Prev = node
	node.Next = l.head
	l.head = node

	l.size++

	l.Resize()

	return nil
}

func (l *LRUCache[T]) Resize() {
	if l.size <= l.capacity {
		return
	}

	// remove from map
	delete(l.data, l.tail.Key)

	// handle single node case
	if l.tail.Prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		tmp := l.tail.Prev
		tmp.Next = nil
		l.tail = tmp
	}
	l.size--
}
