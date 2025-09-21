package main

import "errors"

type CircularQueue[T any] struct {
	data  []T
	front int
	rear  int
	size  int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		data:  make([]T, capacity),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (q *CircularQueue[T]) Enqueue(item T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.rear] = item
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	item := q.data[q.front]
	q.data[q.front] = zero // clear the slot
	q.front = (q.front + 1) % len(q.data)
	q.size--
	return item, nil
}

func (q *CircularQueue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.data[q.front], nil
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.size == len(q.data)
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}

func (q *CircularQueue[T]) Capacity() int {
	return len(q.data)
}

// Example usage
func exampleUsage() {
	// String queue
	strQueue := NewCircularQueue[string](3)
	strQueue.Enqueue("first")
	strQueue.Enqueue("second")
	strQueue.Enqueue("third")

	item, _ := strQueue.Dequeue()
	println(item) // prints "first"

	// Integer queue
	intQueue := NewCircularQueue[int](5)
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)

	val, _ := intQueue.Peek()
	println(val) // prints 10
}
