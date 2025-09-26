package vm

import (
	"math"
	"testing"
)

func TestMovCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.push(0x12345678)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_MOV)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	value, err := mem.ReadWordAt(0x200)
	if err != nil {
		t.Fatal(err)
	}
	if value != 0x12345678 {
		t.Errorf("Expected 0x12345678, got 0x%08X", value)
	}
	if cpu.psw.GetSP() != 31 {
		t.Errorf("Expected SP 31, got %d", cpu.psw.GetSP())
	}
}

func TestAddIntCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 10)
	cpu.push(5)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_ADD_I)
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
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
	if cpu.psw.GetFlag(FLAG_ZERO) {
		t.Error("Zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_NEGATIVE) {
		t.Error("Negative flag should not be set")
	}
}

func TestSubIntCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 5)
	cpu.push(10)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_SUB_I)
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
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	if cpu.psw.GetFlag(FLAG_ZERO) {
		t.Error("Zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_NEGATIVE) {
		t.Error("Negative flag should not be set")
	}
}

func TestMulIntCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 5)
	cpu.push(10)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_MUL_I)
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
	if result != 50 {
		t.Errorf("Expected 50, got %d", result)
	}
	if cpu.psw.GetFlag(FLAG_ZERO) {
		t.Error("Zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_NEGATIVE) {
		t.Error("Negative flag should not be set")
	}
}

func TestDivIntCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 5)
	cpu.push(10)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_DIV_I)
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
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
	if cpu.psw.GetFlag(FLAG_ZERO) {
		t.Error("Zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_NEGATIVE) {
		t.Error("Negative flag should not be set")
	}
}

func TestAddFloatCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, Float32ToUint32(2.5))
	cpu.push(Float32ToUint32(3.5))
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
	if math.Abs(float64(Uint32ToFloat32(result)-6.0)) > 1e-6 {
		t.Errorf("Expected 6.0, got %f", Uint32ToFloat32(result))
	}
	if cpu.psw.GetFlag(FLAG_FZERO) {
		t.Error("Float zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_FNEGATIVE) {
		t.Error("Float negative flag should not be set")
	}
}

func TestSubFloatCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, Float32ToUint32(2.5))
	cpu.push(Float32ToUint32(3.5))
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_SUB_F)
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
	if math.Abs(float64(Uint32ToFloat32(result)-1.0)) > 1e-6 {
		t.Errorf("Expected 1.0, got %f", Uint32ToFloat32(result))
	}
	if cpu.psw.GetFlag(FLAG_FZERO) {
		t.Error("Float zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_FNEGATIVE) {
		t.Error("Float negative flag should not be set")
	}
}

func TestMulFloatCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, Float32ToUint32(2.5))
	cpu.push(Float32ToUint32(3.0))
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_MUL_F)
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
	if math.Abs(float64(Uint32ToFloat32(result)-7.5)) > 1e-6 {
		t.Errorf("Expected 7.5, got %f", Uint32ToFloat32(result))
	}
	if cpu.psw.GetFlag(FLAG_FZERO) {
		t.Error("Float zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_FNEGATIVE) {
		t.Error("Float negative flag should not be set")
	}
}

func TestDivFloatCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, Float32ToUint32(2.5))
	cpu.push(Float32ToUint32(5.0))
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_DIV_F)
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
	if math.Abs(float64(Uint32ToFloat32(result)-2.0)) > 1e-6 {
		t.Errorf("Expected 2.0, got %f", Uint32ToFloat32(result))
	}
	if cpu.psw.GetFlag(FLAG_FZERO) {
		t.Error("Float zero flag should not be set")
	}
	if cpu.psw.GetFlag(FLAG_FNEGATIVE) {
		t.Error("Float negative flag should not be set")
	}
}

func TestCmpIntCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 5)
	cpu.push(5)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_CMP_I)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if !cpu.psw.GetFlag(FLAG_ZERO) {
		t.Error("Zero flag should be set")
	}
	if cpu.psw.GetFlag(FLAG_NEGATIVE) {
		t.Error("Negative flag should not be set")
	}
	if cpu.psw.GetSP() != 31 {
		t.Errorf("Expected SP 31, got %d", cpu.psw.GetSP())
	}
}

func TestCmpFloatCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, Float32ToUint32(2.5))
	cpu.push(Float32ToUint32(2.5))
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_CMP_F)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if !cpu.psw.GetFlag(FLAG_FZERO) {
		t.Error("Float zero flag should be set")
	}
	if cpu.psw.GetFlag(FLAG_FNEGATIVE) {
		t.Error("Float negative flag should not be set")
	}
	if cpu.psw.GetSP() != 31 {
		t.Errorf("Expected SP 31, got %d", cpu.psw.GetSP())
	}
}

func TestJmpCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_JMP)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestJzCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetFlag(FLAG_ZERO, true)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_JZ)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestJnzCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetFlag(FLAG_ZERO, false)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_JNZ)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestJcCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetFlag(FLAG_CARRY, true)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_JC)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestJncCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetFlag(FLAG_CARRY, false)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_JNC)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestCallCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_CALL)
	mem.WriteByteAt(0x101, 0xAA)
	mem.WriteByteAt(0x102, 0xBB)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0xBBAA {
		t.Errorf("Expected IP 0xBBAA, got 0x%04X", cpu.psw.GetIP())
	}

	retAddr, err := cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if retAddr != 0x103 {
		t.Errorf("Expected return address 0x103, got 0x%04X", retAddr)
	}
}

func TestRetCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.push(0x200)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_RET)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x00)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetIP() != 0x200 {
		t.Errorf("Expected IP 0x200, got 0x%04X", cpu.psw.GetIP())
	}
}

func TestPushCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 0x12345678)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_PUSH)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	value, err := cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if value != 0x12345678 {
		t.Errorf("Expected 0x12345678, got 0x%08X", value)
	}
}

func TestPopCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.push(0x12345678)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_POP)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	value, err := mem.ReadWordAt(0x200)
	if err != nil {
		t.Fatal(err)
	}
	if value != 0x12345678 {
		t.Errorf("Expected 0x12345678, got 0x%08X", value)
	}
}

func TestAndCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 0x00FF00FF)
	cpu.push(0x0F0F0F0F)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_AND)
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
	if result != 0x000F000F {
		t.Errorf("Expected 0x000F000F, got 0x%08X", result)
	}
}

func TestOrCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 0x00FF00FF)
	cpu.push(0x0F0F0F0F)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_OR)
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
	if result != 0x0FFF0FFF {
		t.Errorf("Expected 0x0FFF0FFF, got 0x%08X", result)
	}
}

func TestXorCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 0x00FF00FF)
	cpu.push(0x0F0F0F0F)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_XOR)
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
	if result != 0x0FF00FF0 {
		t.Errorf("Expected 0x0FF00FF0, got 0x%08X", result)
	}
}

func TestNotCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	cpu.push(0xFFFF0000)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_NOT)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x00)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	result, err := cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if result != 0x0000FFFF {
		t.Errorf("Expected 0x0000FFFF, got 0x%08X", result)
	}
}

func TestShlCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 4)
	cpu.push(0x0000000F)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_SHL)
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
	if result != 0x000000F0 {
		t.Errorf("Expected 0x000000F0, got 0x%08X", result)
	}
}

func TestShrCommand(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	mem.WriteWordAt(0x200, 4)
	cpu.push(0xF0000000)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_SHR)
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
	if result != 0x0F000000 {
		t.Errorf("Expected 0x0F000000, got 0x%08X", result)
	}
}

func TestAddIntCommandWithOverflow(t *testing.T) {
    mem := NewMemory(1024)
    cpu := NewCPU(mem)
    cpu.psw.SetSP(31)

    mem.WriteWordAt(0x200, math.MaxUint32)
    cpu.push(1)
    cpu.psw.SetIP(0x100)
    mem.WriteByteAt(0x100, OP_ADD_I)
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
    if result != 0 {
        t.Errorf("Expected 0, got %d", result)
    }
    if !cpu.psw.GetFlag(FLAG_CARRY) {
        t.Error("Carry flag should be set")
    }
    if cpu.psw.GetFlag(FLAG_OVERFLOW) {
        t.Error("Overflow flag should not be set")
    }
}

func TestSetdAndClrdCommands(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)

	// Test SETD command
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_SETD)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x00)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if !cpu.psw.GetFlag(FLAG_DESTINATION) {
		t.Error("SETD command failed: DESTINATION flag not set")
	}

	// Test CLRD command
	cpu.psw.SetIP(0x110)
	mem.WriteByteAt(0x110, OP_CLRD)
	mem.WriteByteAt(0x111, 0x00)
	mem.WriteByteAt(0x112, 0x00)

	err = cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.psw.GetFlag(FLAG_DESTINATION) {
		t.Error("CLRD command failed: DESTINATION flag not cleared")
	}
}

