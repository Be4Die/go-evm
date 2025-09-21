package assembly

import (
	"strings"
	"testing"
)

func TestProcessDataLine(t *testing.T) {
	translator := NewTranslator()
	translator.currentAddress = 0x0100

	tests := []struct {
		name    string
		line    string
		wantErr bool
	}{
		{"DB directive", "DB 10, 20", false},
		{"DW directive", "DW 1000, 2000", false},
		{"label with DB", "DATA: DB 10", false},
		{"invalid directive", "INVALID 10", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset data for each test
			translator.data = make([]dataItem, 0)
			translator.currentAddress = 0x0100
			
			err := translator.processDataLine(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processDataLine(%q) error = %v, wantErr %v", tt.line, err, tt.wantErr)
			}
		})
	}
}

func TestProcessDB(t *testing.T) {
	translator := NewTranslator()
	translator.currentAddress = 0x0100
	translator.pass = 1 // Set to first pass to add data items

	tests := []struct {
		name        string
		line        string
		expectedLen int
		wantErr     bool
	}{
		{"single value", "DB 10", 1, false},
		{"multiple values", "DB 10, 20, 30", 3, false},
		{"hex values", "DB 0x0A, 0x14", 2, false},
		{"binary values", "DB 0b1010, 0b10100", 2, false},
		{"invalid value", "DB INVALID", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset data for each test
			translator.data = make([]dataItem, 0)
			translator.currentAddress = 0x0100
			
			err := translator.processDB(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processDB(%q) error = %v, wantErr %v", tt.line, err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(translator.data) != tt.expectedLen {
				t.Errorf("processDB(%q) created %d items, want %d", tt.line, len(translator.data), tt.expectedLen)
			}
		})
	}
}

func TestProcessDW(t *testing.T) {
	translator := NewTranslator()
	translator.currentAddress = 0x0100
	translator.pass = 1 // Set to first pass to add data items

	tests := []struct {
		name        string
		line        string
		expectedLen int
		wantErr     bool
	}{
		{"single value", "DW 1000", 1, false},
		{"multiple values", "DW 1000, 2000", 2, false},
		{"hex values", "DW 0x03E8, 0x07D0", 2, false},
		{"invalid value", "DW INVALID", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset data for each test
			translator.data = make([]dataItem, 0)
			translator.currentAddress = 0x0100
			
			err := translator.processDW(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processDW(%q) error = %v, wantErr %v", tt.line, err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(translator.data) != tt.expectedLen {
				t.Errorf("processDW(%q) created %d items, want %d", tt.line, len(translator.data), tt.expectedLen)
			}
		})
	}
}

func TestProcessEQU(t *testing.T) {
	translator := NewTranslator()
	translator.pass = 1 // Set to first pass to add constants

	tests := []struct {
		name     string
		line     string
		expected uint32
		wantErr  bool
	}{
		{"numeric constant", "MAX EQU 100", 100, false},
		{"hex constant", "ADDR EQU 0x0100", 0x0100, false},
		{"invalid constant", "INVALID EQU VALUE", 0, true},
		{"malformed directive", "MALFORMED", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset constants for each test
			translator.constants = make(map[string]uint32)
			
			err := translator.processEQU(tt.line, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("processEQU(%q) error = %v, wantErr %v", tt.line, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				parts := strings.Fields(tt.line)
				constName := parts[0]
				if translator.constants[constName] != tt.expected {
					t.Errorf("processEQU(%q) constant value = %d, want %d", 
						tt.line, translator.constants[constName], tt.expected)
				}
			}
		})
	}
}