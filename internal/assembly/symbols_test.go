package assembly

import (
	"testing"
)

func TestParseConstant(t *testing.T) {
	translator := NewTranslator()
	translator.symbolTable["LABEL"] = 0x0200
	translator.constants["MAX"] = 100

	tests := []struct {
		name     string
		input    string
		expected uint32
		wantErr  bool
	}{
		{"hex number", "0x10", 16, false},
		{"binary number", "0b1010", 10, false},
		{"decimal number", "42", 42, false},
		{"symbol", "LABEL", 0x0200, false},
		{"constant", "MAX", 100, false},
		{"invalid format", "INVALID", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := translator.parseConstant(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConstant(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("parseConstant(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}