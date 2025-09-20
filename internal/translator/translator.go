// translator.go
package translator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Be4Die/go-evm/internal/vm"
)

// Translator представляет ассемблер для виртуальной машины
type Translator struct {
	symbolTable    map[string]uint32
	constants      map[string]uint32
	instructions   []instruction
	data           []dataItem
	currentAddress uint32
	startAddress   uint32
	pass           int
	dataAddress    uint32 // Добавлено для отслеживания адреса данных
}

type instruction struct {
	lineNum    int
	address    uint32
	opcode     int
	operand    uint32
	label      string
	mnemonic   string
	operandStr string
}

type dataItem struct {
	address uint32
	value   uint32
	isByte  bool
}

// NewTranslator создает новый экземпляр транслятора
func NewTranslator() *Translator {
	return &Translator{
		symbolTable:  make(map[string]uint32),
		constants:    make(map[string]uint32),
		instructions: make([]instruction, 0),
		data:         make([]dataItem, 0),
		dataAddress:  0x0100, // Данные начинаются с 0x0100
	}
}

// Assemble трансляция ассемблерного кода в бинарный формат
func (t *Translator) Assemble(inputFile, outputFile string) error {
	// Первый проход - построение таблицы символов
	t.pass = 1
	t.currentAddress = 0x0200 // Код начинается с 0x0200
	t.startAddress = 0x0200
	if err := t.processFile(inputFile); err != nil {
		return err
	}

	// Второй проход - генерация кода
	t.pass = 2
	t.currentAddress = 0x0200
	t.instructions = make([]instruction, 0)
	if err := t.processFile(inputFile); err != nil {
		return err
	}

	// Запись выходного файла
	return t.writeOutput(outputFile)
}

func (t *Translator) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	inDataSection := false

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Пропуск пустых строк и комментариев
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "//") {
			continue
		}

		// Удаление inline-комментариев
		if commentIndex := strings.Index(line, ";"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}
		if commentIndex := strings.Index(line, "//"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}

		// Обработка директив секций
		if strings.ToUpper(line) == "DS" {
			inDataSection = true
			t.currentAddress = t.dataAddress // Переключаемся на адрес данных
			continue
		}

		if strings.ToUpper(line) == "DE" {
			inDataSection = false
			t.currentAddress = 0x0200 // Возвращаемся к адресу кода
			continue
		}

		// Обработка секции данных
		if inDataSection {
			if err := t.processDataLine(line, lineNum); err != nil {
				return err
			}
			continue
		}

		// Обработка секции кода
		if err := t.processCodeLine(line, lineNum); err != nil {
			return err
		}
	}

	return scanner.Err()
}

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

func (t *Translator) processCodeLine(line string, lineNum int) error {
	// Обработка меток
	parts := strings.SplitN(line, ":", 2)
	if len(parts) > 1 {
		label := strings.TrimSpace(parts[0])
		line = strings.TrimSpace(parts[1])
		
		if t.pass == 1 {
			t.symbolTable[label] = t.currentAddress
		}
		
		if line == "" {
			return nil // Только метка без инструкции
		}
	}

	// Обработка директивы ORG
	if strings.HasPrefix(strings.ToUpper(line), "ORG ") {
		return t.processORG(line, lineNum)
	}

	// Обработка директивы EQU
	if strings.HasPrefix(strings.ToUpper(line), "EQU ") {
		return t.processEQU(line, lineNum)
	}

	// Обработка инструкций
	return t.processInstruction(line, lineNum)
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

func (t *Translator) processInstruction(line string, lineNum int) error {
	// Разбор инструкции и операндов
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}
	
	mnemonic := strings.ToUpper(parts[0])
	operandStr := ""
	if len(parts) > 1 {
		operandStr = strings.Join(parts[1:], " ")
	}
	
	// Обработка псевдоинструкций
	if mnemonic == "HALT" {
		mnemonic = "JMP"
		operandStr = "0x0000"
	}
	
	// Получение опкода
	opcode, err := t.getOpcode(mnemonic)
	if err != nil {
		return fmt.Errorf("unknown instruction '%s' at line %d", mnemonic, lineNum)
	}
	
	// Разбор операнда (только во втором проходе)
	var operand uint32 = 0
	if t.pass == 2 {
		operand, err = t.parseOperand(operandStr)
		if err != nil {
			return fmt.Errorf("invalid operand '%s' at line %d: %v", operandStr, lineNum, err)
		}
	}
	
	if t.pass == 1 {
		t.instructions = append(t.instructions, instruction{
			lineNum:    lineNum,
			address:    t.currentAddress,
			opcode:     opcode,
			operand:    0, // В первом проходе операнд не вычисляем
			mnemonic:   mnemonic,
			operandStr: operandStr,
		})
		t.currentAddress += 3
	} else {
		t.instructions = append(t.instructions, instruction{
			lineNum:    lineNum,
			address:    t.currentAddress,
			opcode:     opcode,
			operand:    operand,
			mnemonic:   mnemonic,
			operandStr: operandStr,
		})
		t.currentAddress += 3
	}
	
	return nil
}

