package loader

import (
	"math"
	"os"
	"testing"

	"github.com/Be4Die/go-evm/vm"
)

func TestLoader(t *testing.T) {
	// Создаем временный файл с программой в новом формате
	program := "0x1000\nDS\n0x0100 42.0\n0x0104 3.14\nDE\n130001\n130401\n060000\n140002\n130801\n1A0000\n140402"
	tmpfile, err := os.CreateTemp("", "test_program")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(program); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Тестируем загрузчик
	mem := vm.NewMemory(64 * 1024)
	ldr := NewLoader()

	startAddr, err := ldr.LoadProgram(tmpfile.Name(), mem)
	if err != nil {
		t.Fatal(err)
	}

	if startAddr != 0x1000 {
		t.Errorf("Expected start address 0x1000, got %04X", startAddr)
	}

	// Проверяем загруженные данные
	expectedData := map[uint16]uint32{
		0x0100: math.Float32bits(42.0),
		0x0104: math.Float32bits(3.14),
	}

	for addr, expectedValue := range expectedData {
		val, err := mem.ReadWordAt(addr)
		if err != nil {
			t.Fatal(err)
		}
		if val != expectedValue {
			t.Errorf("Data at %04X mismatch. Expected %08X, got %08X", addr, expectedValue, val)
		}
	}

	// Проверяем загруженный код
	expectedCode := []byte{0x13, 0x00, 0x01, 0x13, 0x04, 0x01, 0x06, 0x00, 0x00, 0x14, 0x00, 0x02}
	for i, b := range expectedCode {
		val, err := mem.ReadByteAt(0x1000 + uint16(i))
		if err != nil {
			t.Fatal(err)
		}
		if val != b {
			t.Errorf("Byte %d mismatch. Expected %02X, got %02X", i, b, val)
		}
	}
}