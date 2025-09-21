package assembly

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewTranslator(t *testing.T) {
	translator := NewTranslator()

	if translator.symbolTable == nil {
		t.Error("symbolTable not initialized")
	}

	if translator.constants == nil {
		t.Error("constants not initialized")
	}

	if translator.dataAddress != 0x0100 {
		t.Errorf("dataAddress = %x, want %x", translator.dataAddress, 0x0100)
	}
}

func TestAssemble(t *testing.T) {
	// Create a temporary test file
	testAsm := `
ENTRY START
SECTION .data
DATA: DB 10, 20

SECTION .code
START:
    MOV 0x10
    JMP START
`

	tmpDir := t.TempDir()
	asmFile := filepath.Join(tmpDir, "test.asm")
	binFile := filepath.Join(tmpDir, "test.bin")

	err := os.WriteFile(asmFile, []byte(testAsm), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	translator := NewTranslator()
	err = translator.Assemble(asmFile, binFile, false)
	if err != nil {
		t.Fatalf("Assemble failed: %v", err)
	}

	// Check if output file was created
	if _, err := os.Stat(binFile); err != nil {
		t.Error("Output file was not created")
	}

	// Check if entry label was set
	if translator.entryLabel != "START" {
		t.Errorf("entryLabel = %s, want %s", translator.entryLabel, "START")
	}

	// Check if symbols were added to symbol table
	if _, exists := translator.symbolTable["START"]; !exists {
		t.Error("START label not found in symbol table")
	}

	if _, exists := translator.symbolTable["DATA"]; !exists {
		t.Error("DATA label not found in symbol table")
	}
}


func TestAssembleWithDebug(t *testing.T) {
	// Создаем временный тестовый файл
	testAsm := `
ENTRY START
SECTION .data
DATA: DB 10, 20

SECTION .code
START:
    MOV 0x10
    JMP START
`

	tmpDir := t.TempDir()
	asmFile := filepath.Join(tmpDir, "test.asm")
	binFile := filepath.Join(tmpDir, "test.bin")

	err := os.WriteFile(asmFile, []byte(testAsm), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	translator := NewTranslator()
	err = translator.Assemble(asmFile, binFile, true) // Включаем debug режим
	if err != nil {
		t.Fatalf("Assemble failed: %v", err)
	}

	// Проверяем что файл символов был создан
	symFile := strings.TrimSuffix(binFile, filepath.Ext(binFile)) + ".sym"
	if _, err := os.Stat(symFile); err != nil {
		t.Error("Symbol file was not created")
	}

	// Проверяем содержимое файла символов
	content, err := os.ReadFile(symFile)
	if err != nil {
		t.Fatalf("Failed to read symbol file: %v", err)
	}

	expectedContent := "0x0100 DATA\n0x0200 START\n"
	if string(content) != expectedContent {
		t.Errorf("Symbol file content = %s, want %s", string(content), expectedContent)
	}
}

func TestWriteOutput(t *testing.T) {
	translator := NewTranslator()
	translator.startAddress = 0x0200
	translator.entryLabel = "START"
	translator.symbolTable["START"] = 0x0200
	
	// Add some test data
	translator.data = []dataItem{
		{address: 0x0100, value: 10, isByte: true},
		{address: 0x0101, value: 20, isByte: true},
	}
	
	// Add some test instructions
	translator.instructions = []instruction{
		{address: 0x0200, opcode: 0x01, operand: 0x0010},
		{address: 0x0203, opcode: 0x02, operand: 0x0020},
	}

	tmpDir := t.TempDir()
	outputFile := filepath.Join(tmpDir, "output.bin")

	err := translator.writeOutput(outputFile)
	if err != nil {
		t.Fatalf("writeOutput failed: %v", err)
	}

	// Check if file was created and has content
	info, err := os.Stat(outputFile)
	if err != nil {
		t.Fatalf("Failed to stat output file: %v", err)
	}

	if info.Size() == 0 {
		t.Error("Output file is empty")
	}
}