package vm

import "testing"

func TestMemoryReadWrite(t *testing.T) {
    mem := NewMemory(1024)
    
    // Test byte operations
    testAddr := uint16(0x10)
    testValue := byte(0xAB)
    err := mem.WriteByteAt(testAddr, testValue)
    if err != nil {
        t.Fatal(err)
    }
    
    val, err := mem.ReadByteAt(testAddr)
    if err != nil {
        t.Fatal(err)
    }
    if val != testValue {
        t.Errorf("Expected %02x, got %02x", testValue, val)
    }

    // Test word operations
    testWord := uint32(0x12345678)
    err = mem.WriteWordAt(testAddr, testWord)
    if err != nil {
        t.Fatal(err)
    }
    
    word, err := mem.ReadWordAt(testAddr)
    if err != nil {
        t.Fatal(err)
    }
    if word != testWord {
        t.Errorf("Expected %08x, got %08x", testWord, word)
    }
}

func TestMemoryBounds(t *testing.T) {
    mem := NewMemory(100)
    
    // Test out of bounds
    err := mem.WriteByteAt(100, 0xFF)
    if err == nil {
        t.Error("Expected error for out of bounds write")
    }

    _, err = mem.ReadWordAt(99)
    if err == nil {
        t.Error("Expected error for out of bounds read")
    }
}