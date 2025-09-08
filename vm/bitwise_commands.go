package vm

/* -- Битовые операции над числами -- */

// Побитовое И
func bitwiseAND(a, b Word) (Word, bool, bool, bool) {
    return a & b, false, false, false
}

// Побитовое ИЛИ
func bitwiseOR(a, b Word) (Word, bool, bool, bool) {
    return a | b, false, false, false
}

// Побитовое XOR
func bitwiseXOR(a, b Word) (Word, bool, bool, bool) {
    return a ^ b, false, false, false
}

// Побитовое NOT
func bitwiseNOT(a Word) (Word, bool, bool, bool) {
    return ^a, false, false, false
}

// Логический сдвиг влево
func leftShift(a, b Word) (Word, bool, bool, bool) {
    shift := b % 32 // Ограничиваем сдвиг 32 битами
    result := a << shift
    carry := (a >> (32 - shift)) & 1 == 1
    return result, carry, false, false
}

// Логический сдвиг вправо
func rightShift(a, b Word) (Word, bool, bool, bool) {
    shift := b % 32 // Ограничиваем сдвиг 32 битами
    result := a >> shift
    carry := (a >> (shift - 1)) & 1 == 1
    return result, carry, false, false
}

/* -- Конструкторы битовых команд для чисел -- */

// конструктор команды логического И
func NewBitwiseAndCommand(op1 OperandReadWriter, op2 OperandReader) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: op1,
        operation: bitwiseAND,
    }
}

// конструктор команды логического ИЛИ
func NewBitwiseORCommand(op1 OperandReadWriter, op2 OperandReader) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: op1,
        operation: bitwiseOR,
    }
}

// конструктор команды логического исключающее ИЛИ
func NewBitwiseXORCommand(op1 OperandReadWriter, op2 OperandReader) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: op1,
        operation: bitwiseXOR,
    }
}

// конструктор команды логического инвертирования
func NewBitwiseNOTCommand(op OperandReadWriter) Command {
    return &unaryCommand{
        op:    op,
        operation: bitwiseNOT,
    }
}

// конструктор команды логического ИЛИ
func NewBitwiseLeftShiftCommand(op1 OperandReadWriter, op2 OperandReader) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: op1,
        operation: leftShift,
    }
}

// конструктор команды логического исключающее ИЛИ
func NewBitwiseRightShiftCommand(op1 OperandReadWriter, op2 OperandReader) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: op1,
        operation: rightShift,
    }
}
