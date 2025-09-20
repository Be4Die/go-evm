package assembly

import (
	"bufio"
	"fmt"
	"os"
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
	dataAddress    uint32
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