package vm

/* -- Абстракция операндов -- */

// интерфейс для операндов, поддерживающих чтение
type OperandReader interface {
    GetValue(cpu CPU, memory Memory) (Word, error)
}

// интерфейс для операндов, поддерживающих запись
type OperandWriter interface {
    SetValue(cpu CPU, memory Memory, value Word) error
}

// объединяет чтение и запись
type OperandReadWriter interface {
    OperandReader
    OperandWriter
}


/* -- Константный операнд -- */

// тип константного операнда
type constantOperand struct {
    value Word
}

// получение значение константы
func (o *constantOperand) GetValue(cpu CPU, memory Memory) (Word, error) {
    return o.value, nil
}

// конструктор константного операнда
func NewConstantOperand(value Word) OperandReader {
    return &constantOperand{value: value}
}


/* -- Регистровый операнд -- */

// тип регистрового операнда
type registerOperand struct {
    reg RegisterType
}

// получение значения регистра
func (o *registerOperand) GetValue(cpu CPU, memory Memory) (Word, error) {
    return cpu.GetRegister(o.reg)
}

// установка значения регистра
func (o *registerOperand) SetValue(cpu CPU, memory Memory, value Word) error {
    return cpu.SetRegister(o.reg, value)
}

// конструктор регистрового операнда
func NewRegisterOperand(reg RegisterType) OperandReadWriter {
    return &registerOperand{reg: reg}
}


/* -- Адресный операнд -- */

// тип адресного операнда
type addressOperand struct {
    addr Word
}

//  получение по адресу памяти
func (o *addressOperand) GetValue(cpu CPU, memory Memory) (Word, error) {
    return memory.GetValue(o.addr)
}

// установление значения по адресу в памяти
func (o *addressOperand) SetValue(cpu CPU, memory Memory, value Word) error {
    return memory.SetValue(o.addr, value)
}

// конструктор адресного операнда
func NewAddressOperand(addr Word) OperandReadWriter {
    return &addressOperand{addr: addr}
}