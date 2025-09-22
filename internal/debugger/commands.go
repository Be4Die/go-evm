package debugger

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
)

// CommandHandler обрабатывает команды отладчика
type CommandHandler struct {
	debugger   *Debugger
	history    []string
}

// NewCommandHandler создает новый обработчик команд
func NewCommandHandler(debugger *Debugger) *CommandHandler {
	return &CommandHandler{
		debugger: debugger,
		history:  []string{},
	}
}

// HandleCommand обрабатывает введенную команду
func (ch *CommandHandler) HandleCommand(cmd string) {
	ch.history = append(ch.history, cmd)
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "run", "r":
		ch.handleRun()
	case "stop":
		ch.handleStop()
	case "step", "s":
		ch.handleStep()
	case "continue", "c":
		ch.handleContinue()
	case "break", "b":
		ch.handleBreakpoint(parts[1:])
	case "delete", "d":
		ch.handleDeleteBreakpoint(parts[1:])
	case "disable":
		ch.handleDisableBreakpoint(parts[1:])
	case "enable":
		ch.handleEnableBreakpoint(parts[1:])
	case "info", "i":
		ch.handleInfo(parts[1:])
	case "registers", "reg":
		ch.handleRegisters()
	case "memory", "mem":
		ch.handleMemory(parts[1:])
	case "stack":
		ch.handleStack()
	case "trace":
		ch.handleTrace(parts[1:])
	case "quit", "q":
		os.Exit(0)
	case "symbols", "sym":
		ch.handleSymbols()
	case "help", "h":
		ch.handleHelp()
	default:
		fmt.Printf("Unknown command: %s\n", parts[0])
	}
}

// Добавим новый метод handleSymbols:
func (ch *CommandHandler) handleSymbols() {
    symbols := ch.debugger.symbols.Symbols
    if len(symbols) == 0 {
        fmt.Println("No symbols loaded")
        return
    }

    fmt.Println("Symbols:")
    for label, addr := range symbols {
        fmt.Printf("  %s: 0x%04X\n", label, addr)
    }
}

// handleRun обрабатывает команду запуска программы
func (ch *CommandHandler) handleRun() {
	fmt.Println("Running program...")
	err := ch.debugger.Run()
	if err != nil {
		fmt.Printf("Execution error: %v\n", err)
	}
}

// handleStop обрабатывает команду остановки программы
func (ch *CommandHandler) handleStop() {
	fmt.Println("Stopping program...")
	ch.debugger.Stop()
}

// handleStep обрабатывает команду пошагового выполнения
func (ch *CommandHandler) handleStep() {
	err := ch.debugger.Step()
	if err != nil {
		fmt.Printf("Step error: %v\n", err)
	} else {
		ch.showCurrentState()
	}
}

// handleContinue обрабатывает команду продолжения выполнения
func (ch *CommandHandler) handleContinue() {
	err := ch.debugger.Continue()
	if err != nil {
		fmt.Printf("Continue error: %v\n", err)
	}
}

// handleBreakpoint обрабатывает команды управления точками останова
func (ch *CommandHandler) handleBreakpoint(args []string) {
	if len(args) == 0 {
		ch.showBreakpoints()
		return
	}

	addr, err := ch.parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	ch.debugger.SetBreakpoint(addr)
	fmt.Printf("Breakpoint set at 0x%04X\n", addr)
}

// handleDeleteBreakpoint обрабатывает удаление точек останова
func (ch *CommandHandler) handleDeleteBreakpoint(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: delete <address>")
		return
	}

	addr, err := ch.parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	ch.debugger.RemoveBreakpoint(addr)
	fmt.Printf("Breakpoint removed at 0x%04X\n", addr)
}

// handleDisableBreakpoint обрабатывает отключение точек останова
func (ch *CommandHandler) handleDisableBreakpoint(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: disable <address>")
		return
	}

	addr, err := ch.parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	ch.debugger.DisableBreakpoint(addr)
	fmt.Printf("Breakpoint disabled at 0x%04X\n", addr)
}

// handleEnableBreakpoint обрабатывает включение точек останова
func (ch *CommandHandler) handleEnableBreakpoint(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: enable <address>")
		return
	}

	addr, err := ch.parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	ch.debugger.EnableBreakpoint(addr)
	fmt.Printf("Breakpoint enabled at 0x%04X\n", addr)
}

// handleInfo обрабатывает команду info
func (ch *CommandHandler) handleInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: info <breakpoints|registers|stack>")
		return
	}

	switch args[0] {
	case "breakpoints", "b":
		ch.showBreakpoints()
	case "registers", "reg":
		ch.showRegisters()
	case "stack":
		ch.showStack()
	default:
		fmt.Printf("Unknown info command: %s\n", args[0])
	}
}

// handleRegisters обрабатывает команду registers
func (ch *CommandHandler) handleRegisters() {
	ch.showRegisters()
}

// handleMemory обрабатывает команду memory
func (ch *CommandHandler) handleMemory(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: memory <address> <count>")
		return
	}

	addr, err := ch.parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Invalid count: %v\n", err)
		return
	}

	ch.showMemory(addr, count)
}

// handleStack обрабатывает команду stack
func (ch *CommandHandler) handleStack() {
	ch.showStack()
}

// handleTrace обрабатывает команду trace
func (ch *CommandHandler) handleTrace(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: trace <on|off>")
		return
	}

	if args[0] == "on" {
		ch.debugger.SetTracing(true)
		fmt.Println("Tracing enabled")
	} else if args[0] == "off" {
		ch.debugger.SetTracing(false)
		fmt.Println("Tracing disabled")
	} else {
		fmt.Println("Usage: trace <on|off>")
	}
}


