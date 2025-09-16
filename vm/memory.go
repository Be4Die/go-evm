// vm/memory.go
package vm

import (
	"errors"
)

type Memory struct {
	data []byte
	size uint32
}

func NewMemory(size uint32) *Memory {
	return &Memory{
		data: make([]byte, size),
		size: size,
	}
}

func (m *Memory) ReadByteAt(addr uint16) (byte, error) {
	address := uint32(addr)
	if address >= m.size {
		return 0, errors.New("memory read error: address out of bounds")
	}
	return m.data[address], nil
}

func (m *Memory) WriteByteAt(addr uint16, value byte) error {
	address := uint32(addr)
	if address >= m.size {
		return errors.New("memory write error: address out of bounds")
	}
	m.data[address] = value
	return nil
}

func (m *Memory) ReadWordAt(addr uint16) (uint32, error) {
	address := uint32(addr)
	if address+3 >= m.size {
		return 0, errors.New("memory read error: address out of bounds")
	}
	return uint32(m.data[address]) |
		uint32(m.data[address+1])<<8 |
		uint32(m.data[address+2])<<16 |
		uint32(m.data[address+3])<<24, nil
}

func (m *Memory) WriteWordAt(addr uint16, value uint32) error {
	address := uint32(addr)
	if address+3 >= m.size {
		return errors.New("memory write error: address out of bounds")
	}
	m.data[address] = byte(value)
	m.data[address+1] = byte(value >> 8)
	m.data[address+2] = byte(value >> 16)
	m.data[address+3] = byte(value >> 24)
	return nil
}