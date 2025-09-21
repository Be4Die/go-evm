// symbols.go
package assembly

import (
	"fmt"
	"strconv"
	"strings"
)

// parseConstant преобразует строковое представление константы в числовое значение
// Поддерживает шестнадцатеричные (0x), двоичные (0b) и десятичные форматы
// Выполняет поиск в таблицах символов и констант
func (t *Translator) parseConstant(constStr string) (uint32, error) {
	constStr = strings.TrimSpace(constStr)
	
	// Проверка на существующую константу
	if value, exists := t.constants[constStr]; exists {
		return value, nil
	}
	
	// Проверка на существующий символ
	if value, exists := t.symbolTable[constStr]; exists {
		return value, nil
	}
	
	// Шестнадцатеричное число
	if strings.HasPrefix(constStr, "0x") {
		val, err := strconv.ParseUint(constStr[2:], 16, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	}
	
	// Двоичное число
	if strings.HasPrefix(constStr, "0b") {
		val, err := strconv.ParseUint(constStr[2:], 2, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	}
	
	// Десятичное число
	if val, err := strconv.ParseUint(constStr, 10, 32); err == nil {
		return uint32(val), nil
	}
	
	return 0, fmt.Errorf("unknown constant or symbol: %s", constStr)
}