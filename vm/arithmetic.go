package vm

/* -- Абстракция арифметической команды для целочисленных -- */

// базовая структура данных для арифметических команд
type arithmeticArguments struct {
    op1     OperandReader
    op2     OperandReader
    result  OperandWriter
}

// операция которая выполняет вычисления
type arithmeticOperation func(a, b Word) (result Word, carry bool, overflow bool, auxiliaryCarry bool)

// arithmeticCommand базовая структура для арифметических команд
type arithmeticCommand struct {
    args      arithmeticArguments
    operation arithmeticOperation
}

// Execute выполняет арифметическую команду
func (ac *arithmeticCommand) Execute(cpu CPU, memory Memory) error {
    op1, err := ac.args.op1.GetValue(cpu, memory)
    if err != nil {
        return err
    }

    op2, err := ac.args.op2.GetValue(cpu, memory)
    if err != nil {
        return err
    }

    result, carry, overflow, auxiliaryCarry := ac.operation(op1, op2)
    
    flags, err := cpu.GetRegister(REG_FLAGS)
    if err != nil {
        return err
    }

    flags = updateFlags(flags, result, carry, overflow, auxiliaryCarry)
    err = cpu.SetRegister(REG_FLAGS, flags)
    if err != nil {
        return err
    }

    return ac.args.result.SetValue(cpu, memory, result)
}

/* -- Конструкторы арифметических команд -- */

// конструктор команды целочисленного сложения
func NewIntAddCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: integerAdd,
    }
}

// конструктор команды целочисленного вычитания
func NewIntSubCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: integerSub,
    }
}

// конструктор команды целочисленного умножения
func NewIntMulCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: integerMul,
    }
}

// конструктор команды целочисленного деления
func NewIntDivCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: integerDiv,
    }
}

// конструктор команды сложения чисел с плавающей запятой
func NewFloatAddCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: floatAdd,
    }
}

// конструктор команды вычитания чисел с плавающей запятой
func NewFloatSubCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: floatSub,
    }
}

// конструктор команды умножения чисел с плавающей запятой
func NewFloatMulCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: floatMul,
    }
}

// конструктор команды деления чисел с плавающей запятой
func NewFloatDivCommand(op1, op2 OperandReader, result OperandWriter) Command {
    return &arithmeticCommand{
        args: arithmeticArguments{
            op1:    op1,
            op2:    op2,
            result: result,
        },
        operation: floatDiv,
    }
}