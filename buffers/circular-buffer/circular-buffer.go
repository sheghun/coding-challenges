package main

import (
	"errors"
	"io"
)

type CircularBuffer struct {
	data  []byte
	read  int
	write int
	size  int // tracks how much data is currently in buffer
}

func NewCircularBuffer(capacity int) *CircularBuffer {
	return &CircularBuffer{
		data: make([]byte, capacity),
	}
}

func (c *CircularBuffer) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	capacity := len(c.data)
	available := capacity - c.size

	if available == 0 {
		return 0, errors.New("buffer full")
	}

	// Write only what fits
	toWrite := len(p)
	if toWrite > available {
		toWrite = available
	}

	written := 0
	for i := 0; i < toWrite; i++ {
		c.data[c.write] = p[i]
		c.write = (c.write + 1) % capacity
		written++
		c.size++
	}

	return written, nil
}

func (c *CircularBuffer) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if c.size == 0 {
		return 0, io.EOF
	}

	capacity := len(c.data)

	// Read only what's available and what fits in p
	toRead := len(p)
	if toRead > c.size {
		toRead = c.size
	}

	read := 0
	for i := 0; i < toRead; i++ {
		p[i] = c.data[c.read]
		c.read = (c.read + 1) % capacity
		read++
		c.size--
	}

	return read, nil
}

func main() {
	// Example usage
	buf := NewCircularBuffer(5)
	buf.Write([]byte("hello"))

	data := make([]byte, 3)
	n, _ := buf.Read(data)
	println(string(data[:n])) // prints "hel"
}
