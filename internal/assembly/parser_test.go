// parser_test.go
package assembly

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "semicolon comment",
			input:    "MOV AX, BX ; This is a comment",
			expected: "MOV AX, BX",
		},
		{
			name:     "double slash comment",
			input:    "ADD_I 10 // This is another comment",
			expected: "ADD_I 10",
		},
		{
			name:     "mixed comments",
			input:    "SUB_I CX, DX ; comment // with nested",
			expected: "SUB_I CX, DX",
		},
		{
			name:     "no comments",
			input:    "JMP LABEL",
			expected: "JMP LABEL",
		},
		{
			name:     "only comment",
			input:    "; Just a comment",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeComments(tt.input)
			if result != tt.expected {
				t.Errorf("removeComments(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestProcessCodeLine(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantErr    bool
		checkLabel string
	}{
		{
			name:       "label only",
			line:       "LABEL:",
			wantErr:    false,
			checkLabel: "LABEL",
		},
		{
			name:       "label with instruction",
			line:       "LOOP: MOV 0x10",
			wantErr:    false,
			checkLabel: "LOOP",
		},
		{
			name:    "EQU directive with name",
			line:    "CONST EQU 100",
			wantErr: false,
		},
		{
			name:    "instruction",
			line:    "ADD_I 0x20",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := NewTranslator()
			translator.currentAddress = 0x0200
			translator.pass = 1 // Set to first pass to process labels
			
			err := translator.processCodeLine(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processCodeLine() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check if label was added to symbol table
			if tt.checkLabel != "" {
				if _, exists := translator.symbolTable[tt.checkLabel]; !exists {
					t.Errorf("label %s not found in symbol table", tt.checkLabel)
				}
			}
		})
	}
}