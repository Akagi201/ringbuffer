// Package ringbuffer provides a simple implementation of a fixed size ring buffer.
package ringbuffer

// RingBuffer is a fixed size ring buffer.
type RingBuffer struct {
	buf    []interface{}
	size   int // capacity size
	length int // element count
	reader int // next position to read
	writer int // next position to write
}

// NewRingBuffer creates a new buffer of a given size.
func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		return nil
	}
	return &RingBuffer{
		buf:    make([]interface{}, size),
		size:   size,
		length: 0,
		reader: 0,
		writer: 0,
	}
}

// Capacity returns the current capacity of the ring buffer.
func (rb *RingBuffer) Capacity() int {
	return rb.size
}

// Length returns the element counts inside the ring buffer.
func (rb *RingBuffer) Length() int {
	return rb.length
}

func (rb *RingBuffer) mod(i int) int {
	return i % rb.size
}

// IsFull checks if the ring buffer is full
func (rb *RingBuffer) IsFull() bool {
	return rb.length == rb.size
}

// IsEmpty checks if the ring buffer is empty
func (rb *RingBuffer) IsEmpty() bool {
	return rb.length == 0
}

// Write an element into the ring buffer
func (rb *RingBuffer) Write(v interface{}) {
	if rb.IsFull() {
		return
	}
	rb.buf[rb.writer] = v
	rb.writer = rb.mod(rb.writer + 1)
	rb.length++
}

// WriteAt write an element at index i, i could be negative
func (rb *RingBuffer) WriteAt(i int, v interface{}) {
	i = rb.mod(i)
	rb.buf[i] = v
}

func (rb *RingBuffer) seekReader(delta int) {
	rb.reader = rb.mod(rb.reader + delta)
}

// Read read an element from the ring buffer, nil if empty
func (rb *RingBuffer) Read() interface{} {
	if rb.IsEmpty() {
		return nil
	}
	val := rb.buf[rb.reader]
	rb.seekReader(1)
	rb.length--
	return val
}

// ReadAt read an element at index i, i could be negative
func (rb *RingBuffer) ReadAt(i int) interface{} {
	if rb.IsEmpty() {
		return nil
	}
	i = rb.mod(i)
	return rb.buf[i]
}

// Peek peek the reader element and not affect the index
func (rb *RingBuffer) Peek() interface{} {
	if rb.IsEmpty() {
		return nil
	}
	return rb.buf[rb.reader]
}

// Clear clear the ring buffer
func (rb *RingBuffer) Clear() {
	rb.buf = make([]interface{}, rb.size)
	rb.length = 0
	rb.writer = 0
	rb.reader = 0
}
