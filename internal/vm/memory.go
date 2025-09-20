package vm

import (
	"errors"
)

// Memory представляет виртуальную память машины.
// Содержит байтовый массив для хранения данных и размер памяти.
type Memory struct {
	data []byte  // байтовый массив для хранения данных памяти
	size uint32  // общий размер памяти в байтах
}

// NewMemory создает и возвращает новый экземпляр Memory заданного размера.
// size: размер памяти в байтах
func NewMemory(size uint32) *Memory {
	return &Memory{
		data: make([]byte, size),
		size: size,
	}
}

// ReadByteAt читает байт из памяти по указанному адресу.
// addr: 16-битный адрес для чтения
// Возвращает прочитанный байт или ошибку, если адрес вне диапазона.
func (m *Memory) ReadByteAt(addr uint16) (byte, error) {
	address := uint32(addr)
	if address >= m.size {
		return 0, errors.New("memory read error: address out of bounds")
	}
	return m.data[address], nil
}

// WriteByteAt записывает байт в память по указанному адресу.
// addr: 16-битный адрес для записи
// value: байт для записи
// Возвращает ошибку, если адрес вне диапазона.
func (m *Memory) WriteByteAt(addr uint16, value byte) error {
	address := uint32(addr)
	if address >= m.size {
		return errors.New("memory write error: address out of bounds")
	}
	m.data[address] = value
	return nil
}

// ReadWordAt читает 32-битное слово из памяти по указанному адресу.
// Собирает слово из 4 последовательных байтов (little-endian).
// addr: 16-битный адрес для чтения (указывает на первый байт слова)
// Возвращает прочитанное слово или ошибку, если адрес вне диапазона.
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

// WriteWordAt записывает 32-битное слово в память по указанному адресу.
// Разбивает слово на 4 последовательных байта (little-endian).
// addr: 16-битный адрес для записи (указывает на первый байт слова)
// value: 32-битное слово для записи
// Возвращает ошибку, если адрес вне диапазона.
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