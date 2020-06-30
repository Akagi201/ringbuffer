package ringbuffer_test

import (
	"testing"

	"github.com/Akagi201/ringbuffer"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)

	rb := ringbuffer.NewRingBuffer(5)
	assert.True(rb.IsEmpty())
	assert.False(rb.IsFull())
	assert.Zero(rb.Length())
	assert.Equal(5, rb.Capacity())
}

func TestWriteRead(t *testing.T) {
	assert := assert.New(t)
	rb := ringbuffer.NewRingBuffer(5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}
	assert.Equal(5, rb.Length())
	assert.Equal(5, rb.Capacity())

	for i := 0; i < 5; i++ {
		val := rb.Read()
		v := val.(int)
		assert.Equal(i, v)
	}
}

func TestWriteFull(t *testing.T) {
	assert := assert.New(t)
	rb := ringbuffer.NewRingBuffer(5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}

	assert.True(rb.IsFull())
	// can not write
	for i := 5; i < 8; i++ {
		rb.Write(i)
	}

	for i := 0; i < 5; i++ {
		val := rb.Read()
		v := val.(int)
		assert.Equal(i, v)
	}
}

func TestReadEmpty(t *testing.T) {
	assert := assert.New(t)
	rb := ringbuffer.NewRingBuffer(5)
	assert.Nil(rb.Read())
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}
	for i := 0; i < 5; i++ {
		val := rb.Read()
		v := val.(int)
		assert.Equal(i, v)
	}
	assert.Nil(rb.Read())
}

func TestReuseBuffer(t *testing.T) {
	assert := assert.New(t)
	rb := ringbuffer.NewRingBuffer(5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}
	for i := 0; i < 2; i++ {
		rb.Read()
	}
	for i := 0; i < 2; i++ {
		rb.Write(i)
	}

	for i := 0; i < 5; i++ {
		val := rb.Read()
		v := val.(int)
		assert.Equal((i+2)%5, v)
	}
}

func TestWriteAt(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
	rb := ringbuffer.NewRingBuffer(5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}
	rb.WriteAt(3, 100)
	for i := 0; i < 5; i++ {
		val := rb.Read()
		v := val.(int)
		if i == 3 {
			assert.Equal(100, v)
		} else {
			assert.Equal(i, v)
		}
	}
}

func TestReadAt(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
	rb := ringbuffer.NewRingBuffer(5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}
	rb.WriteAt(3, 100)
	assert.Equal(100, rb.ReadAt(3))
}
