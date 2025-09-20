package assembly

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (t *Translator) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	currentSection := "" // "data" или "code"

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Пропуск пустых строк и комментариев
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "//") {
			continue
		}

		// Удаление inline-комментариев
		line = removeComments(line)

		// Обработка директивы entry
		if strings.HasPrefix(strings.ToUpper(line), "ENTRY ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				return fmt.Errorf("invalid entry directive at line %d", lineNum)
			}
			if t.pass == 1 {
				t.entryLabel = parts[1]
			}
			continue
		}

		// Обработка директив секций
		if strings.HasPrefix(strings.ToUpper(line), "SECTION ") {
			section := strings.TrimSpace(line[8:])
			switch section {
			case ".data":
				currentSection = "data"
				t.currentAddress = t.dataAddress
			case ".code":
				currentSection = "code"
				t.currentAddress = 0x0200
			default:
				return fmt.Errorf("unknown section: %s at line %d", section, lineNum)
			}
			continue
		}

		// Обработка секции данных
		if currentSection == "data" {
			if err := t.processDataLine(line, lineNum); err != nil {
				return err
			}
			continue
		}

		// Обработка секции кода
		if currentSection == "code" {
			if err := t.processCodeLine(line, lineNum); err != nil {
				return err
			}
			continue
		}

		return fmt.Errorf("line %d is not in any section", lineNum)
	}

	return scanner.Err()
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

	// Обработка директивы EQU
	if strings.HasPrefix(strings.ToUpper(line), "EQU ") {
		return t.processEQU(line, lineNum)
	}

	// Обработка инструкций
	return t.processInstruction(line, lineNum)
}

// removeComments удаляет комментарии из строки
func removeComments(line string) string {
	if commentIndex := strings.Index(line, ";"); commentIndex != -1 {
		line = strings.TrimSpace(line[:commentIndex])
	}
	if commentIndex := strings.Index(line, "//"); commentIndex != -1 {
		line = strings.TrimSpace(line[:commentIndex])
	}
	return line
}