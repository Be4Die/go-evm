package vm

import (
	"fmt"
)

// CPU интерфейс определяет основные операции центрального процессора
type CPU interface {
    // GetRegister возвращает текущее значение указанного регистра
    GetRegister(registerType RegisterType) (Word, error)

    // SetRegister устанавливает значение указанного регистра
    SetRegister(registerType RegisterType, value Word) error

    // Run начинает выполнение программного кода
    Run() error

    // Step выполняет одну следующую инструкцию
    Step() error

    // Reset сбрасывает состояние процессора в начальное положение
    Reset() error
}


type cpu struct {
    registers map[RegisterType]Word
    memory    Memory
    running   bool
}

// создает новый процессор с указанной памятью
func NewCPU(memory Memory) CPU {
    cpu := &cpu{
        registers: make(map[RegisterType]Word),
        memory:    memory,
    }
    cpu.Reset()
    return cpu
}

func (c *cpu) GetRegister(reg RegisterType) (Word, error) {
    if value, ok := c.registers[reg]; ok {
        return value, nil
    }
    return 0, fmt.Errorf("register %v not found", reg)
}

func (c *cpu) SetRegister(reg RegisterType, value Word) error {
    if _, ok := c.registers[reg]; !ok {
        return fmt.Errorf("register %v not found", reg)
    }
    c.registers[reg] = value
    return nil
}

func (c *cpu) Run() error {
    c.running = true
    for c.running {
        if err := c.Step(); err != nil {
            return err
        }
    }
    return nil
}

func (c *cpu) Step() error {
    ip, err := c.GetRegister(REG_IP)
    if err != nil {
        return err
    }

    // Чтение команды из памяти (базовая реализация)
    // В реальной реализации нужно декодировать команду
    cmdValue, err := c.memory.GetValue(ip)
    if err != nil {
        return err
    }

    // Увеличиваем указатель команд
    if err := c.SetRegister(REG_IP, ip+4); err != nil {
        return err
    }

    // Здесь должна быть логика декодирования и выполнения команды
    // Это placeholder для реальной реализации
    fmt.Printf("Executing command: %08x\n", cmdValue)

    return nil
}

func (c *cpu) Reset() error {
    // Инициализация регистров нулевыми значениями
    c.registers = map[RegisterType]Word{
        REG_IP:    0,
        REG_ACC:   0,
        REG_SP:    0xFFFF, // Начало стека
        REG_FLAGS: 0,
        REG_BP:    0,
        REG_R1:    0,
        REG_R2:    0,
        REG_R3:    0,
        REG_R4:    0,
        REG_R5:    0,
        REG_R6:    0,
        REG_R7:    0,
        REG_R8:    0,
    }
    c.running = false
    return nil
}