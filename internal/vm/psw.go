package vm

// PSW представляет слово состояния процессора (Processor Status Word).
// Содержит регистры управления выполнением программы и флаги состояния.
type PSW struct {
	ip    uint16 // Instruction Pointer - указатель на следующую инструкцию
	sp    uint8  // Stack Pointer - указатель вершины стека
	flags uint16 // регистр флагов состояния процессора
}

// NewPSW создает и возвращает новый экземпляр PSW.
// Инициализирует указатель стека значением 31 (вершина стека размером 31 слово).
func NewPSW() *PSW {
	return &PSW{
		sp: 31, // Инициализация SP на вершине стека (31 слово)
	}
}

// GetIP возвращает текущее значение указателя инструкций (IP).
func (p *PSW) GetIP() uint16 {
	return p.ip
}

// SetIP устанавливает значение указателя инструкций (IP).
func (p *PSW) SetIP(ip uint16) {
	p.ip = ip
}

// GetSP возвращает текущее значение указателя стека (SP).
func (p *PSW) GetSP() uint8 {
	return p.sp
}

// SetSP устанавливает значение указателя стека (SP).
func (p *PSW) SetSP(sp uint8) {
	p.sp = sp
}

// GetFlags возвращает текущее значение регистра флагов.
func (p *PSW) GetFlags() uint16 {
	return p.flags
}

// SetFlags устанавливает значение регистра флагов.
func (p *PSW) SetFlags(flags uint16) {
	p.flags = flags
}

// SetFlag устанавливает или сбрасывает конкретный бит в регистре флагов.
// bit: номер бита (0-15) для установки
// value: true - установить бит, false - сбросить бит
func (p *PSW) SetFlag(bit uint8, value bool) {
	if value {
		p.flags |= 1 << bit
	} else {
		p.flags &^= 1 << bit
	}
}

// GetFlag возвращает значение конкретного бита из регистра флагов.
// bit: номер бита (0-15) для чтения
// Возвращает true если бит установлен, false если сброшен.
func (p *PSW) GetFlag(bit uint8) bool {
	return (p.flags>>bit)&1 == 1
}