// Обновим метод handleHelp:
func (ch *CommandHandler) handleHelp() {
    fmt.Println(`Available commands:
  run (r)        - Run program
  stop           - Stop program execution
  step (s)       - Step into instruction
  continue (c)   - Continue execution
  break (b)      - Set breakpoint
  delete (d)     - Delete breakpoint
  disable        - Disable breakpoint
  enable         - Enable breakpoint
  info (i)       - Show information
  registers (reg)- Show registers
  memory (mem)   - Show memory
  stack          - Show stack
  symbols (sym)  - Show symbols 
  trace          - Enable/disable tracing
  quit (q)       - Quit debugger
  help (h)       - Show this help`)
}

// showCurrentState показывает текущее состояние отладчика
func (ch *CommandHandler) showCurrentState() {
	ip := ch.debugger.cpu.GetPSW().GetIP()
	fmt.Printf("Stopped at IP: 0x%04X\n", ip)
	
	// Читаем опкод текущей инструкции
	opcode, err := ch.debugger.memory.ReadByteAt(ip)
	if err == nil {
		fmt.Printf("Opcode: 0x%02X\n", opcode)
	}
	
	ch.showRegisters()
}

// showBreakpoints показывает список точек останова
func (ch *CommandHandler) showBreakpoints() {
	breakpoints := ch.debugger.GetBreakpoints()
	if len(breakpoints) == 0 {
		fmt.Println("No breakpoints set")
		return
	}

	fmt.Println("Breakpoints:")
	for _, addr := range breakpoints {
		status := "enabled"
		if !ch.debugger.IsBreakpointEnabled(addr) {
			status = "disabled"
		}
		
		if label, exists := ch.debugger.symbols.ResolveLabel(addr); exists {
			fmt.Printf("  0x%04X (%s) - %s\n", addr, label, status)
		} else {
			fmt.Printf("  0x%04X - %s\n", addr, status)
		}
	}
}

// showRegisters показывает состояние регистров
func (ch *CommandHandler) showRegisters() {
	psw := ch.debugger.cpu.GetPSW()
	fmt.Printf("Registers:\n")
	fmt.Printf("  IP: 0x%04X\n", psw.GetIP())
	fmt.Printf("  SP: %d\n", psw.GetSP())
	fmt.Printf("  Flags: ")
	
	flags := psw.GetFlags()
	flagNames := []string{"ZERO", "CARRY", "OVERFLOW", "NEGATIVE", "FZERO", "FOVERFLOW", "FNEGATIVE"}
	for i, name := range flagNames {
		if i < 7 && flags&(1<<uint(i)) != 0 {
			fmt.Printf("%s ", name)
		}
	}
	fmt.Println()
}

// showMemory показывает содержимое памяти
func (ch *CommandHandler) showMemory(addr uint16, count int) {
	fmt.Printf("Memory at 0x%04X:\n", addr)
	for i := 0; i < count; i++ {
		currentAddr := addr + uint16(i*4)
		value, err := ch.debugger.memory.ReadWordAt(currentAddr)
		if err != nil {
			fmt.Printf("  Error reading memory at 0x%04X: %v\n", currentAddr, err)
			continue
		}
		
		if label, exists := ch.debugger.symbols.ResolveLabel(currentAddr); exists {
			fmt.Printf("  0x%04X (%s): 0x%08X\n", currentAddr, label, value)
		} else {
			fmt.Printf("  0x%04X: 0x%08X\n", currentAddr, value)
		}
	}
}

// showStack показывает содержимое стека
func (ch *CommandHandler) showStack() {
	sp := ch.debugger.cpu.GetPSW().GetSP()
	fmt.Printf("Stack (SP: %d):\n", sp)
	
	for i := sp; i <= 31; i++ {
		stackAddr := uint16(31-i) * 4
		value, err := ch.debugger.memory.ReadWordAt(stackAddr)
		if err != nil {
			fmt.Printf("  Error reading stack at 0x%04X: %v\n", stackAddr, err)
			continue
		}
		
		marker := " "
		if i == sp {
			marker = ">"
		}
		
		fmt.Printf("  %s [%d] 0x%04X: 0x%08X\n", marker, i, stackAddr, value)
	}
}

// parseAddress парсит адрес из строки
func (ch *CommandHandler) parseAddress(addrStr string) (uint16, error) {
	// Пытаемся разрешить как метку
	if addr, exists := ch.debugger.symbols.ResolveAddress(addrStr); exists {
		return addr, nil
	}
	
	// Пытаемся парсить как число
	addr, err := strconv.ParseUint(addrStr, 0, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid address: %s", addrStr)
	}
	
	return uint16(addr), nil
}

// Completer предоставляет автодополнение для команд
func (ch *CommandHandler) Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "run", Description: "Run program"},
		{Text: "stop", Description: "Stop program execution"},
		{Text: "step", Description: "Step into instruction"},
		{Text: "continue", Description: "Continue execution"},
		{Text: "break", Description: "Set breakpoint"},
		{Text: "delete", Description: "Delete breakpoint"},
		{Text: "disable", Description: "Disable breakpoint"},
		{Text: "enable", Description: "Enable breakpoint"},
		{Text: "info", Description: "Show information"},
		{Text: "registers", Description: "Show registers"},
		{Text: "memory", Description: "Show memory"},
		{Text: "stack", Description: "Show stack"},
		{Text: "trace", Description: "Enable/disable tracing"},
		{Text: "quit", Description: "Quit debugger"},
		{Text: "help", Description: "Show help"},
		{Text: "symbols", Description: "Show symbols"},
        {Text: "sym", Description: "Show symbols"},
	}
	
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}