package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Be4Die/go-evm/internal/assembly"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Generate debug symbol file")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: asm [--debug] <input.asm> <output.txt>")
		os.Exit(1)
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	translator := assembly.NewTranslator()
	if err := translator.Assemble(inputFile, outputFile, *debugFlag); err != nil {
		log.Fatalf("Assembly failed: %v", err)
	}

	fmt.Printf("Assembly completed successfully. Output written to %s\n", outputFile)
	if *debugFlag {
		symFile := strings.TrimSuffix(outputFile, filepath.Ext(outputFile)) + ".sym"
		fmt.Printf("Debug symbols written to %s\n", symFile)
	}
}