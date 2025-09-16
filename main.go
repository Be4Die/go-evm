// main.go
package main

import (
	"fmt"
	"os"

	"github.com/Be4Die/go-evm/loader"
	"github.com/Be4Die/go-evm/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vm <program_file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	memory := vm.NewMemory(65535) // 64KB - 1 (максимальный адрес для uint16)
	cpu := vm.NewCPU(memory)

	// Load program at address 0x1000
	if err := loader.LoadProgram(filename, 0x1000, memory); err != nil {
		fmt.Printf("Load error: %v\n", err)
		os.Exit(1)
	}

	// Set initial IP
	cpu.GetPSW().SetIP(0x1000)

	// Run program
	if err := cpu.Run(); err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}
}