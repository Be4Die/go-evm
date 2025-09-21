// instructions.go
package assembly

import (
	"fmt"
	"strings"

	"github.com/Be4Die/go-evm/internal/vm"
)

// processInstruction обрабатывает инструкцию ассемблера
// Выполняет трансляцию мнемоники в опкод и разбор операндов
func (t *Translator) processInstruction(line string, lineNum int) error {
	// Разбор инструкции и операндов
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}
	
	mnemonic := strings.ToUpper(parts[0])
	operandStr := ""
	if len(parts) > 1 {
		operandStr = strings.Join(parts[1:], " ")
	}
	
	// Обработка псевдоинструкций
	if mnemonic == "HALT" {
		mnemonic = "JMP"
		operandStr = "0x0000"
	}
	
	// Получение опкода
	opcode, err := t.getOpcode(mnemonic)
	if err != nil {
		return fmt.Errorf("unknown instruction '%s' at line %d", mnemonic, lineNum)
	}
	
	// Разбор операнда (только во втором проходе)
	var operand uint32 = 0
	if t.pass == 2 {
		operand, err = t.parseOperand(operandStr)
		if err != nil {
			return fmt.Errorf("invalid operand '%s' at line %d: %v", operandStr, lineNum, err)
		}
	}
	
	if t.pass == 1 {
		t.instructions = append(t.instructions, instruction{
			lineNum:    lineNum,
			address:    t.currentAddress,
			opcode:     opcode,
			operand:    0, // В первом проходе операнд не вычисляем
			mnemonic:   mnemonic,
			operandStr: operandStr,
		})
		t.currentAddress += 3
	} else {
		t.instructions = append(t.instructions, instruction{
			lineNum:    lineNum,
			address:    t.currentAddress,
			opcode:     opcode,
			operand:    operand,
			mnemonic:   mnemonic,
			operandStr: operandStr,
		})
		t.currentAddress += 3
	}
	
	return nil
}

// parseOperand преобразует строковое представление операнда в числовое значение
// Поддерживает прямую и косвенную адресацию
func (t *Translator) parseOperand(operandStr string) (uint32, error) {
	if operandStr == "" {
		return 0, nil
	}
	
	// Косвенная адресация [addr]
	if strings.HasPrefix(operandStr, "[") && strings.HasSuffix(operandStr, "]") {
		addrStr := operandStr[1 : len(operandStr)-1]
		return t.parseConstant(addrStr)
	}
	
	// Прямая адресация (для переходов)
	return t.parseConstant(operandStr)
}

// getOpcode возвращает числовой код операции для мнемоники инструкции
// Использует таблицу опкодов виртуальной машины
func (t *Translator) getOpcode(mnemonic string) (int, error) {
	opcodes := map[string]int{
		"MOV":   vm.OP_MOV,
		"ADD_I": vm.OP_ADD_I,
		"SUB_I": vm.OP_SUB_I,
		"MUL_I": vm.OP_MUL_I,
		"DIV_I": vm.OP_DIV_I,
		"ADD_F": vm.OP_ADD_F,
		"SUB_F": vm.OP_SUB_F,
		"MUL_F": vm.OP_MUL_F,
		"DIV_F": vm.OP_DIV_F,
		"CMP_I": vm.OP_CMP_I,
		"CMP_F": vm.OP_CMP_F,
		"JMP":   vm.OP_JMP,
		"JZ":    vm.OP_JZ,
		"JNZ":   vm.OP_JNZ,
		"JC":    vm.OP_JC,
		"JNC":   vm.OP_JNC,
		"CALL":  vm.OP_CALL,
		"RET":   vm.OP_RET,
		"PUSH":  vm.OP_PUSH,
		"POP":   vm.OP_POP,
		"IN":    vm.OP_IN,
		"OUT":   vm.OP_OUT,
		"AND":   vm.OP_AND,
		"OR":    vm.OP_OR,
		"XOR":   vm.OP_XOR,
		"NOT":   vm.OP_NOT,
		"SHL":   vm.OP_SHL,
		"SHR":   vm.OP_SHR,
	}
	
	if opcode, exists := opcodes[mnemonic]; exists {
		return opcode, nil
	}
	
	return 0, fmt.Errorf("unknown mnemonic")
}