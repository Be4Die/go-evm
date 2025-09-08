package vm

// Флаги состояния процессора (битовые маски)
// Зарезервированы биты для возможного расширения (1, 3, 5, 12-15)
const (
	CF = 1 << 0  // Carry Flag - флаг переноса
	PF = 1 << 2  // Parity Flag - флаг четности
	AF = 1 << 4  // Auxiliary Flag - вспомогательный флаг переноса
	ZF = 1 << 6  // Zero Flag - флаг нуля
	SF = 1 << 7  // Sign Flag - флаг знака
	OF = 1 << 11 // Overflow Flag - флаг переполнения
	TF = 1 << 8  // Trap Flag - флаг трассировки
	IF = 1 << 9  // Interrupt Flag - флаг прерываний
	DF = 1 << 10 // Direction Flag - флаг направления
)

// updateFlags обновляет флаги на основе результата операции
func updateFlags(flags Word, result Word, carry bool, overflow bool, auxiliaryCarry bool) Word {
    // Сброс арифметических флагов
    flags &^= (CF | PF | AF | ZF | SF | OF)

    // Установка флага переноса
    if carry {
        flags |= CF
    }

    // Установка вспомогательного флага переноса
    if auxiliaryCarry {
        flags |= AF
    }

    // Установка флага переполнения
    if overflow {
        flags |= OF
    }

    // Установка флага нуля
    if result == 0 {
        flags |= ZF
    }

    // Установка флага знака (старший бит)
    if result&(1<<31) != 0 {
        flags |= SF
    }

    // Установка флага четности
    if parityEven(result) {
        flags |= PF
    }

    return flags
}

// parityEven проверяет четность количества установленных битов
func parityEven(value Word) bool {
    count := 0
    for value != 0 {
        count++
        value &= value - 1
    }
    return count%2 == 0
}