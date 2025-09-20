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
	inDataSection := false

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Пропуск пустых строк и комментариев
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "//") {
			continue
		}

		// Удаление inline-комментариев
		line = removeComments(line)

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