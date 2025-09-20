package vm

import (
	"testing"
)

func TestIntegration(t *testing.T) {
	mem := NewMemory(64 * 1024)
	cpu := NewCPU(mem)

	mem.WriteWordAt(0x100, Float32ToUint32(42.0))
	mem.WriteWordAt(0x104, Float32ToUint32(3.14))
	mem.WriteWordAt(0x108, 0xFFFFFFFF)

	program := []byte{
		OP_PUSH, 0x00, 0x01,  // PUSH [0x0100] (42.0)
		OP_ADD_F, 0x04, 0x01, // ADD_F [0x0104] (3.14)
		OP_POP, 0x00, 0x02,   // POP [0x0200] (результат 45.14)
		OP_PUSH, 0x08, 0x01,  // PUSH [0x0108] (0xFFFFFFFF)
		OP_NOT,               // NOT
		OP_POP, 0x04, 0x02,   // POP [0x0204] (результат NOT)
	}

	for i, b := range program {
		mem.WriteByteAt(0x200+uint16(i), b)
	}

	cpu.psw.SetIP(0x200)

	for i := 0; i < 6; i++ {
		if err := cpu.Step(); err != nil {
			t.Fatalf("Step %d failed: %v", i, err)
		}
	}

	result, err := mem.ReadWordAt(0x200)
	if err != nil {
		t.Fatal(err)
	}
	
	expected := Float32ToUint32(45.14)
	if result != expected {
		actualFloat := Uint32ToFloat32(result)
		expectedFloat := Uint32ToFloat32(expected)
		t.Errorf("Float addition failed. Expected %f (%08X), got %f (%08X)", 
			expectedFloat, expected, actualFloat, result)
	}

	result, err = mem.ReadWordAt(0x204)
	if err != nil {
		t.Fatal(err)
	}
	if result != 0x00000000 {
		t.Errorf("NOT operation failed. Expected 00000000, got %08X", result)
	}

	if cpu.psw.GetSP() != 31 {
		t.Errorf("Stack pointer not restored. Expected 31, got %d", cpu.psw.GetSP())
	}
}

func TestFibonacciProgram(t *testing.T) {
    mem := NewMemory(64 * 1024)
    cpu := NewCPU(mem)
    
    program := []byte{
        OP_PUSH, 0x00, 0x03, // PUSH 0
        OP_POP, 0x00, 0x10,   // POP [0x1000] - a
        OP_PUSH, 0x01, 0x03,  // PUSH 1
        OP_POP, 0x04, 0x10,   // POP [0x1004] - b
        OP_PUSH, 0x0A, 0x03,  // PUSH 10 (счетчик)
        
        // loop:
        OP_PUSH, 0x00, 0x10,  // PUSH [0x1000]
        OP_PUSH, 0x04, 0x10,  // PUSH [0x1004]
        OP_ADD_I, 0x00, 0x00, // ADD
        OP_POP, 0x08, 0x10,   // POP [0x1008] - c
        OP_PUSH, 0x04, 0x10,  // PUSH b
        OP_POP, 0x00, 0x10,   // POP a
        OP_PUSH, 0x08, 0x10,  // PUSH c
        OP_POP, 0x04, 0x10,   // POP b
        OP_PUSH, 0x08, 0x03,  // PUSH counter
        OP_PUSH, 0x01, 0x03,  // PUSH 1
        OP_SUB_I, 0x00, 0x00, // DEC counter
        OP_POP, 0x08, 0x03,   // POP counter
        OP_PUSH, 0x08, 0x03,  // PUSH counter
        OP_JNZ, 0x0F, 0x02,   // JNZ loop (абсолютный адрес 0x020F)
        
        OP_PUSH, 0x04, 0x10,  // PUSH result
        OP_OUT, 0x00, 0x00,   // OUTPUT
        OP_JMP, 0x00, 0x00,   // JMP 0 (остановка)
    }
    
    // Загрузка программы в память
    for i, b := range program {
        mem.WriteByteAt(0x200+uint16(i), b)
    }
    
    cpu.psw.SetIP(0x200)
    err := cpu.Run()
    if err != nil {
        t.Fatal(err)
    }
    
    // Проверяем что программа завершилась корректно
    if cpu.psw.GetIP() != 0 {
        t.Errorf("Expected IP 0, got %04X", cpu.psw.GetIP())
    }
}

