package assembly

import (
	"fmt"
	"strings"
)

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

func (t *Translator) processORG(line string, lineNum int) error {
	addrStr := strings.TrimSpace(line[4:])
	addr, err := t.parseConstant(addrStr)
	if err != nil {
		return fmt.Errorf("invalid address '%s' at line %d: %v", addrStr, lineNum, err)
	}
	
	if t.pass == 1 {
		t.currentAddress = addr
		if t.startAddress == 0 {
			t.startAddress = addr
		}
	}
	
	return nil
}

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