package vm

import (
	"errors"
	"fmt"
	"math"
)

// CommandFunc представляет тип функции для выполнения команд процессора.
// Каждая функция возвращает ошибку в случае неудачного выполнения.
type CommandFunc func() error

// push помещает значение на вершину стека.
// value: 32-битное значение для помещения в стек
// Возвращает ошибку в случае переполнения стека.
func (c *CPU) push(value uint32) error {
	if c.psw.sp == 0 {
		return errors.New("stack overflow")
	}
	c.psw.sp--
	stackAddr := c.getStackAddress()
	return c.memory.WriteWordAt(stackAddr, value)
}

// pop извлекает значение с вершины стека.
// Возвращает извлеченное значение или ошибку в случае проблем с памятью.
func (c *CPU) pop() (uint32, error) {
	stackAddr := c.getStackAddress()
	value, err := c.memory.ReadWordAt(stackAddr)
	if err != nil {
		return 0, err
	}
	c.psw.sp++
	if c.psw.sp > 31 {
		c.psw.sp = 31 // Сброс на вершину если превысили лимит
	}
	return value, nil
}

// getStackAddress вычисляет адрес в памяти для текущей позиции стека.
// Возвращает 16-битный адрес памяти.
func (c *CPU) getStackAddress() uint16 {
	return uint16(31-c.psw.sp) * 4 // 32 слова по 4 байта каждое
}

// readImmediateOffset читает 16-битное непосредственное значение из памяти.
// Используется для чтения аргументов команд.
// Возвращает прочитанное значение или ошибку.
func (c *CPU) readImmediateOffset() (uint16, error) {
    low, err := c.memory.ReadByteAt(c.psw.ip)
    if err != nil {
        return 0, err
    }
    high, err := c.memory.ReadByteAt(c.psw.ip + 1)
    if err != nil {
        return 0, err
    }
    c.psw.ip += 2
    return uint16(low) | (uint16(high) << 8), nil
}

// Реализации команд процессора

// movCommand реализует команду перемещения данных (OP_MOV).
// Перемещает значение с вершины стека в память по указанному смещению.
func (c *CPU) movCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	value, err := c.pop()
	if err != nil {
		return err
	}

	// Записать в память по смещению
	return c.memory.WriteWordAt(offset, value)
}

// addIntCommand реализует команду целочисленного сложения (OP_ADD_I).
// Складывает значение из памяти со значением на вершине стека.
// Устанавливает флаги ZERO, NEGATIVE, OVERFLOW и CARRY.
func (c *CPU) addIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить сложение
	result := stackValue + memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Проверить переполнение (если знаки одинаковые, но результат имеет другой знак)
	if ((stackValue^memValue)&0x80000000) == 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Проверить перенос
	carry := (stackValue > math.MaxUint32-memValue)
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Поместить результат обратно в стек
	return c.push(result)
}

// subIntCommand реализует команду целочисленного вычитания (OP_SUB_I).
// Вычитает значение из памяти из значения на вершине стека.
// Устанавливает флаги ZERO, NEGATIVE, OVERFLOW и CARRY.
func (c *CPU) subIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить вычитание
	result := stackValue - memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Проверить переполнение
	if ((stackValue^memValue)&0x80000000) != 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Проверить перенос/заём
	carry := stackValue >= memValue
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Поместить результат обратно в стек
	return c.push(result)
}

// mulIntCommand реализует команду целочисленного умножения (OP_MUL_I).
// Умножает значение из памяти на значение на вершине стека.
// Устанавливает флаги ZERO, NEGATIVE и OVERFLOW.
func (c *CPU) mulIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить умножение
	result := stackValue * memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Проверить переполнение (упрощённо)
	if (memValue != 0 && result/memValue != stackValue) {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Поместить результат обратно в стек
	return c.push(result)
}

