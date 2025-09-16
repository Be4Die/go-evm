package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Be4Die/go-evm/loader"
	"github.com/Be4Die/go-evm/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <program_file> [base_address]\n", os.Args[0])
		fmt.Println("  program_file - text file with program in hex format")
		fmt.Println("  base_address - optional memory address to load program (default: 0x1000)")
		os.Exit(1)
	}

	filename := os.Args[1]
	baseAddr := uint16(0x1000)

	if len(os.Args) > 2 {
		addr, err := strconv.ParseUint(os.Args[2], 0, 16)
		if err != nil {
			fmt.Printf("Invalid base address: %v\n", err)
			os.Exit(1)
		}
		baseAddr = uint16(addr)
	}

	// Initialize memory (64KB)
	memory := vm.NewMemory(64 * 1024)
	
	// Initialize CPU
	cpu := vm.NewCPU(memory)
	
	// Initialize loader
	ldr := loader.NewLoader()

	// Load program into memory
	startAddr, err := ldr.LoadProgram(filename, baseAddr, memory)
	if err != nil {
		fmt.Printf("Load error: %v\n", err)
		os.Exit(1)
	}

	// Set initial IP to start address
	cpu.GetPSW().SetIP(startAddr)
	fmt.Printf("Program loaded at %04X\n", startAddr)

	// Execute program
	maxInstructions := 1000
	instructionCount := 0
	
	for instructionCount < maxInstructions {
		if err := cpu.Step(); err != nil {
			fmt.Printf("Execution error: %v\n", err)
			break
		}
		
		instructionCount++
		
		// Check if we've reached the end of program
		if cpu.GetPSW().GetIP() == 0x0000 {
			fmt.Println("Program terminated normally")
			break
		}
	}
	
	if instructionCount >= maxInstructions {
		fmt.Println("Reached maximum instruction limit")
	}
	
	fmt.Printf("Execution finished after %d instructions\n", instructionCount)
	fmt.Printf("Final IP: %04X, SP: %d\n", cpu.GetPSW().GetIP(), cpu.GetPSW().GetSP())
}