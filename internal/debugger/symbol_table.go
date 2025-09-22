package debugger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SymbolTable представляет таблицу символов (метки -> адреса)
type SymbolTable struct {
	Symbols map[string]uint16
	Labels  map[uint16]string
}

// NewSymbolTable создает новую таблицу символов
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Symbols: make(map[string]uint16),
		Labels:  make(map[uint16]string),
	}
}

// LoadFromFile загружает таблицу символов из файла
func (st *SymbolTable) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open symbol file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		addr, err := strconv.ParseUint(parts[0], 0, 16)
		if err != nil {
			continue
		}

		label := parts[1]
		st.Symbols[label] = uint16(addr)
		st.Labels[uint16(addr)] = label
	}

	return scanner.Err()
}

// ResolveAddress разрешает имя метки в адрес
func (st *SymbolTable) ResolveAddress(label string) (uint16, bool) {
	addr, exists := st.Symbols[label]
	return addr, exists
}

// ResolveLabel разрешает адрес в имя метки
func (st *SymbolTable) ResolveLabel(addr uint16) (string, bool) {
	label, exists := st.Labels[addr]
	return label, exists
}

func (st *SymbolTable) GetAllSymbols() map[string]uint16 {
    return st.Symbols
}