func TestAddIntWithDestinationFlag(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	// Test with DESTINATION flag set (result in memory)
	cpu.psw.SetFlag(FLAG_DESTINATION, true)
	mem.WriteWordAt(0x200, 10)
	cpu.push(5)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_ADD_I)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x02)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	// Check result in memory
	result, err := mem.ReadWordAt(0x200)
	if err != nil {
		t.Fatal(err)
	}
	if result != 15 {
		t.Errorf("Expected 15 in memory, got %d", result)
	}

	// Test with DESTINATION flag cleared (result in stack)
	cpu.psw.SetFlag(FLAG_DESTINATION, false)
	mem.WriteWordAt(0x200, 20)
	cpu.push(10)
	cpu.psw.SetIP(0x110)
	mem.WriteByteAt(0x110, OP_ADD_I)
	mem.WriteByteAt(0x111, 0x00)
	mem.WriteByteAt(0x112, 0x02)

	err = cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	// Check result in stack
	result, err = cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if result != 30 {
		t.Errorf("Expected 30 in stack, got %d", result)
	}
}

func TestArithmeticCommandsWithDestinationFlag(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	commands := []struct {
		opcode byte
		name   string
		a      uint32
		b      uint32
		result uint32
	}{
		{OP_SUB_I, "SUB_I", 10, 5, 5},
		{OP_MUL_I, "MUL_I", 5, 4, 20},
		{OP_DIV_I, "DIV_I", 20, 5, 4},
		{OP_ADD_F, "ADD_F", Float32ToUint32(2.5), Float32ToUint32(1.5), Float32ToUint32(4.0)},
		{OP_SUB_F, "SUB_F", Float32ToUint32(5.0), Float32ToUint32(2.5), Float32ToUint32(2.5)},
		{OP_MUL_F, "MUL_F", Float32ToUint32(2.5), Float32ToUint32(4.0), Float32ToUint32(10.0)},
		{OP_DIV_F, "DIV_F", Float32ToUint32(10.0), Float32ToUint32(2.5), Float32ToUint32(4.0)},
	}

	for _, cmd := range commands {
		t.Run(cmd.name+"_MemoryDestination", func(t *testing.T) {
			cpu.psw.SetFlag(FLAG_DESTINATION, true)
			mem.WriteWordAt(0x200, cmd.b)
			cpu.push(cmd.a)
			cpu.psw.SetIP(0x100)
			mem.WriteByteAt(0x100, cmd.opcode)
			mem.WriteByteAt(0x101, 0x00)
			mem.WriteByteAt(0x102, 0x02)

			err := cpu.Step()
			if err != nil {
				t.Fatal(err)
			}

			result, err := mem.ReadWordAt(0x200)
			if err != nil {
				t.Fatal(err)
			}
			if result != cmd.result {
				t.Errorf("%s: expected %d, got %d", cmd.name, cmd.result, result)
			}
		})

		t.Run(cmd.name+"_StackDestination", func(t *testing.T) {
			cpu.psw.SetFlag(FLAG_DESTINATION, false)
			mem.WriteWordAt(0x200, cmd.b)
			cpu.push(cmd.a)
			cpu.psw.SetIP(0x110)
			mem.WriteByteAt(0x110, cmd.opcode)
			mem.WriteByteAt(0x111, 0x00)
			mem.WriteByteAt(0x112, 0x02)

			err := cpu.Step()
			if err != nil {
				t.Fatal(err)
			}

			result, err := cpu.pop()
			if err != nil {
				t.Fatal(err)
			}
			if result != cmd.result {
				t.Errorf("%s: expected %d, got %d", cmd.name, cmd.result, result)
			}
		})
	}
}

