package vm

// реализует безусловный переход
type jumpCommand struct {
    target OperandReader
}

func (j *jumpCommand) Execute(cpu CPU, memory Memory) error {
    target, err := j.target.GetValue(cpu, memory)
    if err != nil {
        return err
    }
    return cpu.SetRegister(REG_IP, target)
}

// реализует условный переход
type conditionalJumpCommand struct {
    target    OperandReader
    condition func(Word) bool
}

func (cj *conditionalJumpCommand) Execute(cpu CPU, memory Memory) error {
    flags, err := cpu.GetRegister(REG_FLAGS)
    if err != nil {
        return err
    }
    
    if cj.condition(flags) {
        target, err := cj.target.GetValue(cpu, memory)
        if err != nil {
            return err
        }
        return cpu.SetRegister(REG_IP, target)
    }
    return nil
}

// реализует вызов подпрограммы
type callCommand struct {
    target OperandReader
}

func (c *callCommand) Execute(cpu CPU, memory Memory) error {
    // Сохраняем адрес возврата (следующая инструкция)
    ip, err := cpu.GetRegister(REG_IP)
    if err != nil {
        return err
    }
    
    // Получаем указатель стека
    sp, err := cpu.GetRegister(REG_SP)
    if err != nil {
        return err
    }
    
    // Сохраняем адрес возврата в стеке
    if err := memory.SetValue(sp, ip); err != nil {
        return err
    }
    
    // Обновляем указатель стека
    if err := cpu.SetRegister(REG_SP, sp-4); err != nil {
        return err
    }
    
    // Переходим по целевому адресу
    target, err := c.target.GetValue(cpu, memory)
    if err != nil {
        return err
    }
    return cpu.SetRegister(REG_IP, target)
}

// реализует возврат из подпрограммы
type returnCommand struct{}

func (r *returnCommand) Execute(cpu CPU, memory Memory) error {
    // Получаем указатель стека
    sp, err := cpu.GetRegister(REG_SP)
    if err != nil {
        return err
    }
    
    // Восстанавливаем адрес возврата из стека
    returnAddr, err := memory.GetValue(sp + 4)
    if err != nil {
        return err
    }
    
    // Обновляем указатель стека
    if err := cpu.SetRegister(REG_SP, sp+4); err != nil {
        return err
    }
    
    // Возвращаем управление
    return cpu.SetRegister(REG_IP, returnAddr)
}

/* -- Конструкторы управляющих команд -- */

func NewJumpCommand(target OperandReader) Command {
    return &jumpCommand{target: target}
}

func NewCallCommand(target OperandReader) Command {
    return &callCommand{target: target}
}

func NewReturnCommand() Command {
    return &returnCommand{}
}

func NewJumpIfEqualCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&ZF != 0
        },
    }
}

func NewJumpIfNotEqualCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&ZF == 0
        },
    }
}

func NewJumpIfGreaterCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&ZF == 0 && flags&SF == flags&OF
        },
    }
}

func NewJumpIfGreaterOrEqualCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&SF == flags&OF
        },
    }
}

func NewJumpIfLessCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&SF != flags&OF
        },
    }
}

func NewJumpIfLessOrEqualCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&ZF != 0 || flags&SF != flags&OF
        },
    }
}

func NewJumpIfCarryCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&CF != 0
        },
    }
}

func NewJumpIfOverflowCommand(target OperandReader) Command {
    return &conditionalJumpCommand{
        target: target,
        condition: func(flags Word) bool {
            return flags&OF != 0
        },
    }
}