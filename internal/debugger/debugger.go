package debugger

import (
	"fmt"
	"time"

	"github.com/Be4Die/go-evm/internal/loader"
	"github.com/Be4Die/go-evm/internal/vm"
	"github.com/c-bata/go-prompt"
)

// MemoryReader интерфейс для чтения памяти
type MemoryReader interface {
	ReadByteAt(addr uint16) (byte, error)
	ReadWordAt(addr uint16) (uint32, error)
}

// Debugger представляет отладчик EVM
type Debugger struct {
	cpu           *vm.CPU
	memory        *vm.Memory
	loader        *loader.Loader
	breakpoints   *BreakpointManager
	symbols       *SymbolTable
	cmdHandler    *CommandHandler
	isRunning     bool
	shouldStop    bool
	tracing       bool
	callDepth     int
}

// NewDebugger создает новый отладчик
func NewDebugger() *Debugger {
	memory := vm.NewMemory(64 * 1024)
	cpu := vm.NewCPU(memory)
	ldr := loader.NewLoader()
	
	debugger := &Debugger{
		cpu:         cpu,
		memory:      memory,
		loader:      ldr,
		breakpoints: NewBreakpointManager(),
		symbols:     NewSymbolTable(),
		isRunning:   false,
		shouldStop:  false,
		tracing:     false,
		callDepth:   0,
	}
	
	debugger.cmdHandler = NewCommandHandler(debugger)
	
	return debugger
}

// LoadProgram загружает программу в отладчик
func (d *Debugger) LoadProgram(filename string) error {
	startAddr, err := d.loader.LoadProgram(filename, d.memory)
	if err != nil {
		return fmt.Errorf("failed to load program: %v", err)
	}
	
	d.cpu.GetPSW().SetIP(startAddr)
	fmt.Printf("Program loaded at 0x%04X\n", startAddr)
	
	return nil
}

// LoadSymbols загружает символы из файла
func (d *Debugger) LoadSymbols(filename string) error {
	err := d.symbols.LoadFromFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load symbols: %v", err)
	}
	
	fmt.Printf("Loaded %d symbols\n", len(d.symbols.Symbols))
	return nil
}

// Run запускает выполнение программы
func (d *Debugger) Run() error {
	d.isRunning = true
	d.shouldStop = false
	
	for !d.shouldStop {
		ip := d.cpu.GetPSW().GetIP()
		
		// Проверяем точки останова
		if d.breakpoints.IsEnabled(ip) {
			fmt.Printf("Breakpoint hit at 0x%04X\n", ip)
			d.isRunning = false
			return nil
		}
		
		// Выполняем шаг
		err := d.cpu.Step()
		if err != nil {
			d.isRunning = false
			return fmt.Errorf("execution error: %v", err)
		}
		
		// Если включена трассировка, выводим информацию
		if d.tracing {
			d.printTrace()
		}
		
		// Проверяем завершение программы
		if d.cpu.GetPSW().GetIP() == 0 {
			fmt.Println("Program terminated normally")
			d.isRunning = false
			return nil
		}
		
		// Небольшая задержка для удобства наблюдения
		time.Sleep(10 * time.Millisecond)
	}
	
	d.isRunning = false
	return nil
}

// Stop останавливает выполнение программы
func (d *Debugger) Stop() {
	d.shouldStop = true
}

// Step выполняет один шаг (с заходом в функции)
func (d *Debugger) Step() error {
	ip := d.cpu.GetPSW().GetIP()
	
	// Проверяем, является ли текущая инструкция CALL
	opcode, err := d.memory.ReadByteAt(ip)
	if err != nil {
		return fmt.Errorf("failed to read opcode: %v", err)
	}
	
	if opcode == vm.OP_CALL {
		d.callDepth++
	}
	
	// Выполняем шаг
	err = d.cpu.Step()
	if err != nil {
		return fmt.Errorf("step error: %v", err)
	}
	
	// Если включена трассировка, выводим информацию
	if d.tracing {
		d.printTrace()
	}
	
	return nil
}

// StepOut выполняет шаг с возвратом из функции
func (d *Debugger) StepOut() error {
	if d.callDepth == 0 {
		return d.Step()
	}
	
	// Устанавливаем временную точку останова на инструкцию после RET
	currentDepth := d.callDepth
	d.callDepth = 0
	
	// Запускаем выполнение пока не вернемся на предыдущий уровень
	for {
		err := d.Step()
		if err != nil {
			return err
		}
		
		ip := d.cpu.GetPSW().GetIP()
		opcode, err := d.memory.ReadByteAt(ip)
		if err != nil {
			return fmt.Errorf("failed to read opcode: %v", err)
		}
		
		if opcode == vm.OP_RET {
			d.callDepth = currentDepth - 1
			break
		}
	}
	
	return nil
}

// Continue продолжает выполнение до следующей точки останова
func (d *Debugger) Continue() error {
	return d.Run()
}

// SetBreakpoint устанавливает точку останова
func (d *Debugger) SetBreakpoint(addr uint16) {
	d.breakpoints.SetBreakpoint(addr)
}

// RemoveBreakpoint удаляет точку останова
func (d *Debugger) RemoveBreakpoint(addr uint16) {
	d.breakpoints.RemoveBreakpoint(addr)
}

// DisableBreakpoint отключает точку останова
func (d *Debugger) DisableBreakpoint(addr uint16) {
	d.breakpoints.DisableBreakpoint(addr)
}

// EnableBreakpoint включает точку останова
func (d *Debugger) EnableBreakpoint(addr uint16) {
	d.breakpoints.EnableBreakpoint(addr)
}

// IsBreakpointEnabled проверяет, активна ли точка останова
func (d *Debugger) IsBreakpointEnabled(addr uint16) bool {
	return d.breakpoints.IsEnabled(addr)
}

// GetBreakpoints возвращает список точек останова
func (d *Debugger) GetBreakpoints() []uint16 {
	return d.breakpoints.GetBreakpoints()
}

// SetTracing включает или выключает трассировку
func (d *Debugger) SetTracing(enabled bool) {
	d.tracing = enabled
}

// printTrace выводит информацию трассировки
func (d *Debugger) printTrace() {
	ip := d.cpu.GetPSW().GetIP()
	sp := d.cpu.GetPSW().GetSP()
	
	// Читаем опкод текущей инструкции
	opcode, err := d.memory.ReadByteAt(ip)
	if err != nil {
		fmt.Printf("TRACE: IP=0x%04X SP=%d\n", ip, sp)
		return
	}
	
	fmt.Printf("TRACE: IP=0x%04X SP=%d | OPCODE=0x%02X\n", ip, sp, opcode)
}

// StartInteractive запускает интерактивный режим отладчика
func (d *Debugger) StartInteractive() {
	fmt.Println("EVM Debugger v0.1")
	fmt.Println("Type 'help' for commands")
	
	// Запускаем интерактивную оболочку
	p := prompt.New(
		d.cmdHandler.HandleCommand,
		d.cmdHandler.Completer,
		prompt.OptionPrefix("(dbg) "),
		prompt.OptionTitle("EVM Debugger"),
	)
	
	p.Run()
}