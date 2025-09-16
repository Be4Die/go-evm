package loader

import (
	"os"
	"testing"

	"github.com/Be4Die/go-evm/vm"
)

func TestLoader(t *testing.T) {
	// Создаем временный файл с программой
	program := "130001\n130401\n060000\n140002\n130801\n1A0000\n140402"
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

	if err := ldr.LoadProgram(tmpfile.Name(), 0x1000, mem); err != nil {
		t.Fatal(err)
	}

	// Проверяем загруженные данные
	expected := []byte{0x13, 0x00, 0x01, 0x13, 0x04, 0x01, 0x06, 0x00, 0x00, 0x14, 0x00, 0x02}
	for i, b := range expected {
		val, err := mem.ReadByteAt(0x1000 + uint16(i))
		if err != nil {
			t.Fatal(err)
		}
		if val != b {
			t.Errorf("Byte %d mismatch. Expected %02X, got %02X", i, b, val)
		}
	}
}