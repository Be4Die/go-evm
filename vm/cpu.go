// vm/cpu.go
package vm

import (
	"errors"
)

type CPU struct {
	memory   *Memory
	psw      *PSW
	commands map[uint8]Command
}

func NewCPU(memory *Memory) *CPU {
	cpu := &CPU{
		memory: memory,
		psw:    NewPSW(),
	}
	cpu.initCommands()
	return cpu
}

func (c *CPU) GetPSW() *PSW {
	return c.psw
}

func (c *CPU) Step() error {
    opcode, err := c.memory.ReadByteAt(c.psw.ip)
    if err != nil {
        return err
    }

    command, exists := c.commands[opcode]
    if !exists {
        return errors.New("unknown opcode")
    }

    c.psw.ip++ // Increment IP before execution for immediate values
    if err := command.Execute(); err != nil { // Убрали передачу cpu
        return err
    }
    return nil
}

func (c *CPU) Run() error {
	for {
		if err := c.Step(); err != nil {
			return err
		}
	}
}

func (c *CPU) initCommands() {
	c.commands = make(map[uint8]Command)
	
	// Register all commands
	c.commands[OP_MOV] = CommandFunc(c.movCommand)
	c.commands[OP_ADD_I] = CommandFunc(c.addIntCommand)
	c.commands[OP_SUB_I] = CommandFunc(c.subIntCommand)
	c.commands[OP_MUL_I] = CommandFunc(c.mulIntCommand)
	c.commands[OP_DIV_I] = CommandFunc(c.divIntCommand)
	c.commands[OP_ADD_F] = CommandFunc(c.addFloatCommand)
	c.commands[OP_SUB_F] = CommandFunc(c.subFloatCommand)
	c.commands[OP_MUL_F] = CommandFunc(c.mulFloatCommand)
	c.commands[OP_DIV_F] = CommandFunc(c.divFloatCommand)
	c.commands[OP_CMP_I] = CommandFunc(c.cmpIntCommand)
	c.commands[OP_CMP_F] = CommandFunc(c.cmpFloatCommand)
	c.commands[OP_JMP] = CommandFunc(c.jmpCommand)
	c.commands[OP_JZ] = CommandFunc(c.jzCommand)
	c.commands[OP_JNZ] = CommandFunc(c.jnzCommand)
	c.commands[OP_JC] = CommandFunc(c.jcCommand)
	c.commands[OP_JNC] = CommandFunc(c.jncCommand)
	c.commands[OP_CALL] = CommandFunc(c.callCommand)
	c.commands[OP_RET] = CommandFunc(c.retCommand)
	c.commands[OP_PUSH] = CommandFunc(c.pushCommand)
	c.commands[OP_POP] = CommandFunc(c.popCommand)
	c.commands[OP_IN] = CommandFunc(c.inCommand)
	c.commands[OP_OUT] = CommandFunc(c.outCommand)
	c.commands[OP_AND] = CommandFunc(c.andCommand)
	c.commands[OP_OR] = CommandFunc(c.orCommand)
	c.commands[OP_XOR] = CommandFunc(c.xorCommand)
	c.commands[OP_NOT] = CommandFunc(c.notCommand)
	c.commands[OP_SHL] = CommandFunc(c.shlCommand)
	c.commands[OP_SHR] = CommandFunc(c.shrCommand)
}