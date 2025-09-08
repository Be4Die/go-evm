package vm

import "math"

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