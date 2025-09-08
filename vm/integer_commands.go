package vm

/* -- Арифметические операции для целочисленных -- */

// операция сложения для целых чисел
func integerAdd(a, b Word) (Word, bool, bool, bool) {
    result := a + b
    carry := result < a
    overflow := (a^b)&(a^result)>>31 != 0
    // Вспомогательный перенос происходит при переносе из 3-го бита в 4-й
    auxiliaryCarry := (a&0xF + b&0xF) > 0xF
    return result, carry, overflow, auxiliaryCarry
}

// операция вычитания для целых чисел
func integerSub(a, b Word) (Word, bool, bool, bool) {
    result := a - b
    carry := a < b
    overflow := (a^b)&(a^result)>>31 != 0
    // Вспомогательный перенос происходит при заёме из 4-го бита
    auxiliaryCarry := (a&0xF) < (b&0xF)
    return result, carry, overflow, auxiliaryCarry
}

// операция умножения для целых чисел
func integerMul(a, b Word) (Word, bool, bool, bool) {
    result := a * b
    // Проверка переполнения: если a != 0 и result / a != b, то произошло переполнение
    overflow := a != 0 && result/a != b
    carry := overflow
    // Для умножения AF обычно не определяется, возвращаем false
    return result, carry, overflow, false
}

// операция деления для целых чисел
func integerDiv(a, b Word) (Word, bool, bool, bool) {
    if b == 0 {
        // Деление на ноль: устанавливаем флаги и возвращаем 0
        return 0, true, true, false
    }
    result := a / b
    // Для деления AF обычно не определяется, возвращаем false
    return result, false, false, false
}


/* -- Конструкторы арифметических команд для целочисленных -- */

// конструктор команды целочисленного сложения
func NewIntAddCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: integerAdd,
    }
}

// конструктор команды целочисленного вычитания
func NewIntSubCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: integerSub,
    }
}

// конструктор команды целочисленного умножения
func NewIntMulCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: integerMul,
    }
}

// конструктор команды целочисленного деления
func NewIntDivCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: integerDiv,
    }
}