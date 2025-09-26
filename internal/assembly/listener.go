package assembly

import (
	"fmt"

	"github.com/Be4Die/go-evm/internal/assembly/grammar"
)

// AssemblerListener реализует listener для обхода AST
type AssemblerListener struct {
	*grammar.BaseAssemblerListener
	translator *Translator
}

// NewAssemblerListener создает новый listener
func NewAssemblerListener(translator *Translator) *AssemblerListener {
	return &AssemblerListener{
		translator: translator,
	}
}

// EnterEntryDirective обрабатывает директиву ENTRY
func (l *AssemblerListener) EnterEntryDirective(ctx *grammar.EntryDirectiveContext) {
	if l.translator.pass == 1 {
		label := ctx.LabelName().GetText()
		l.translator.entryLabel = label
	}
}

// EnterSectionDirective обрабатывает директиву SECTION
func (l *AssemblerListener) EnterSectionDirective(ctx *grammar.SectionDirectiveContext) {
	section := ctx.SectionName().GetText()
	l.translator.currentSection = section

	if section == ".data" {
		l.translator.currentAddress = l.translator.dataAddress
	} else if section == ".code" {
		l.translator.currentAddress = l.translator.startAddress
	}
}

// EnterDataDirective обрабатывает директивы DB/DW
func (l *AssemblerListener) EnterDataDirective(ctx *grammar.DataDirectiveContext) {
    if l.translator.pass == 1 {
        // Обработка метки (с проверкой на nil)
        if ctx.LabelName() != nil {
            label := ctx.LabelName().GetText()
            l.translator.symbolTable[label] = l.translator.currentAddress
        }

		// Обработка данных
		allConstants := ctx.AllConstant()
		isByte := ctx.DB() != nil

		for _, constantCtx := range allConstants {
			value, err := l.parseConstant(constantCtx)
			if err != nil {
				// Ошибка обрабатывается в error listener
				continue
			}

			l.translator.data = append(l.translator.data, DataItem{
				Address: l.translator.currentAddress,
				Value:   value,
				IsByte:  isByte,
			})

			if isByte {
				l.translator.currentAddress++
			} else {
				l.translator.currentAddress += 4
			}
		}
	}
}

// EnterInstruction обрабатывает инструкции
func (l *AssemblerListener) EnterInstruction(ctx *grammar.InstructionContext) {
    var label string
    // Обработка метки (с проверкой на nil)
    if ctx.LabelName() != nil {
        label = ctx.LabelName().GetText()
        if l.translator.pass == 1 {
            l.translator.symbolTable[label] = l.translator.currentAddress
        }
    }

    mnemonicNode := ctx.Mnemonic()
    if mnemonicNode == nil {
        return // Если мнемоника отсутствует, пропускаем инструкцию
    }
    
    mnemonic := mnemonicNode.GetText()
    opcode, err := l.translator.getOpcode(mnemonic)
    if err != nil {
        return
    }

    var operand uint32
    var operandStr string

    // Обработка операнда (с проверкой на nil)
    if ctx.Operand() != nil {
        operandStr = ctx.Operand().GetText()
        if l.translator.pass == 2 {
            operand, err = l.parseOperand(ctx.Operand())
            if err != nil {
                return
            }
        }
    }

    // Специальная обработка HALT
    if mnemonic == "HALT" {
        opcode = 0x0C // JMP
        operand = 0x0000
        operandStr = "0x0000"
    }

    if l.translator.pass == 1 {
        l.translator.instructions = append(l.translator.instructions, Instruction{
            Address:    l.translator.currentAddress,
            Opcode:     opcode,
            Operand:    0,
            Label:      label, // Используем переменную label вместо прямого вызова
            Mnemonic:   mnemonic,
            OperandStr: operandStr,
        })
        l.translator.currentAddress += 3
    } else {
        l.translator.instructions = append(l.translator.instructions, Instruction{
            Address:    l.translator.currentAddress,
            Opcode:     opcode,
            Operand:    operand,
            Label:      label, // Используем переменную label вместо прямого вызова
            Mnemonic:   mnemonic,
            OperandStr: operandStr,
        })
        l.translator.currentAddress += 3
    }
}

// EnterLabel обрабатывает отдельные метки
func (l *AssemblerListener) EnterLabel(ctx *grammar.LabelContext) {
	if l.translator.pass == 1 {
		label := ctx.LabelName().GetText()
		l.translator.symbolTable[label] = l.translator.currentAddress
	}
}

// parseConstant преобразует константу в число
func (l *AssemblerListener) parseConstant(ctx grammar.IConstantContext) (uint32, error) {
	if ctx.Number() != nil {
		return l.translator.parseConstant(ctx.Number().GetText())
	}
	if ctx.LabelName() != nil {
		return l.translator.parseConstant(ctx.LabelName().GetText())
	}
	return 0, fmt.Errorf("invalid constant")
}

// parseOperand преобразует операнд в число
func (l *AssemblerListener) parseOperand(ctx grammar.IOperandContext) (uint32, error) {
	if indirectCtx, ok := ctx.(*grammar.IndirectOperandContext); ok {
		return l.parseConstant(indirectCtx.IndirectAddress().Constant())
	}
	if directCtx, ok := ctx.(*grammar.DirectOperandContext); ok {
		return l.parseConstant(directCtx.Constant())
	}
	return 0, fmt.Errorf("invalid operand type")
}