// divIntCommand реализует команду целочисленного деления (OP_DIV_I).
// Делит значение на вершине стека на значение из памяти.
// Устанавливает флаги ZERO и NEGATIVE.
// Возвращает ошибку при делении на ноль.
func (c *CPU) divIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	if memValue == 0 {
		return errors.New("division by zero")
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить деление
	result := stackValue / memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// addFloatCommand реализует команду сложения чисел с плавающей точкой (OP_ADD_F).
// Складывает значение из памяти со значением на вершине стека.
// Устанавливает флаги FZERO и FNEGATIVE.
func (c *CPU) addFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Конвертировать в float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Выполнить сложение
	resultFloat := stackFloat + memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Установить флаги
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// subFloatCommand реализует команду вычитания чисел с плавающей точкой (OP_SUB_F).
// Вычитает значение из памяти из значения на вершине стека.
// Устанавливает флаги FZERO и FNEGATIVE.
func (c *CPU) subFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Конвертировать в float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Выполнить вычитание
	resultFloat := stackFloat - memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Установить флаги
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// mulFloatCommand реализует команду умножения чисел с плавающей точкой (OP_MUL_F).
// Умножает значение из памяти на значение на вершине стека.
// Устанавливает флаги FZERO и FNEGATIVE.
func (c *CPU) mulFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Конвертировать в float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Выполнить умножение
	resultFloat := stackFloat * memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Установить флаги
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// divFloatCommand реализует команду деления чисел с плавающей точкой (OP_DIV_F).
// Делит значение на вершине стека на значение из памяти.
// Устанавливает флаги FZERO и FNEGATIVE.
// Возвращает ошибку при делении на ноль.
func (c *CPU) divFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Конвертировать в float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	if memFloat == 0 {
		return errors.New("division by zero")
	}

	// Выполнить деление
	resultFloat := stackFloat / memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Установить флаги
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// cmpIntCommand реализует команду целочисленного сравнения (OP_CMP_I).
// Сравнивает значение на вершине стека со значением из памяти.
// Устанавливает флаги ZERO, NEGATIVE, OVERFLOW и CARRY.
// Не изменяет значения в стеке.
func (c *CPU) cmpIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Сравнить значения
	result := stackValue - memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Проверить переполнение
	if ((stackValue^memValue)&0x80000000) != 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Проверить перенос/заём
	carry := stackValue >= memValue
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Не возвращать результат в стек - сравнение только устанавливает флаги
	return nil
}

// cmpFloatCommand реализует команду сравнения чисел с плавающей точкой (OP_CMP_F).
// Сравнивает значение на вершине стека со значением из памяти.
// Устанавливает флаги FZERO и FNEGATIVE.
// Не изменяет значения в стеке.
func (c *CPU) cmpFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Конвертировать в float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Сравнить значения
	result := stackFloat - memFloat

	// Установить флаги
	c.psw.SetFlag(FLAG_FZERO, result == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, result < 0)

	// Не возвращать результат в стек - сравнение только устанавливает флаги
	return nil
}

// jmpCommand реализует команду безусловного перехода (OP_JMP).
// Устанавливает указатель инструкций на новое значение.
func (c *CPU) jmpCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Установить IP на новый адрес
	c.psw.SetIP(offset)
	return nil
}

// jzCommand реализует команду перехода если равно нулю (OP_JZ).
// Выполняет переход, если установлен флаг ZERO.
func (c *CPU) jzCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Переход если установлен флаг нуля
	if c.psw.GetFlag(FLAG_ZERO) {
		c.psw.SetIP(offset)
	}
	return nil
}

// jnzCommand реализует команду перехода если не равно нулю (OP_JNZ).
// Выполняет переход, если не установлен флаг ZERO.
func (c *CPU) jnzCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Переход если не установлен флаг нуля
	if !c.psw.GetFlag(FLAG_ZERO) {
		c.psw.SetIP(offset)
	}
	return nil
}

// jcCommand реализует команду перехода если перенос (OP_JC).
// Выполняет переход, если установлен флаг CARRY.
func (c *CPU) jcCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Переход если установлен флаг переноса
	if c.psw.GetFlag(FLAG_CARRY) {
		c.psw.SetIP(offset)
	}
	return nil
}

// jncCommand реализует команду перехода если нет переноса (OP_JNC).
// Выполняет переход, если не установлен флаг CARRY.
func (c *CPU) jncCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Переход если не установлен флаг переноса
	if !c.psw.GetFlag(FLAG_CARRY) {
		c.psw.SetIP(offset)
	}
	return nil
}

// callCommand реализует команду вызова подпрограммы (OP_CALL).
// Сохраняет адрес возврата в стеке и выполняет переход.
func (c *CPU) callCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Поместить адрес возврата (текущий IP) в стек
	returnAddress := c.psw.ip
	if err := c.push(uint32(returnAddress)); err != nil {
		return err
	}

	// Установить IP на новый адрес
	c.psw.SetIP(offset)
	return nil
}