func TestFactorialProgram(t *testing.T) {
    mem := NewMemory(64 * 1024)
    cpu := NewCPU(mem)
    
    program := []byte{
        OP_PUSH, 0x05, 0x03,  // PUSH 5
        OP_POP, 0x00, 0x10,    // POP [0x1000] - n
        OP_PUSH, 0x01, 0x03,   // PUSH 1
        OP_POP, 0x04, 0x10,    // POP [0x1004] - result
        
        // loop:
        OP_PUSH, 0x00, 0x10,   // PUSH n
        OP_PUSH, 0x04, 0x10,   // PUSH result
        OP_MUL_I, 0x00, 0x00,  // MUL
        OP_POP, 0x04, 0x10,    // POP result
        OP_PUSH, 0x00, 0x10,   // PUSH n
        OP_PUSH, 0x01, 0x03,   // PUSH 1
        OP_SUB_I, 0x00, 0x00,  // DEC n
        OP_POP, 0x00, 0x10,    // POP n
        OP_PUSH, 0x00, 0x10,   // PUSH n
        OP_JNZ, 0x0F, 0x02,    // JNZ loop (абсолютный адрес 0x020F)
        
        OP_PUSH, 0x04, 0x10,   // PUSH result
        OP_OUT, 0x00, 0x00,    // OUTPUT
        OP_JMP, 0x00, 0x00,    // JMP 0 (остановка)
    }
    
    for i, b := range program {
        mem.WriteByteAt(0x200+uint16(i), b)
    }
    
    cpu.psw.SetIP(0x200)
    err := cpu.Run()
    if err != nil {
        t.Fatal(err)
    }
}

func TestStackOverflow(t *testing.T) {
    mem := NewMemory(1024)
    cpu := NewCPU(mem)
    
    cpu.psw.SetSP(0)
    for i := 0; i < 31; i++ {
        cpu.push(uint32(i))
    }
    
    err := cpu.push(42)
    if err == nil {
        t.Error("Expected stack overflow error")
    }
}

func TestInvalidOpcode(t *testing.T) {
    mem := NewMemory(1024)
    cpu := NewCPU(mem)
    
    mem.WriteByteAt(0x100, 0xFF)
    cpu.psw.SetIP(0x100)
    
    err := cpu.Step()
    if err == nil {
        t.Error("Expected invalid opcode error")
    }
}

func TestDivisionByZero(t *testing.T) {
    mem := NewMemory(1024)
    cpu := NewCPU(mem)
    
    mem.WriteWordAt(0x200, 0)
    cpu.push(10)
    cpu.psw.SetIP(0x100)
    mem.WriteByteAt(0x100, OP_DIV_I)
    mem.WriteByteAt(0x101, 0x00)
    mem.WriteByteAt(0x102, 0x02)
    
    err := cpu.Step()
    if err == nil {
        t.Error("Expected division by zero error")
    }
    
    mem.WriteWordAt(0x200, Float32ToUint32(0.0))
    cpu.push(Float32ToUint32(10.0))
    cpu.psw.SetIP(0x100)
    mem.WriteByteAt(0x100, OP_DIV_F)
    mem.WriteByteAt(0x101, 0x00)
    mem.WriteByteAt(0x102, 0x02)
    
    err = cpu.Step()
    if err == nil {
        t.Error("Expected division by zero error")
    }
}

func TestSpecialFloatValues(t *testing.T) {
    mem := NewMemory(1024)
    cpu := NewCPU(mem)
    
    nan := uint32(0x7F800001)
    mem.WriteWordAt(0x200, nan)
    cpu.push(Float32ToUint32(1.0))
    cpu.psw.SetIP(0x100)
    mem.WriteByteAt(0x100, OP_ADD_F)
    mem.WriteByteAt(0x101, 0x00)
    mem.WriteByteAt(0x102, 0x02)
    
    err := cpu.Step()
    if err != nil {
        t.Fatal(err)
    }
    
    result, err := cpu.pop()
    if err != nil {
        t.Fatal(err)
    }
    if result>>23 != 0xFF {
        t.Error("Expected NaN result")
    }
}