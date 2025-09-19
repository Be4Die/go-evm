package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Be4Die/go-evm/loader"
	"github.com/Be4Die/go-evm/translator"
	"github.com/Be4Die/go-evm/vm"
)

func main() {
	// Определение флагов
	translateFlag := flag.Bool("translate", false, "Translate assembly file to bytecode")
	runAsmFlag := flag.Bool("run-asm", false, "Run assembly file (through temporary translation)")
	flag.Parse()

	if *translateFlag {
		// Режим трансляции: -translate <input.asm> <output.bin>
		if len(flag.Args()) != 2 {
			fmt.Printf("Usage: %s -translate <input_file.asm> <output_file.bin>\n", os.Args[0])
			fmt.Println("  input_file.asm - text file with assembly program")
			fmt.Println("  output_file.bin - output file for bytecode")
			os.Exit(1)
		}

		inputFile := flag.Arg(0)
		outputFile := flag.Arg(1)

		if err := translateFile(inputFile, outputFile); err != nil {
			fmt.Printf("Translation error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("File successfully translated: %s -> %s\n", inputFile, outputFile)
	} else if *runAsmFlag {
		// Режим запуска ассемблера: -run-asm <input.asm>
		if len(flag.Args()) != 1 {
			fmt.Printf("Usage: %s -run-asm <program_file.asm>\n", os.Args[0])
			fmt.Println("  program_file.asm - text file with assembly program")
			os.Exit(1)
		}

		filename := flag.Arg(0)
		if err := runAsmFile(filename); err != nil {
			fmt.Printf("Execution error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Стандартный режим: запуск байт-кода
		if len(flag.Args()) != 1 {
			fmt.Printf("Usage: %s <program_file.bin>\n", os.Args[0])
			fmt.Println("  program_file.bin - binary file with program in bytecode format")
			os.Exit(1)
		}

		filename := flag.Arg(0)
		if err := runBytecodeFile(filename); err != nil {
			fmt.Printf("Execution error: %v\n", err)
			os.Exit(1)
		}
	}
}

// translateFile транслирует файл с ассемблером в файл с байт-кодом
func translateFile(inputFile, outputFile string) error {
	// Чтение исходного файла
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Трансляция
	bytecode, err := translator.Translate(string(data))
	if err != nil {
		return fmt.Errorf("translation failed: %v", err)
	}

	// Запись байт-кода в выходной файл
	if err := os.WriteFile(outputFile, bytecode, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}

// runAsmFile запускает файл с ассемблером через временный файл
func runAsmFile(filename string) error {
	// Чтение исходного файла
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Трансляция
	bytecode, err := translator.Translate(string(data))
	if err != nil {
		return fmt.Errorf("translation failed: %v", err)
	}

	// Создание временного файла
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, "go-evm_temp.bin")
	defer os.Remove(tmpFile) // Удаляем временный файл после использования

	if err := os.WriteFile(tmpFile, bytecode, 0644); err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}

	// Запуск байт-кода
	return runBytecodeFile(tmpFile)
}

// runBytecodeFile запускает файл с байт-кодом
func runBytecodeFile(filename string) error {
	// Инициализация памяти (64KB)
	memory := vm.NewMemory(64 * 1024)

	// Инициализация CPU
	cpu := vm.NewCPU(memory)

	// Инициализация загрузчика
	ldr := loader.NewLoader()

	// Загрузка программы в память
	startAddr, err := ldr.LoadProgram(filename, memory)
	if err != nil {
		return fmt.Errorf("load error: %v", err)
	}

	// Установка начального IP
	cpu.GetPSW().SetIP(startAddr)
	fmt.Printf("Program loaded at %04X\n", startAddr)

	// Выполнение программы
	if err := cpu.Run(); err != nil {
		return fmt.Errorf("execution error: %v", err)
	}

	fmt.Println("Program terminated normally")
	fmt.Printf("Final IP: %04X, SP: %d\n", cpu.GetPSW().GetIP(), cpu.GetPSW().GetSP())
	return nil
}