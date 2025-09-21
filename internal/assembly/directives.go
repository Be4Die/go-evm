package assembly

import (
	"fmt"
	"strings"
)

// processDataLine обрабатывает строку в секции данных
// Выполняет разбор меток и директив определения данных
func (t *Translator) processDataLine(line string, lineNum int) error {
	// Обработка меток в данных
	parts := strings.SplitN(line, ":", 2)
	if len(parts) > 1 {
		label := strings.TrimSpace(parts[0])
		line = strings.TrimSpace(parts[1])
		
		if t.pass == 1 {
			t.symbolTable[label] = t.currentAddress
		}
	}

	// Обработка директив данных
	if strings.HasPrefix(strings.ToUpper(line), "DB ") {
		return t.processDB(line, lineNum)
	} else if strings.HasPrefix(strings.ToUpper(line), "DW ") {
		return t.processDW(line, lineNum)
	}

	return fmt.Errorf("invalid data directive at line %d", lineNum)
}

// processDB обрабатывает директиву DB (Define Byte)
// Инициализирует байтовые значения в памяти данных
func (t *Translator) processDB(line string, lineNum int) error {
	// Извлечение значений после DB
	valuesStr := strings.TrimSpace(line[3:])
	values := strings.Split(valuesStr, ",")
	
	for _, valStr := range values {
		valStr = strings.TrimSpace(valStr)
		value, err := t.parseConstant(valStr)
		if err != nil {
			return fmt.Errorf("invalid value '%s' at line %d: %v", valStr, lineNum, err)
		}
		
		if t.pass == 1 {
			t.data = append(t.data, dataItem{address: t.currentAddress, value: value, isByte: true})
			t.currentAddress++
		} else {
			t.currentAddress++
		}
	}
	
	return nil
}

// processDW обрабатывает директиву DW (Define Word)
// Инициализирует словные значения (4 байта) в памяти данных
func (t *Translator) processDW(line string, lineNum int) error {
	// Извлечение значений после DW
	valuesStr := strings.TrimSpace(line[3:])
	values := strings.Split(valuesStr, ",")
	
	for _, valStr := range values {
		valStr = strings.TrimSpace(valStr)
		value, err := t.parseConstant(valStr)
		if err != nil {
			return fmt.Errorf("invalid value '%s' at line %d: %v", valStr, lineNum, err)
		}
		
		if t.pass == 1 {
			t.data = append(t.data, dataItem{address: t.currentAddress, value: value, isByte: false})
			t.currentAddress += 4
		} else {
			t.currentAddress += 4
		}
	}
	
	return nil
}

// processEQU обрабатывает директиву EQU (Equate)
// Создает символические константы для использования в коде
func (t *Translator) processEQU(line string, lineNum int) error {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return fmt.Errorf("invalid EQU directive at line %d", lineNum)
	}
	
	name := parts[0]
	valueStr := parts[2]
	
	value, err := t.parseConstant(valueStr)
	if err != nil {
		return fmt.Errorf("invalid constant value '%s' at line %d: %v", valueStr, lineNum, err)
	}
	
	if t.pass == 1 {
		t.constants[name] = value
	}
	
	return nil
}