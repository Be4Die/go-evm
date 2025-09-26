package assembly

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Be4Die/go-evm/internal/assembly/grammar"
	"github.com/antlr4-go/antlr/v4"
)

// Translator представляет основной транслятор ассемблерного кода в машинный
type Translator struct {
	symbolTable    map[string]uint32  // Таблица символов (меток)
	constants      map[string]uint32  // Таблица констант
	instructions   []Instruction      // Список инструкций
	data           []DataItem         // Секция данных
	currentAddress uint32             // Текущий адрес компиляции
	startAddress   uint32             // Стартовый адрес программы
	pass           int                // Номер текущего прохода (1 или 2)
	dataAddress    uint32             // Базовый адрес секции данных
	entryLabel     string             // Метка точки входа
	currentSection string             // Текущая секция (.data или .code)
}

// Instruction представляет машинную инструкцию
type Instruction struct {
	Address    uint32 // Адрес инструкции
	Opcode     int    // Код операции
	Operand    uint32 // Операнд
	Label      string // Метка инструкции
	Mnemonic   string // Мнемоника
	OperandStr string // Строковое представление операнда
}

// DataItem представляет элемент данных
type DataItem struct {
	Address uint32 // Адрес в памяти
	Value   uint32 // Значение
	IsByte  bool   // Флаг типа данных (байт/слово)
}

// NewTranslator создает новый экземпляр транслятора
func NewTranslator() *Translator {
	return &Translator{
		symbolTable:  make(map[string]uint32),
		constants:    make(map[string]uint32),
		instructions: make([]Instruction, 0),
		data:         make([]DataItem, 0),
		dataAddress:  0x0100,
		startAddress: 0x0200,
		currentSection: ".code",
	}
}

// Assemble выполняет двухпроходную трансляцию
func (t *Translator) Assemble(inputFile, outputFile string, debug bool) error {
	// Первый проход - построение таблицы символов
	t.pass = 1
	t.currentAddress = t.startAddress
	if err := t.processFile(inputFile); err != nil {
		return err
	}

	// Второй проход - генерация кода
	t.pass = 2
	t.currentAddress = t.startAddress
	t.instructions = make([]Instruction, 0)
	if err := t.processFile(inputFile); err != nil {
		return err
	}

	// Запись выходного файла
	if err := t.writeOutput(outputFile); err != nil {
		return err
	}

	// Запись файла символов при включенном debug-режиме
	if debug {
		symFile := strings.TrimSuffix(outputFile, filepath.Ext(outputFile)) + ".sym"
		if err := t.writeSymbols(symFile); err != nil {
			return err
		}
	}

	return nil
}

// processFile обрабатывает файл ассемблера
func (t *Translator) processFile(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	lexer := grammar.NewAssemblerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewAssemblerParser(stream)

	parser.RemoveErrorListeners()
	errorListener := NewErrorListener()
	parser.AddErrorListener(errorListener)

	tree := parser.Program()
	listener := NewAssemblerListener(t)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if errorListener.HasErrors() {
		return fmt.Errorf("parsing errors: %v", errorListener.ErrorMessages())
	}

	return nil
}

// writeOutput записывает результат трансляции в файл
func (t *Translator) writeOutput(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Запись стартового адреса
	if t.entryLabel != "" {
		if addr, exists := t.symbolTable[t.entryLabel]; exists {
			t.startAddress = addr
		} else {
			return fmt.Errorf("entry label %s not found", t.entryLabel)
		}
	}
	fmt.Fprintf(writer, "0x%04X\n", t.startAddress)

	// Запись секции данных
	fmt.Fprintln(writer, "DS")
	for _, item := range t.data {
		fmt.Fprintf(writer, "0x%04X %d\n", item.Address, item.Value)
	}
	fmt.Fprintln(writer, "DE")

	// Запись секции кода
	for _, instr := range t.instructions {
		operandLow := byte(instr.Operand & 0xFF)
		operandHigh := byte(instr.Operand >> 8)
		line := fmt.Sprintf("%02X%02X%02X\n", instr.Opcode, operandLow, operandHigh)
		writer.WriteString(line)
	}

	return writer.Flush()
}

// writeSymbols записывает файл символов
func (t *Translator) writeSymbols(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create symbol file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for label, addr := range t.symbolTable {
		fmt.Fprintf(writer, "0x%04X %s\n", addr, label)
	}

	return writer.Flush()
}

// parseConstant преобразует строку в числовое значение
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

// getOpcode возвращает числовой код операции для мнемоники
func (t *Translator) getOpcode(mnemonic string) (int, error) {
	opcodes := map[string]int{
		"MOV":   0x01, "ADD_I": 0x02, "SUB_I": 0x03, "MUL_I": 0x04,
		"DIV_I": 0x05, "ADD_F": 0x06, "SUB_F": 0x07, "MUL_F": 0x08,
		"DIV_F": 0x09, "CMP_I": 0x0A, "CMP_F": 0x0B, "JMP": 0x0C,
		"JZ":    0x0D, "JNZ": 0x0E, "JC": 0x0F, "JNC": 0x10,
		"CALL":  0x11, "RET": 0x12, "PUSH": 0x13, "POP": 0x14,
		"IN":    0x15, "OUT": 0x16, "AND": 0x17, "OR": 0x18,
		"XOR":   0x19, "NOT": 0x1A, "SHL": 0x1B, "SHR": 0x1C,
		"HALT":  0x1D,
	}

	if opcode, exists := opcodes[mnemonic]; exists {
		return opcode, nil
	}

	return 0, fmt.Errorf("unknown mnemonic: %s", mnemonic)
}