func (t *Translator) parseOperand(operandStr string) (uint32, error) {
	if operandStr == "" {
		return 0, nil
	}
	
	// Косвенная адресация [addr]
	if strings.HasPrefix(operandStr, "[") && strings.HasSuffix(operandStr, "]") {
		addrStr := operandStr[1 : len(operandStr)-1]
		return t.parseConstant(addrStr)
	}
	
	// Прямая адресация (для переходов)
	return t.parseConstant(operandStr)
}

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

func (t *Translator) getOpcode(mnemonic string) (int, error) {
	opcodes := map[string]int{
		"MOV":   vm.OP_MOV,
		"ADD_I": vm.OP_ADD_I,
		"SUB_I": vm.OP_SUB_I,
		"MUL_I": vm.OP_MUL_I,
		"DIV_I": vm.OP_DIV_I,
		"ADD_F": vm.OP_ADD_F,
		"SUB_F": vm.OP_SUB_F,
		"MUL_F": vm.OP_MUL_F,
		"DIV_F": vm.OP_DIV_F,
		"CMP_I": vm.OP_CMP_I,
		"CMP_F": vm.OP_CMP_F,
		"JMP":   vm.OP_JMP,
		"JZ":    vm.OP_JZ,
		"JNZ":   vm.OP_JNZ,
		"JC":    vm.OP_JC,
		"JNC":   vm.OP_JNC,
		"CALL":  vm.OP_CALL,
		"RET":   vm.OP_RET,
		"PUSH":  vm.OP_PUSH,
		"POP":   vm.OP_POP,
		"IN":    vm.OP_IN,
		"OUT":   vm.OP_OUT,
		"AND":   vm.OP_AND,
		"OR":    vm.OP_OR,
		"XOR":   vm.OP_XOR,
		"NOT":   vm.OP_NOT,
		"SHL":   vm.OP_SHL,
		"SHR":   vm.OP_SHR,
	}
	
	if opcode, exists := opcodes[mnemonic]; exists {
		return opcode, nil
	}
	
	return 0, fmt.Errorf("unknown mnemonic")
}

func (t *Translator) writeOutput(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	
	// Запись стартового адреса
	fmt.Fprintf(writer, "0x%04X\n", t.startAddress)
	
	// Запись секции данных
	fmt.Fprintln(writer, "DS")
	for _, item := range t.data {
		if item.isByte {
		    fmt.Fprintf(writer, "0x%04X %d\n", item.address, item.value)
		} else {
		    fmt.Fprintf(writer, "0x%04X %d\n", item.address, item.value)
		}
	}
	fmt.Fprintln(writer, "DE")
	
	// Запись секции кода
	for _, instr := range t.instructions {
		// Преобразование операнда в 2 байта в little-endian порядке
		operandLow := byte(instr.operand & 0xFF)
		operandHigh := byte(instr.operand >> 8)
		
		// Формирование строки из 6 шестнадцатеричных символов
		line := fmt.Sprintf("%02X%02X%02X\n", instr.opcode, operandLow, operandHigh)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	
	return writer.Flush()
}