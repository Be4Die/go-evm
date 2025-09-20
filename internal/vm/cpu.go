package vm

import (
	"errors"
	"fmt"
	"math"
)

// CPU представляет центральный процессор виртуальной машины.
// Обрабатывает выполнение инструкций, управление памятью и регистрами состояния.
type CPU struct {
	memory   *Memory               // Память виртуальной машины
	psw      *PSW                  // Слово состояния процессора
	commands map[uint8]CommandFunc // Карта опкодов и соответствующих функций команд
}

// NewCPU создает и возвращает новый экземпляр CPU.
// memory: экземпляр памяти для работы процессора
func NewCPU(memory *Memory) *CPU {
	cpu := &CPU{
		memory: memory,
		psw:    NewPSW(),
	}
	cpu.initCommands()
	return cpu
}

// GetPSW возвращает текущее слово состояния процессора (PSW).
func (c *CPU) GetPSW() *PSW {
	return c.psw
}

// Step выполняет одну инструкцию по текущему адресу IP.
// Возвращает ошибку в случае неизвестного опкода или проблем с памятью.
func (c *CPU) Step() error {
	opcode, err := c.memory.ReadByteAt(c.psw.ip)
	if err != nil {
		return err
	}

	command, exists := c.commands[opcode]
	if !exists {
		return fmt.Errorf("unknown opcode: %02X at address %04X", opcode, c.psw.ip)
	}

	c.psw.ip++
	
	if err := command(); err != nil {
		return err
	}
	
	fmt.Printf("  IP now: %04X, SP: %d\n", c.psw.ip, c.psw.sp)
	return nil
}

// Run выполняет инструкции до тех пор, пока IP не станет равным 0.
// Возвращает ошибку в случае выхода IP за границы памяти или ошибки выполнения инструкции.
func (c *CPU) Run() error {
	for c.psw.ip != 0 {
		if err := c.Step(); err != nil {
			return err
		}
		
		if c.psw.ip > math.MaxUint16-1 {
			return errors.New("instruction pointer out of bounds")
		}
	}
	return nil
}

// initCommands инициализирует карту команд процессора.
// Сопоставляет опкоды с соответствующими функциями-обработчиками.
func (c *CPU) initCommands() {
	c.commands = make(map[uint8]CommandFunc)
	
	c.commands[OP_MOV] = c.movCommand
	c.commands[OP_ADD_I] = c.addIntCommand
	c.commands[OP_SUB_I] = c.subIntCommand
	c.commands[OP_MUL_I] = c.mulIntCommand
	c.commands[OP_DIV_I] = c.divIntCommand
	c.commands[OP_ADD_F] = c.addFloatCommand
	c.commands[OP_SUB_F] = c.subFloatCommand
	c.commands[OP_MUL_F] = c.mulFloatCommand
	c.commands[OP_DIV_F] = c.divFloatCommand
	c.commands[OP_CMP_I] = c.cmpIntCommand
	c.commands[OP_CMP_F] = c.cmpFloatCommand
	c.commands[OP_JMP] = c.jmpCommand
	c.commands[OP_JZ] = c.jzCommand
	c.commands[OP_JNZ] = c.jnzCommand
	c.commands[OP_JC] = c.jcCommand
	c.commands[OP_JNC] = c.jncCommand
	c.commands[OP_CALL] = c.callCommand
	c.commands[OP_RET] = c.retCommand
	c.commands[OP_PUSH] = c.pushCommand
	c.commands[OP_POP] = c.popCommand
	c.commands[OP_IN] = c.inCommand
	c.commands[OP_OUT] = c.outCommand
	c.commands[OP_AND] = c.andCommand
	c.commands[OP_OR] = c.orCommand
	c.commands[OP_XOR] = c.xorCommand
	c.commands[OP_NOT] = c.notCommand
	c.commands[OP_SHL] = c.shlCommand
	c.commands[OP_SHR] = c.shrCommand
}