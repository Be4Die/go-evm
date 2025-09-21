package assembly

import (
	"testing"

	"github.com/Be4Die/go-evm/internal/vm"
)

func TestGetOpcode(t *testing.T) {
	translator := NewTranslator()

	tests := []struct {
		mnemonic string
		expected int
		wantErr  bool
	}{
		{"MOV", vm.OP_MOV, false},
		{"ADD_I", vm.OP_ADD_I, false},
		{"JMP", vm.OP_JMP, false},
		{"INVALID", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.mnemonic, func(t *testing.T) {
			result, err := translator.getOpcode(tt.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOpcode(%q) error = %v, wantErr %v", tt.mnemonic, err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("getOpcode(%q) = %d, want %d", tt.mnemonic, result, tt.expected)
			}
		})
	}
}

func TestParseOperand(t *testing.T) {
	translator := NewTranslator()
	translator.symbolTable["LABEL"] = 0x0200
	translator.constants["MAX"] = 100

	tests := []struct {
		name     string
		operand  string
		expected uint32
		wantErr  bool
	}{
		{"hex literal", "0x10", 0x10, false},
		{"decimal literal", "42", 42, false},
		{"binary literal", "0b1010", 10, false},
		{"symbol", "LABEL", 0x0200, false},
		{"constant", "MAX", 100, false},
		{"indirect addressing", "[0x20]", 0x20, false},
		{"invalid operand", "INVALID", 0, true},
		{"empty operand", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := translator.parseOperand(tt.operand)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOperand(%q) error = %v, wantErr %v", tt.operand, err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("parseOperand(%q) = %d, want %d", tt.operand, result, tt.expected)
			}
		})
	}
}

func TestProcessInstruction(t *testing.T) {
	translator := NewTranslator()
	translator.currentAddress = 0x0200
	translator.symbolTable["LOOP"] = 0x0205

	tests := []struct {
		name    string
		line    string
		wantErr bool
	}{
		{"valid instruction", "MOV 0x10", false},
		{"instruction with label", "JMP LOOP", false},
		{"halt pseudo-instruction", "HALT", false},
		{"invalid instruction", "INVALID OP", true},
		{"empty line", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := translator.processInstruction(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processInstruction(%q) error = %v, wantErr %v", tt.line, err, tt.wantErr)
			}
		})
	}
}