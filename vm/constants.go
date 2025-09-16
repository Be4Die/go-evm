// vm/constants.go
package vm

const (
	// Opcodes
	OP_MOV   = 0x01
	OP_ADD_I = 0x02
	OP_SUB_I = 0x03
	OP_MUL_I = 0x04
	OP_DIV_I = 0x05
	OP_ADD_F = 0x06
	OP_SUB_F = 0x07
	OP_MUL_F = 0x08
	OP_DIV_F = 0x09
	OP_CMP_I = 0x0A
	OP_CMP_F = 0x0B
	OP_JMP   = 0x0C
	OP_JZ    = 0x0D
	OP_JNZ   = 0x0E
	OP_JC    = 0x0F
	OP_JNC   = 0x10
	OP_CALL  = 0x11
	OP_RET   = 0x12
	OP_PUSH  = 0x13
	OP_POP   = 0x14
	OP_IN    = 0x15
	OP_OUT   = 0x16
	OP_AND   = 0x17
	OP_OR    = 0x18
	OP_XOR   = 0x19
	OP_NOT   = 0x1A
	OP_SHL   = 0x1B
	OP_SHR   = 0x1C

	// Flag bits
	FLAG_ZERO      = 0  // Zero flag
	FLAG_CARRY     = 1  // Carry flag
	FLAG_OVERFLOW  = 2  // Overflow flag
	FLAG_NEGATIVE  = 3  // Negative flag
	FLAG_FZERO     = 4  // Float zero flag
	FLAG_FOVERFLOW = 5  // Float overflow flag
	FLAG_FNEGATIVE = 6  // Float negative flag
)