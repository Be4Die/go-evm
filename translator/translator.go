package translator

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/Be4Die/go-evm/vm"
)

// DirectiveType представляет тип директивы ассемблера
type DirectiveType int

const (
	DirectiveNone DirectiveType = iota
	DirectiveDS
	DirectiveORG
)

// Instruction представляет инструкцию с мнемоникой и операндами
type Instruction struct {
	LineNum    int
	Label      string
	Mnemonic   string
	Operands   []string
	Address    uint16
	Directive  DirectiveType
	Values     []uint32
	OrgAddress uint16
}

// SymbolTable представляет таблицу символов (меток)
type SymbolTable map[string]uint16

// Translator представляет транслятор ассемблера
type Translator struct {
	instructions []Instruction
	symbols      SymbolTable
	currentAddr  uint16
}

// NewTranslator создает новый транслятор
func NewTranslator() *Translator {
	return &Translator{
		symbols: make(SymbolTable),
	}
}

// Parse разбирает исходный код и строит список инструкций
func (t *Translator) Parse(source string) error {
	scanner := bufio.NewScanner(strings.NewReader(source))
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Пропускаем пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		
		// Удаляем комментарии в конце строки
		if idx := strings.Index(line, ";"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}
		
		// Парсим директивы и инструкции
		instr, err := t.parseLine(line, lineNum)
		if err != nil {
			return fmt.Errorf("line %d: %v", lineNum, err)
		}
		
		t.instructions = append(t.instructions, instr)
	}
	
	return scanner.Err()
}

// parseLine разбирает строку ассемблера
func (t *Translator) parseLine(line string, lineNum int) (Instruction, error) {
	instr := Instruction{LineNum: lineNum}
	
	// Проверяем наличие метки
	if idx := strings.Index(line, ":"); idx != -1 {
		instr.Label = strings.TrimSpace(line[:idx])
		line = strings.TrimSpace(line[idx+1:])
		
		if instr.Label == "" {
			return instr, errors.New("empty label")
		}
	}
	
	// Разбираем мнемонику и операнды
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return instr, errors.New("no instruction")
	}
	
	mnemonic := strings.ToUpper(fields[0])
	instr.Mnemonic = mnemonic
	
	// Обрабатываем директивы
	switch mnemonic {
	case "DS":
		instr.Directive = DirectiveDS
		if len(fields) < 2 {
			return instr, errors.New("DS directive requires values")
		}
		
		for _, valStr := range fields[1:] {
			val, err := parseValue(valStr)
			if err != nil {
				return instr, fmt.Errorf("invalid value '%s': %v", valStr, err)
			}
			instr.Values = append(instr.Values, val)
		}
		
	case "ORG":
		instr.Directive = DirectiveORG
		if len(fields) != 2 {
			return instr, errors.New("ORG directive requires address")
		}
		
		addr, err := parseAddress(fields[1])
		if err != nil {
			return instr, fmt.Errorf("invalid address '%s': %v", fields[1], err)
		}
		instr.OrgAddress = addr
		
	default:
		// Обрабатываем инструкции
		if len(fields) > 1 {
			operandsStr := strings.Join(fields[1:], " ")
			instr.Operands = t.parseOperands(operandsStr)
		}
	}
	
	return instr, nil
}

// parseOperands разбирает операнды инструкции
func (t *Translator) parseOperands(operandsStr string) []string {
	var operands []string
	current := strings.Builder{}
	inBrackets := false
	
	for _, r := range operandsStr {
		switch {
		case r == '[':
			inBrackets = true
			current.WriteRune(r)
		case r == ']':
			inBrackets = false
			current.WriteRune(r)
		case r == ',' && !inBrackets:
			operands = append(operands, strings.TrimSpace(current.String()))
			current.Reset()
		default:
			current.WriteRune(r)
		}
	}
	
	if current.Len() > 0 {
		operands = append(operands, strings.TrimSpace(current.String()))
	}
	
	return operands
}

// parseValue парсит числовое значение
func parseValue(s string) (uint32, error) {
	s = strings.TrimSpace(s)
	
	// Шестнадцатеричное число
	if strings.HasPrefix(s, "0x") {
		val, err := strconv.ParseUint(s[2:], 16, 32)
		return uint32(val), err
	}
	
	// Двоичное число
	if strings.HasPrefix(s, "0b") {
		val, err := strconv.ParseUint(s[2:], 2, 32)
		return uint32(val), err
	}
	
	// Число с плавающей точкой
	if strings.Contains(s, ".") {
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, err
		}
		return math.Float32bits(float32(val)), nil
	}
	
	// Десятичное число
	val, err := strconv.ParseUint(s, 10, 32)
	return uint32(val), err
}

// parseAddress парсит адрес
func parseAddress(s string) (uint16, error) {
	s = strings.TrimSpace(s)
	
	// Убираем скобки если есть
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = s[1 : len(s)-1]
	}
	
	// Шестнадцатеричное число
	if strings.HasPrefix(s, "0x") {
		val, err := strconv.ParseUint(s[2:], 16, 16)
		return uint16(val), err
	}
	
	// Десятичное число
	val, err := strconv.ParseUint(s, 10, 16)
	return uint16(val), err
}

// resolveOperand разрешает операнд в числовое значение
func (t *Translator) resolveOperand(operand string) (uint16, error) {
	// Если это числовое значение
	if isNumeric(operand) {
		return parseAddress(operand)
	}
	
	// Если это адрес в скобках
	if strings.HasPrefix(operand, "[") && strings.HasSuffix(operand, "]") {
		inner := operand[1 : len(operand)-1]
		if isNumeric(inner) {
			return parseAddress(inner)
		}
		// Ищем метку
		if addr, ok := t.symbols[inner]; ok {
			return addr, nil
		}
		return 0, fmt.Errorf("undefined label '%s'", inner)
	}
	
	// Ищем метку
	if addr, ok := t.symbols[operand]; ok {
		return addr, nil
	}
	
	return 0, fmt.Errorf("undefined label '%s'", operand)
}

