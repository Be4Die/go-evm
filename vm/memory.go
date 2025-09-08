package vm

import "errors"

// Memory интерфейс определяет операции работы с памятью виртуальной машины
type Memory interface {
    // SetValue записывает значение по указанному адресу
    // Возвращает ошибку если адрес недопустим
    SetValue(addr Word, value Word) error
    
    // GetValue читает значение по указанному адресу
    // Возвращает ошибку если адрес недопустим
    GetValue(addr Word) (Word, error)
}

type memory struct {
    data []byte
    size uint32
}

// NewMemory создает новую память указанного размера
func NewMemory(size uint32) Memory {
    return &memory{
        data: make([]byte, size),
        size: size,
    }
}

func (m *memory) SetValue(addr Word, value Word) error {
    if uint32(addr)+4 > m.size {
        return errors.New("address out of memory bounds")
    }
    
    // Младший байт first (little-endian)
    m.data[addr] = byte(value)
    m.data[addr+1] = byte(value >> 8)
    m.data[addr+2] = byte(value >> 16)
    m.data[addr+3] = byte(value >> 24)
    return nil
}

func (m *memory) GetValue(addr Word) (Word, error) {
    if uint32(addr)+4 > m.size {
        return 0, errors.New("address out of memory bounds")
    }
    
    value := Word(m.data[addr]) |
        Word(m.data[addr+1])<<8 |
        Word(m.data[addr+2])<<16 |
        Word(m.data[addr+3])<<24
    return value, nil
}