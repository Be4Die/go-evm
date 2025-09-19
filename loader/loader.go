package loader

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Be4Die/go-evm/vm"
)

// Loader представляет загрузчик программ для виртуальной машины.
// Отвечает за чтение и парсинг файлов программ, а также загрузку их в память.
type Loader struct{}

// NewLoader создает и возвращает новый экземпляр Loader.
func NewLoader() *Loader {
	return &Loader{}
}

// LoadProgram загружает программу из файла в память виртуальной машины.
// filename: путь к файлу программы
// memory: экземпляр памяти для загрузки программы
// Возвращает стартовый адрес программы и ошибку в случае неудачи.
func (l *Loader) LoadProgram(filename string, memory *vm.Memory) (uint16, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentAddr := uint16(0)
	lineNum := 0
	startAddr := uint16(0)
	dataSection := false
	codeSection := false

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Пропустить пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		// Удалить inline-комментарии
		if commentIndex := strings.Index(line, "//"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}
		if commentIndex := strings.Index(line, "#"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}

		// Парсить стартовый адрес из первой строки
		if lineNum == 1 {
			addr, err := strconv.ParseUint(line, 0, 16)
			if err != nil {
				return 0, fmt.Errorf("invalid start address: %v", err)
			}
			startAddr = uint16(addr)
			currentAddr = startAddr
			continue
		}

		// Проверить директиву DS (начало секции данных)
		if line == "DS" {
			dataSection = true
			continue
		}

		// Проверить директиву DE (конец секции данных)
		if line == "DE" {
			dataSection = false
			codeSection = true
			continue
		}

		// Обработать секцию данных
		if dataSection {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				return 0, fmt.Errorf("invalid data at line %d", lineNum)
			}
			addrStr := parts[0]
			valueStr := parts[1]

			addr, err := strconv.ParseUint(addrStr, 0, 16)
			if err != nil {
				return 0, fmt.Errorf("invalid address at line %d: %v", lineNum, err)
			}

			var value uint32
			// Попытаться парсить как float сначала
			if strings.Contains(valueStr, ".") {
				floatVal, err := strconv.ParseFloat(valueStr, 32)
				if err != nil {
					return 0, fmt.Errorf("invalid float value at line %d: %v", lineNum, err)
				}
				value = math.Float32bits(float32(floatVal))
			} else {
				// Парсить как integer
				intVal, err := strconv.ParseUint(valueStr, 0, 32)
				if err != nil {
					return 0, fmt.Errorf("invalid integer value at line %d: %v", lineNum, err)
				}
				value = uint32(intVal)
			}

			if err := memory.WriteWordAt(uint16(addr), value); err != nil {
				return 0, fmt.Errorf("memory write error at address %04X: %v", addr, err)
			}
			continue
		}

		// Обработать секцию кода
		if codeSection {
			bytes, err := l.parseHexLine(line)
			if err != nil {
				return 0, fmt.Errorf("parse error at line %d: %v", lineNum, err)
			}

			if len(bytes) != 3 {
				return 0, fmt.Errorf("invalid command length at line %d: expected 3 bytes, got %d", lineNum, len(bytes))
			}

			// Записать байты в память
			for i, b := range bytes {
				addr := currentAddr + uint16(i)
				if err := memory.WriteByteAt(addr, b); err != nil {
					return 0, fmt.Errorf("memory write error at address %04X: %v", addr, err)
				}
			}
			currentAddr += 3
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("file read error: %v", err)
	}

	fmt.Printf("Program loaded successfully from %s at address %04X\n", filename, startAddr)
	return startAddr, nil
}

// parseHexLine парсит строку с шестнадцатеричными значениями в массив байт.
// line: строка с шестнадцатеричными значениями (6 символов)
// Возвращает массив из 3 байт или ошибку в случае невалидных данных.
func (l *Loader) parseHexLine(line string) ([]byte, error) {
	// Удалить все пробелы и преобразовать в верхний регистр
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ToUpper(line)
	
	if len(line) != 6 {
		return nil, fmt.Errorf("invalid line length: expected 6 hex characters, got %d", len(line))
	}

	bytes := make([]byte, 3)
	for i := 0; i < 3; i++ {
		hexByte := line[i*2 : i*2+2]
		val, err := strconv.ParseUint(hexByte, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("invalid hex value %s: %v", hexByte, err)
		}
		bytes[i] = byte(val)
	}

	return bytes, nil
}