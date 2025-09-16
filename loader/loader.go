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

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

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
	codeSection := false

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		// Remove any inline comments
		if commentIndex := strings.Index(line, "//"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}
		if commentIndex := strings.Index(line, "#"); commentIndex != -1 {
			line = strings.TrimSpace(line[:commentIndex])
		}

		// Check for START directive
		if strings.HasPrefix(line, "START:") {
			addrStr := strings.TrimSpace(line[len("START:"):])
			addr, err := strconv.ParseUint(addrStr, 0, 16)
			if err != nil {
				return 0, fmt.Errorf("invalid start address: %v", err)
			}
			startAddr = uint16(addr)
			currentAddr = startAddr
			codeSection = true
			continue
		}

		// Check for DATA directive
		if strings.HasPrefix(line, "DATA:") {
			parts := strings.Fields(line[len("DATA:"):])
			if len(parts) < 2 {
				return 0, fmt.Errorf("invalid DATA directive at line %d", lineNum)
			}
			addrStr := parts[0]
			valueStr := parts[1]

			addr, err := strconv.ParseUint(addrStr, 0, 16)
			if err != nil {
				return 0, fmt.Errorf("invalid address in DATA directive at line %d: %v", lineNum, err)
			}

			var value uint32
			// Try to parse as float first
			if strings.Contains(valueStr, ".") {
				floatVal, err := strconv.ParseFloat(valueStr, 32)
				if err != nil {
					return 0, fmt.Errorf("invalid float value in DATA directive at line %d: %v", lineNum, err)
				}
				value = math.Float32bits(float32(floatVal))
			} else {
				// Parse as integer
				intVal, err := strconv.ParseUint(valueStr, 0, 32)
				if err != nil {
					return 0, fmt.Errorf("invalid integer value in DATA directive at line %d: %v", lineNum, err)
				}
				value = uint32(intVal)
			}

			if err := memory.WriteWordAt(uint16(addr), value); err != nil {
				return 0, fmt.Errorf("memory write error at address %04X: %v", addr, err)
			}
			continue
		}

		// If we haven't encountered a START directive yet, assume code starts at 0
		if !codeSection {
			codeSection = true
			currentAddr = 0
		}

		// Parse hex values from line (command)
		bytes, err := l.parseHexLine(line)
		if err != nil {
			return 0, fmt.Errorf("parse error at line %d: %v", lineNum, err)
		}

		if len(bytes) != 3 {
			return 0, fmt.Errorf("invalid command length at line %d: expected 3 bytes, got %d", lineNum, len(bytes))
		}

		// Write bytes to memory
		for i, b := range bytes {
			addr := currentAddr + uint16(i)
			if err := memory.WriteByteAt(addr, b); err != nil {
				return 0, fmt.Errorf("memory write error at address %04X: %v", addr, err)
			}
		}
		currentAddr += 3
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("file read error: %v", err)
	}

	fmt.Printf("Program loaded successfully from %s at address %04X\n", filename, startAddr)
	return startAddr, nil
}

func (l *Loader) parseHexLine(line string) ([]byte, error) {
	// Remove all whitespace and convert to uppercase
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