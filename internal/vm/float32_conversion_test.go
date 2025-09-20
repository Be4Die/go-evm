package vm

import (
	"math"
	"testing"
)

func TestFloat32ToUint32(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected uint32
	}{
		{"Zero", 0.0, 0x00000000},
		{"PositiveOne", 1.0, 0x3F800000},
		{"NegativeOne", -1.0, 0xBF800000},
		{"PositiveInfinity", float32(math.Inf(1)), 0x7F800000},
		{"NegativeInfinity", float32(math.Inf(-1)), 0xFF800000},
		{"NaN", float32(math.NaN()), 0x7FC00000}, // Стандартный NaN
		{"MaxFloat32", math.MaxFloat32, 0x7F7FFFFF},
		{"SmallestNormalFloat32", 1.17549435e-38, 0x00800000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float32ToUint32(tt.input)
			
			// Особый случай для NaN
			if tt.name == "NaN" {
				// Проверяем что это NaN (экспонента вся 1, мантисса не нулевая)
				if (result&0x7F800000) != 0x7F800000 || (result&0x007FFFFF) == 0 {
					t.Errorf("Float32ToUint32(%f) = %#x, expected NaN pattern", tt.input, result)
				}
			} else if result != tt.expected {
				t.Errorf("Float32ToUint32(%f) = %#x, expected %#x", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUint32ToFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected float32
	}{
		{"Zero", 0x00000000, 0.0},
		{"PositiveOne", 0x3F800000, 1.0},
		{"NegativeOne", 0xBF800000, -1.0},
		{"PositiveInfinity", 0x7F800000, float32(math.Inf(1))},
		{"NegativeInfinity", 0xFF800000, float32(math.Inf(-1))},
		{"NaN", 0x7FC00000, float32(math.NaN())},
		{"MaxFloat32", 0x7F7FFFFF, math.MaxFloat32},
		{"SmallestNormalFloat32", 0x00800000, 1.17549435e-38},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint32ToFloat32(tt.input)
			
			// Особый случай для NaN
			if tt.name == "NaN" {
				if !math.IsNaN(float64(result)) {
					t.Errorf("Uint32ToFloat32(%#x) = %f, expected NaN", tt.input, result)
				}
			} else if tt.name == "PositiveInfinity" || tt.name == "NegativeInfinity" {
				// Проверка бесконечностей
				if math.IsInf(float64(result), 0) != math.IsInf(float64(tt.expected), 0) {
					t.Errorf("Uint32ToFloat32(%#x) = %f, expected %f", tt.input, result, tt.expected)
				}
			} else if result != tt.expected {
				t.Errorf("Uint32ToFloat32(%#x) = %f, expected %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConversionRoundTrip(t *testing.T) {
	testValues := []float32{
		0.0,
		1.0,
		-1.0,
		math.Pi,
		math.MaxFloat32,
		math.SmallestNonzeroFloat32,
		float32(math.Inf(1)),
		float32(math.Inf(-1)),
		float32(math.NaN()),
	}

	for _, value := range testValues {
		t.Run("", func(t *testing.T) {
			// Преобразование туда и обратно
			bits := Float32ToUint32(value)
			result := Uint32ToFloat32(bits)
			
			// Особые случаи
			if math.IsNaN(float64(value)) {
				if !math.IsNaN(float64(result)) {
					t.Errorf("Round trip failed for NaN: got %f", result)
				}
				return
			}
			
			if math.IsInf(float64(value), 0) {
				if !math.IsInf(float64(result), 0) {
					t.Errorf("Round trip failed for Infinity: got %f", result)
				}
				return
			}
			
			if result != value {
				t.Errorf("Round trip failed: original %f, got %f", value, result)
			}
		})
	}
}