// isNumeric проверяет, является ли строка числом
func isNumeric(s string) bool {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") {
		_, err := strconv.ParseUint(s[2:], 16, 64)
		return err == nil
	}
	if strings.HasPrefix(s, "0b") {
		_, err := strconv.ParseUint(s[2:], 2, 64)
		return err == nil
	}
	
	// Проверяем, что все символы - цифры или знак минус
	for i, r := range s {
		if !unicode.IsDigit(r) && !(i == 0 && r == '-') {
			return false
		}
	}
	
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// BuildSymbolTable строит таблицу символов
func (t *Translator) BuildSymbolTable() error {
	t.currentAddr = 0
	
	for i := range t.instructions {
		instr := &t.instructions[i]
		
		// Обрабатываем директиву ORG
		if instr.Directive == DirectiveORG {
			t.currentAddr = instr.OrgAddress
			continue
		}
		
		// Добавляем метку в таблицу символов
		if instr.Label != "" {
			if _, exists := t.symbols[instr.Label]; exists {
				return fmt.Errorf("line %d: duplicate label '%s'", instr.LineNum, instr.Label)
			}
			t.symbols[instr.Label] = t.currentAddr
		}
		
		// Вычисляем адрес для следующей инструкции
		if instr.Directive == DirectiveDS {
			instr.Address = t.currentAddr
			t.currentAddr += uint16(len(instr.Values) * 4)
		} else {
			instr.Address = t.currentAddr
			// Инструкция занимает 1 байт для опкода + 2 байта для каждого операнда
			t.currentAddr += 1 + uint16(len(instr.Operands))*2
		}
	}
	
	return nil
}

// TranslateFile транслирует файл с ассемблером в байт-код
func TranslateFile(inputFilename string) ([]byte, error) {
	data, err := os.ReadFile(inputFilename)
	if err != nil {
		return nil, err
	}
	
	return Translate(string(data))
}

// Translate транслирует инструкции в байт-код
func (t *Translator) Translate() ([]byte, error) {
	var code []byte
	
	for _, instr := range t.instructions {
		// Пропускаем директивы ORG
		if instr.Directive == DirectiveORG {
			continue
		}
		
		// Обрабатываем директиву DS
		if instr.Directive == DirectiveDS {
			for _, val := range instr.Values {
				// Записываем значение little-endian
				code = append(code, 
					byte(val),
					byte(val >> 8),
					byte(val >> 16),
					byte(val >> 24))
			}
			continue
		}
		
		// Обрабатываем инструкции
		opcode, err := t.getOpcode(instr.Mnemonic)
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", instr.LineNum, err)
		}
		
		code = append(code, opcode)
		
		// Обрабатываем операнды
		for _, operand := range instr.Operands {
			value, err := t.resolveOperand(operand)
			if err != nil {
				return nil, fmt.Errorf("line %d: %v", instr.LineNum, err)
			}
			
			// Записываем операнд little-endian
			code = append(code, byte(value), byte(value>>8))
		}
	}
	
	return code, nil
}

// getOpcode возвращает опкод для мнемоники
func (t *Translator) getOpcode(mnemonic string) (byte, error) {
	switch mnemonic {
	case "MOV":
		return vm.OP_MOV, nil
	case "ADD_I":
		return vm.OP_ADD_I, nil
	case "SUB_I":
		return vm.OP_SUB_I, nil
	case "MUL_I":
		return vm.OP_MUL_I, nil
	case "DIV_I":
		return vm.OP_DIV_I, nil
	case "ADD_F":
		return vm.OP_ADD_F, nil
	case "SUB_F":
		return vm.OP_SUB_F, nil
	case "MUL_F":
		return vm.OP_MUL_F, nil
	case "DIV_F":
		return vm.OP_DIV_F, nil
	case "CMP_I":
		return vm.OP_CMP_I, nil
	case "CMP_F":
		return vm.OP_CMP_F, nil
	case "JMP":
		return vm.OP_JMP, nil
	case "JZ":
		return vm.OP_JZ, nil
	case "JNZ":
		return vm.OP_JNZ, nil
	case "JC":
		return vm.OP_JC, nil
	case "JNC":
		return vm.OP_JNC, nil
	case "CALL":
		return vm.OP_CALL, nil
	case "RET":
		return vm.OP_RET, nil
	case "PUSH":
		return vm.OP_PUSH, nil
	case "POP":
		return vm.OP_POP, nil
	case "IN":
		return vm.OP_IN, nil
	case "OUT":
		return vm.OP_OUT, nil
	case "AND":
		return vm.OP_AND, nil
	case "OR":
		return vm.OP_OR, nil
	case "XOR":
		return vm.OP_XOR, nil
	case "NOT":
		return vm.OP_NOT, nil
	case "SHL":
		return vm.OP_SHL, nil
	case "SHR":
		return vm.OP_SHR, nil
	default:
		return 0, fmt.Errorf("unknown mnemonic '%s'", mnemonic)
	}
}

// Translate переводит исходный код в байт-код
func Translate(source string) ([]byte, error) {
	translator := NewTranslator()
	
	// Парсим исходный код
	if err := translator.Parse(source); err != nil {
		return nil, err
	}
	
	// Строим таблицу символов
	if err := translator.BuildSymbolTable(); err != nil {
		return nil, err
	}
	
	// Транслируем в байт-код
	return translator.Translate()
}