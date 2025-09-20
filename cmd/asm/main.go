// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Be4Die/go-evm/internal/translator"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: asm <input.asm> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	translator := translator.NewTranslator()
	if err := translator.Assemble(inputFile, outputFile); err != nil {
		log.Fatalf("Assembly failed: %v", err)
	}

	fmt.Printf("Assembly completed successfully. Output written to %s\n", outputFile)
}