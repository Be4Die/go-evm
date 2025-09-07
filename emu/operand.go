package emu

// OperandType определяет тип операнда
type OperandType uint8

const (
    OPERAND_REGISTER OperandType = iota
    OPERAND_CONSTANT
    OPERAND_MEMORY
)

// Operand интерфейс представляет операнд команды
type Operand interface {
    Type() OperandType
    Value(cpu CPU, memory Memory) (Word, error)
}

// RegisterOperand представляет регистровый операнд
type RegisterOperand struct {
    Register RegisterType
}

func (r RegisterOperand) Type() OperandType {
    return OPERAND_REGISTER
}

func (r RegisterOperand) Value(cpu CPU, memory Memory) (Word, error) {
    return cpu.GetRegister(r.Register)
}

// ConstantOperand представляет непосредственное значение
type ConstantOperand struct {
    Value Word
}

func (c ConstantOperand) Type() OperandType {
    return OPERAND_CONSTANT
}

func (c ConstantOperand) Value(cpu CPU, memory Memory) (Word, error) {
    return c.Value, nil
}

// MemoryOperand представляет операнд в памяти
type MemoryOperand struct {
    Address Word
}

func (m MemoryOperand) Type() OperandType {
    return OPERAND_MEMORY
}

func (m MemoryOperand) Value(cpu CPU, memory Memory) (Word, error) {
    return memory.GetValue(m.Address)
}