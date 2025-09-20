package main

import (
	"fmt"
	"os"

	"github.com/Be4Die/go-evm/loader"
	"github.com/Be4Die/go-evm/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <program_file>\n", os.Args[0])
		fmt.Println("  program_file - text file with program in hex format")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Initialize memory (64KB)
	memory := vm.NewMemory(64 * 1024)
	
	// Initialize CPU
	cpu := vm.NewCPU(memory)
	
	// Initialize loader
	ldr := loader.NewLoader()

	// Load program into memory
	startAddr, err := ldr.LoadProgram(filename, memory)
	if err != nil {
		fmt.Printf("Load error: %v\n", err)
		os.Exit(1)
	}

	// Set initial IP to start address
	cpu.GetPSW().SetIP(startAddr)
	fmt.Printf("Program loaded at %04X\n", startAddr)

	// Execute program
	if err := cpu.Run(); err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Program terminated normally")
	fmt.Printf("Final IP: %04X, SP: %d\n", cpu.GetPSW().GetIP(), cpu.GetPSW().GetSP())
}