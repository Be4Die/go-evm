// operations.go
package vm

/* -- Абстракция операционных команд -- */

// unaryOperation представляет функцию унарной операции
type unaryOperation func(operand Word) (result Word, carry bool, overflow bool, auxiliaryCarry bool)

// binaryOperation представляет функцию бинарной операции
type binaryOperation func(operand1, operand2 Word) (result Word, carry bool, overflow bool, auxiliaryCarry bool)

// команда для унарных операций
type unaryCommand struct {
    op        OperandReadWriter
    operation unaryOperation
}

// выполнение унарной команды
func (c *unaryCommand) Execute(cpu CPU, memory Memory) error {
    val, err := c.op.GetValue(cpu, memory)
    if err != nil {
        return err
    }
    
    result, carry, overflow, aux := c.operation(val)
    
    flags, _ := cpu.GetRegister(REG_FLAGS)
    flags = updateFlags(flags, result, carry, overflow, aux)
    cpu.SetRegister(REG_FLAGS, flags)
    
    return c.op.SetValue(cpu, memory, result)
}

// команда для бинарных операций
type binaryCommand struct {
    op1     OperandReader
    op2     OperandReader
    result  OperandWriter
    operation binaryOperation
}

// выполнение бинарной команды
func (bc *binaryCommand) Execute(cpu CPU, memory Memory) error {
    op1, err := bc.op1.GetValue(cpu, memory)
    if err != nil {
        return err
    }

    op2, err := bc.op2.GetValue(cpu, memory)
    if err != nil {
        return err
    }

    result, carry, overflow, auxiliaryCarry := bc.operation(op1, op2)
    
    flags, err := cpu.GetRegister(REG_FLAGS)
    if err != nil {
        return err
    }

    flags = updateFlags(flags, result, carry, overflow, auxiliaryCarry)
    err = cpu.SetRegister(REG_FLAGS, flags)
    if err != nil {
        return err
    }

    return bc.result.SetValue(cpu, memory, result)
}

