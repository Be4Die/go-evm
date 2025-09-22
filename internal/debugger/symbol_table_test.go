package debugger

import (
	"os"
	"testing"
)

func TestSymbolTable(t *testing.T) {
    st := NewSymbolTable()

    // Create test symbol file
    testData := `0x1000 main
0x2000 loop
0x3000 exit`

    tmpfile, err := os.CreateTemp("", "test_symbols")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())

    if _, err := tmpfile.WriteString(testData); err != nil {
        t.Fatal(err)
    }
    tmpfile.Close()

    // Test LoadFromFile
    err = st.LoadFromFile(tmpfile.Name())
    if err != nil {
        t.Errorf("Failed to load symbols: %v", err)
    }

    // Test ResolveAddress
    addr, exists := st.ResolveAddress("main")
    if !exists || addr != 0x1000 {
        t.Error("Failed to resolve address")
    }

    // Test ResolveLabel
    label, exists := st.ResolveLabel(0x2000)
    if !exists || label != "loop" {
        t.Error("Failed to resolve label")
    }

    // Test non-existent symbol
    _, exists = st.ResolveAddress("nonexistent")
    if exists {
        t.Error("Non-existent symbol should not be found")
    }
}