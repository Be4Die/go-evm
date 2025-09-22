package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Be4Die/go-evm/internal/debugger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <program_file> [symbol_file]\n", os.Args[0])
		os.Exit(1)
	}

	programFile := os.Args[1]
	var symbolFile string

	if len(os.Args) > 2 {
		symbolFile = os.Args[2]
	} else {
		// Пытаемся найти файл символов по умолчанию
		base := programFile[:len(programFile)-len(filepath.Ext(programFile))]
		symbolFile = base + ".sym"
	}

	// Создаем отладчик
	dbg := debugger.NewDebugger()

	// Загружаем программу
	err := dbg.LoadProgram(programFile)
	if err != nil {
		fmt.Printf("Failed to load program: %v\n", err)
		os.Exit(1)
	}

	// Пытаемся загрузить символы, если файл существует
	if _, err := os.Stat(symbolFile); err == nil {
		err = dbg.LoadSymbols(symbolFile)
		if err != nil {
			fmt.Printf("Warning: failed to load symbols: %v\n", err)
		}
	}

	// Запускаем интерактивный режим
	dbg.StartInteractive()
}