// retCommand реализует команду возврата из подпрограммы (OP_RET).
// Восстанавливает адрес возврата из стека и устанавливает IP.
func (c *CPU) retCommand() error {
	// Прочитать и игнорировать непосредственное смещение (2 байта)
	_, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Извлечь адрес возврата из стека
	returnAddress, err := c.pop()
	if err != nil {
		return err
	}

	// Установить IP на адрес возврата
	c.psw.SetIP(uint16(returnAddress))
	return nil
}

// pushCommand реализует команду помещения в стек (OP_PUSH).
// Читает значение из памяти и помещает его на вершину стека.
func (c *CPU) pushCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	value, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Поместить значение в стек
	return c.push(value)
}

// popCommand реализует команду извлечения из стека (OP_POP).
// Извлекает значение с вершины стека и записывает его в память.
func (c *CPU) popCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	value, err := c.pop()
	if err != nil {
		return err
	}

	// Записать значение в память
	return c.memory.WriteWordAt(offset, value)
}

// inCommand реализует команду ввода данных (OP_IN).
// Читает целое число из стандартного ввода и записывает его в память.
func (c *CPU) inCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать ввод от пользователя
	var input uint32
	fmt.Print("Input (integer): ")
	_, err = fmt.Scan(&input)
	if err != nil {
		return err
	}

	// Записать ввод в память
	return c.memory.WriteWordAt(offset, input)
}

// outCommand реализует команду вывода данных (OP_OUT).
// Читает значение из памяти и выводит его на стандартный вывод.
func (c *CPU) outCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	value, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Вывести значение
	fmt.Printf("Output: %d\n", value)
	return nil
}

// andCommand реализует команду логического И (OP_AND).
// Выполняет побитовое И между значением из памяти и значением на вершине стека.
// Устанавливает флаги ZERO и NEGATIVE.
func (c *CPU) andCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить операцию AND
	result := stackValue & memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// orCommand реализует команду логического ИЛИ (OP_OR).
// Выполняет побитовое ИЛИ между значением из памяти и значением на вершине стека.
// Устанавливает флаги ZERO и NEGATIVE.
func (c *CPU) orCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить операцию OR
	result := stackValue | memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// xorCommand реализует команду исключающего ИЛИ (OP_XOR).
// Выполняет побитовое XOR между значением из памяти и значением на вершине стека.
// Устанавливает флаги ZERO и NEGATIVE.
func (c *CPU) xorCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать значение из памяти
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить операцию XOR
	result := stackValue ^ memValue

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// notCommand реализует команду логического НЕ (OP_NOT).
// Выполняет побитовую инверсию значения на вершине стека.
// Устанавливает флаги ZERO и NEGATIVE.
func (c *CPU) notCommand() error {
    // Извлечь значение из стека
    stackValue, err := c.pop()
    if err != nil {
        return err
    }

    // Выполнить операцию NOT
    result := ^stackValue

    // Установить флаги
    c.psw.SetFlag(FLAG_ZERO, result == 0)
    c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

    // Поместить результат обратно в стек
    return c.push(result)
}

// shlCommand реализует команду логического сдвига влево (OP_SHL).
// Сдвигает значение на вершине стека влево на количество бит из памяти.
// Устанавливает флаги ZERO, NEGATIVE и CARRY.
func (c *CPU) shlCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать количество сдвига из памяти
	shiftCount, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить сдвиг влево
	result := stackValue << (shiftCount & 0x1F) // Ограничить сдвиг 31 битом

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	c.psw.SetFlag(FLAG_CARRY, (stackValue>>(32-(shiftCount&0x1F))&1) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}

// shrCommand реализует команду логического сдвига вправо (OP_SHR).
// Сдвигает значение на вершине стека вправо на количество бит из памяти.
// Устанавливает флаги ZERO, NEGATIVE и CARRY.
func (c *CPU) shrCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Прочитать количество сдвига из памяти
	shiftCount, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Извлечь значение из стека
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Выполнить сдвиг вправо
	result := stackValue >> (shiftCount & 0x1F) // Ограничить сдвиг 31 битом

	// Установить флаги
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	c.psw.SetFlag(FLAG_CARRY, (stackValue>>((shiftCount&0x1F)-1)&1) != 0)

	// Поместить результат обратно в стек
	return c.push(result)
}