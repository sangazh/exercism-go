package circular

import (
	"errors"
)

type Buffer struct {
	buffer []byte
	size   int
	cursor int
}

func (b *Buffer) len() int {
	return len(b.buffer)
}

func (b *Buffer) isFull() bool {
	return b.len() == b.size
}

func (b *Buffer) next() {
	b.cursor = (b.cursor + 1) % b.size
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, size),
		size:   size,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.len() == 0 {
		return 0, errors.New("empty")
	}
	oldest := b.buffer[0]
	b.buffer = b.buffer[1:]
	return oldest, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isFull() {
		return errors.New("buffer is full")
	}
	b.buffer = append(b.buffer, c)
	b.next()
	return nil
}
func (b *Buffer) Overwrite(c byte) {
	if b.isFull() {
		b.buffer = append(b.buffer[1:], c)
	} else {
		b.buffer = append(b.buffer, c)
	}
	b.next()
}
func (b *Buffer) Reset() {
	b.buffer = make([]byte, 0, b.size)
	b.cursor = 0
}
