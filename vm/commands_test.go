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