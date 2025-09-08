package vm

import "math"

/* -- Арифметические операции для чисел с плавающей запятой -- */

// операция сложения для чисел с плавающей точкой
func floatAdd(a, b Word) (Word, bool, bool, bool) {
    fa := math.Float32frombits(uint32(a))
    fb := math.Float32frombits(uint32(b))
    result := math.Float32bits(fa + fb)
    return Word(result), false, false, false
}

// операция вычитания для чисел с плавающей точкой
func floatSub(a, b Word) (Word, bool, bool, bool) {
    fa := math.Float32frombits(uint32(a))
    fb := math.Float32frombits(uint32(b))
    result := math.Float32bits(fa - fb)
    return Word(result), false, false, false
}

// операция умножения для чисел с плавающей точкой
func floatMul(a, b Word) (Word, bool, bool, bool) {
    fa := math.Float32frombits(uint32(a))
    fb := math.Float32frombits(uint32(b))
    result := math.Float32bits(fa * fb)
    return Word(result), false, false, false
}

// операция деления для чисел с плавающей точкой
func floatDiv(a, b Word) (Word, bool, bool, bool) {
    fa := math.Float32frombits(uint32(a))
    fb := math.Float32frombits(uint32(b))
    result := math.Float32bits(fa / fb)
    return Word(result), false, false, false
}


/* -- Конструкторы арифметических команд для чисел с плавающей запятой -- */

// конструктор команды сложения чисел с плавающей запятой
func NewFloatAddCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: floatAdd,
    }
}

// конструктор команды вычитания чисел с плавающей запятой
func NewFloatSubCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: floatSub,
    }
}

// конструктор команды умножения чисел с плавающей запятой
func NewFloatMulCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: floatMul,
    }
}

// конструктор команды деления чисел с плавающей запятой
func NewFloatDivCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &binaryCommand{
        op1:    op1,
        op2:    op2,
        result: result,
        operation: floatDiv,
    }
}