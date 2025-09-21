// translator.go
package assembly

import (
	"bufio"
	"fmt"
	"os"
)

// Translator представляет основной транслятор ассемблерного кода в машинный
// Выполняет двухпроходную компиляцию с построением таблиц символов и констант
type Translator struct {
	symbolTable    map[string]uint32  // Таблица символов (меток)
	constants      map[string]uint32  // Таблица констант
	instructions   []instruction      // Список инструкций
	data           []dataItem         // Секция данных
	currentAddress uint32             // Текущий адрес компиляции
	startAddress   uint32             // Стартовый адрес программы
	pass           int                // Номер текущего прохода (1 или 2)
	dataAddress    uint32             // Базовый адрес секции данных
	entryLabel     string             // Метка точки входа
}

// instruction представляет машинную инструкцию с опкодом и операндом
type instruction struct {
	lineNum    int    // Номер строки в исходном коде
	address    uint32 // Адрес инструкции в памяти
	opcode     int    // Код операции
	operand    uint32 // Значение операнда
	label      string // Метка инструкции
	mnemonic   string // Мнемоника инструкции
	operandStr string // Строковое представление операнда
}

// dataItem представляет элемент данных в памяти
type dataItem struct {
	address uint32 // Адрес в памяти
	value   uint32 // Значение данных
	isByte  bool   // Флаг типа данных (байт/слово)
}

// NewTranslator создает новый экземпляр транслятора с инициализированными таблицами
func NewTranslator() *Translator {
	return &Translator{
		symbolTable:  make(map[string]uint32),
		constants:    make(map[string]uint32),
		instructions: make([]instruction, 0),
		data:         make([]dataItem, 0),
		dataAddress:  0x0100, // Данные начинаются с 0x0100
		entryLabel:   "",
	}
}

// Assemble выполняет двухпроходную трансляцию ассемблерного кода в бинарный формат
// inputFile - путь к исходному файлу, outputFile - путь для выходного файла
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

// writeOutput записывает результат трансляции в выходной файл
// Формат файла включает стартовый адрес, секцию данных и машинный код
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