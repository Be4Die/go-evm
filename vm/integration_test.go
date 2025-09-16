package vm

import (
	"testing"
)

func TestIntegration(t *testing.T) {
	mem := NewMemory(64 * 1024)
	cpu := NewCPU(mem)

	// Записываем тестовые данные в память
	mem.WriteWordAt(0x100, Float32ToUint32(42.0)) // Дробное число
	mem.WriteWordAt(0x104, Float32ToUint32(3.14)) // Дробное число
	mem.WriteWordAt(0x108, 0xFFFFFFFF)            // Максимальное unsigned

	// Загружаем программу напрямую в память (имитация загрузчика)
	program := []byte{
		OP_PUSH, 0x00, 0x01, // PUSH [0x0100] (42.0)
		OP_ADD_F, 0x04, 0x01, // ADD_F [0x0104] (3.14)
		OP_POP, 0x00, 0x02,  // POP [0x0200] (результат 45.14)
		OP_PUSH, 0x08, 0x01, // PUSH [0x0108] (0xFFFFFFFF)
		OP_NOT,                // NOT
		OP_POP, 0x04, 0x02,  // POP [0x0204] (результат NOT)
	}

	for i, b := range program {
		mem.WriteByteAt(0x200+uint16(i), b)
	}

	// Устанавливаем начальный IP
	cpu.psw.SetIP(0x200)

	// Выполняем программу
	for i := 0; i < 6; i++ {
		if err := cpu.Step(); err != nil {
			t.Fatalf("Step %d failed: %v", i, err)
		}
	}

	// Проверяем результаты
	// 1. Проверка дробного сложения
	result, err := mem.ReadWordAt(0x200)
	if err != nil {
		t.Fatal(err)
	}
	
	// Ожидаемый результат: 42.0 + 3.14 = 45.14
	expected := Float32ToUint32(45.14)
	if result != expected {
		actualFloat := Uint32ToFloat32(result)
		expectedFloat := Uint32ToFloat32(expected)
		t.Errorf("Float addition failed. Expected %f (%08X), got %f (%08X)", 
			expectedFloat, expected, actualFloat, result)
	}

	// 2. Проверка битовой операции
	result, err = mem.ReadWordAt(0x204)
	if err != nil {
		t.Fatal(err)
	}
	if result != 0x00000000 {
		t.Errorf("NOT operation failed. Expected 00000000, got %08X", result)
	}

	// 3. Проверка состояния регистров
	if cpu.psw.GetSP() != 31 {
		t.Errorf("Stack pointer not restored. Expected 31, got %d", cpu.psw.GetSP())
	}
}