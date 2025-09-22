package debugger

// BreakpointManager управляет точками останова
type BreakpointManager struct {
	breakpoints map[uint16]bool
	disabled    map[uint16]bool
}

// NewBreakpointManager создает новый менеджер точек останова
func NewBreakpointManager() *BreakpointManager {
	return &BreakpointManager{
		breakpoints: make(map[uint16]bool),
		disabled:    make(map[uint16]bool),
	}
}

// SetBreakpoint устанавливает точку останова по адресу
func (bm *BreakpointManager) SetBreakpoint(addr uint16) {
	bm.breakpoints[addr] = true
	delete(bm.disabled, addr)
}

// RemoveBreakpoint удаляет точку останова
func (bm *BreakpointManager) RemoveBreakpoint(addr uint16) {
	delete(bm.breakpoints, addr)
	delete(bm.disabled, addr)
}

// DisableBreakpoint временно отключает точку останова
func (bm *BreakpointManager) DisableBreakpoint(addr uint16) {
	if bm.IsBreakpoint(addr) {
		bm.disabled[addr] = true
	}
}

// EnableBreakpoint включает точку останова
func (bm *BreakpointManager) EnableBreakpoint(addr uint16) {
	delete(bm.disabled, addr)
}

// IsBreakpoint проверяет, установлена ли точка останова по адресу
func (bm *BreakpointManager) IsBreakpoint(addr uint16) bool {
	return bm.breakpoints[addr]
}

// IsEnabled проверяет, активна ли точка останова
func (bm *BreakpointManager) IsEnabled(addr uint16) bool {
	return bm.breakpoints[addr] && !bm.disabled[addr]
}

// GetBreakpoints возвращает список всех точек останова
func (bm *BreakpointManager) GetBreakpoints() []uint16 {
	var addrs []uint16
	for addr := range bm.breakpoints {
		addrs = append(addrs, addr)
	}
	return addrs
}