func TestLogicalCommandsWithDestinationFlag(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	commands := []struct {
		opcode byte
		name   string
		a      uint32
		b      uint32
		result uint32
	}{
		{OP_AND, "AND", 0x00FF00FF, 0x0F0F0F0F, 0x000F000F},
		{OP_OR, "OR", 0x00FF00FF, 0x0F0F0F0F, 0x0FFF0FFF},
		{OP_XOR, "XOR", 0x00FF00FF, 0x0F0F0F0F, 0x0FF00FF0},
		{OP_SHL, "SHL", 0x0000000F, 4, 0x000000F0},
		{OP_SHR, "SHR", 0xF0000000, 4, 0x0F000000},
	}

	for _, cmd := range commands {
		t.Run(cmd.name+"_MemoryDestination", func(t *testing.T) {
			cpu.psw.SetFlag(FLAG_DESTINATION, true)
			mem.WriteWordAt(0x200, cmd.b)
			cpu.push(cmd.a)
			cpu.psw.SetIP(0x100)
			mem.WriteByteAt(0x100, cmd.opcode)
			mem.WriteByteAt(0x101, 0x00)
			mem.WriteByteAt(0x102, 0x02)

			err := cpu.Step()
			if err != nil {
				t.Fatal(err)
			}

			result, err := mem.ReadWordAt(0x200)
			if err != nil {
				t.Fatal(err)
			}
			if result != cmd.result {
				t.Errorf("%s: expected %08X, got %08X", cmd.name, cmd.result, result)
			}
		})

		t.Run(cmd.name+"_StackDestination", func(t *testing.T) {
			cpu.psw.SetFlag(FLAG_DESTINATION, false)
			mem.WriteWordAt(0x200, cmd.b)
			cpu.push(cmd.a)
			cpu.psw.SetIP(0x110)
			mem.WriteByteAt(0x110, cmd.opcode)
			mem.WriteByteAt(0x111, 0x00)
			mem.WriteByteAt(0x112, 0x02)

			err := cpu.Step()
			if err != nil {
				t.Fatal(err)
			}

			result, err := cpu.pop()
			if err != nil {
				t.Fatal(err)
			}
			if result != cmd.result {
				t.Errorf("%s: expected %08X, got %08X", cmd.name, cmd.result, result)
			}
		})
	}
}

func TestNotCommandWithDestinationFlag(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	// NOT command always uses stack for result (as per current implementation)
	// Test both flag states to verify behavior
	cpu.push(0xFFFF0000)
	cpu.psw.SetFlag(FLAG_DESTINATION, true)
	cpu.psw.SetIP(0x100)
	mem.WriteByteAt(0x100, OP_NOT)
	mem.WriteByteAt(0x101, 0x00)
	mem.WriteByteAt(0x102, 0x00)

	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	result, err := cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if result != 0x0000FFFF {
		t.Errorf("NOT with DESTINATION=true: expected 0x0000FFFF, got 0x%08X", result)
	}

	cpu.push(0x00FF00FF)
	cpu.psw.SetFlag(FLAG_DESTINATION, false)
	cpu.psw.SetIP(0x110)
	mem.WriteByteAt(0x110, OP_NOT)
	mem.WriteByteAt(0x111, 0x00)
	mem.WriteByteAt(0x112, 0x00)

	err = cpu.Step()
	if err != nil {
		t.Fatal(err)
	}

	result, err = cpu.pop()
	if err != nil {
		t.Fatal(err)
	}
	if result != 0xFF00FF00 {
		t.Errorf("NOT with DESTINATION=false: expected 0xFF00FF00, got 0x%08X", result)
	}
}

func TestComparisonCommandsIgnoreDestinationFlag(t *testing.T) {
	mem := NewMemory(1024)
	cpu := NewCPU(mem)
	cpu.psw.SetSP(31)

	commands := []struct {
		opcode byte
		name   string
	}{
		{OP_CMP_I, "CMP_I"},
		{OP_CMP_F, "CMP_F"},
	}

	for _, cmd := range commands {
		t.Run(cmd.name, func(t *testing.T) {
			// Test with DESTINATION flag set
			cpu.psw.SetFlag(FLAG_DESTINATION, true)
			if cmd.name == "CMP_I" {
				mem.WriteWordAt(0x200, 5)
				cpu.push(5)
			} else {
				mem.WriteWordAt(0x200, Float32ToUint32(2.5))
				cpu.push(Float32ToUint32(2.5))
			}
			
			cpu.psw.SetIP(0x100)
			mem.WriteByteAt(0x100, cmd.opcode)
			mem.WriteByteAt(0x101, 0x00)
			mem.WriteByteAt(0x102, 0x02)

			err := cpu.Step()
			if err != nil {
				t.Fatal(err)
			}

			// Verify SP is restored (comparison shouldn't leave result)
			if cpu.psw.GetSP() != 31 {
				t.Errorf("%s: SP not restored properly, got %d", cmd.name, cpu.psw.GetSP())
			}
		})
	}
}